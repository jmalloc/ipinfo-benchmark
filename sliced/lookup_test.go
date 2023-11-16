package sliced_test

import (
	"testing"

	"github.com/jmalloc/ipinfo-benchmark/sliced"
)

func TestLookup_ViaSlice(t *testing.T) {
	rec, ok := sliced.Lookup("217.64.127.219")
	if !ok {
		t.Fatal("expected record not found")
	}

	{
		expect := "NordVPN"
		if rec.ServiceName != expect {
			t.Fatalf("unexpected service name: got %s, want %s", rec.ServiceName, expect)
		}
	}

	{
		expect := sliced.Hosting | sliced.VPN
		if rec.ServiceTypes != expect {
			t.Fatalf("unexpected service types: got %s, want %s", rec.ServiceTypes, expect)
		}
	}
}

func BenchmarkLookup_ViaSlice_ArbitraryRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, ok := sliced.Lookup("103.205.29.68")
		if !ok {
			b.Fatal("expected record not found")
		}
	}
}

func BenchmarkLookup_ViaSlice_NonExistentRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, ok := sliced.Lookup("8.8.8.8")
		if ok {
			b.Fatal("unexpected record found")
		}
	}
}
