package mapped

import (
	"compress/gzip"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"io"
	"net/netip"
	"os"
	"runtime"
	"time"
)

var records map[uint32]Record

// Lookup returns the record associated with the given IP.
func Lookup(ip string) (Record, bool) {
	parsed := netip.MustParseAddr(ip)
	if parsed.Is6() {
		return Record{}, false
	}

	arrayIP := parsed.As4()
	intIP := binary.BigEndian.Uint32(arrayIP[:])

	rec, ok := records[intIP]
	return rec, ok
}

func load() error {
	f, err := os.Open("../standard_privacy.csv.gz")
	if err != nil {
		return err
	}
	defer f.Close()

	z, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer z.Close()

	c := csv.NewReader(z)
	if err != nil {
		return err
	}
	c.ReuseRecord = true

	// Skip the header.
	if _, err := c.Read(); err != nil {
		return err
	}

	for {
		row, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if row[0] != row[1] {
			continue
		}

		ip := netip.MustParseAddr(row[0])
		if ip.Is6() {
			continue
		}

		data := ip.As4()
		key := binary.BigEndian.Uint32(data[:])
		rec := Record{}

		if row[3] != "" {
			rec.ServiceTypes |= Hosting
		}
		if row[4] != "" {
			rec.ServiceTypes |= Proxy
		}
		if row[5] != "" {
			rec.ServiceTypes |= Tor
		}
		if row[6] != "" {
			rec.ServiceTypes |= VPN
		}
		if row[7] != "" {
			rec.ServiceTypes |= Relay
		}

		rec.ServiceName = row[8]

		records[key] = rec
	}

	return nil
}

func init() {
	fmt.Fprintf(os.Stderr, "\n====== MAP ======\n\n")

	start := time.Now()
	records = make(map[uint32]Record, 6_000_000)
	if err := load(); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "csv ingest time = %s\n\n", time.Since(start))

	runtime.GC()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Fprintf(os.Stderr, "memory usage:\n")
	fmt.Fprintf(os.Stderr, "  virtual = %d MB\n", mega(m.Sys))
	fmt.Fprintf(os.Stderr, "  in use  = %d MB\n\n", mega(m.HeapInuse+m.StackInuse))
}

func mega(b uint64) uint64 {
	return b / 1024 / 1024
}
