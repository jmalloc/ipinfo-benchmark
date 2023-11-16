package mapped_test

import (
	"testing"

	"github.com/jmalloc/ipinfo-benchmark/mapped"
)

func TestLookup_ViaMap(t *testing.T) {
	rec, ok := mapped.Lookup("217.64.127.219")
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
		expect := mapped.Hosting | mapped.VPN
		if rec.ServiceTypes != expect {
			t.Fatalf("unexpected service types: got %s, want %s", rec.ServiceTypes, expect)
		}
	}
}

func BenchmarkLookup_ViaMap_ArbitraryRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, ok := mapped.Lookup("103.205.29.68")
		if !ok {
			b.Fatal("expected record not found")
		}
	}
}

func BenchmarkLookup_ViaMap_NonExistentRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, ok := mapped.Lookup("8.8.8.8")
		if ok {
			b.Fatal("unexpected record found")
		}
	}
}
