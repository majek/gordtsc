package gordtsc

import (
	"fmt"
	"testing"
)

func TestRdtsc(t *testing.T) {
	t0 := Start()
	t1 := Stop()
	td := t1 - t0
	if td < 0 || td > 1000000 {
		t.Fatal("Something's wring with RDTSC clock counter")
	}
}

func BenchmarkRdtsc(b *testing.B) {
	sum := uint64(0)
	for i := 0; i < b.N; i++ {
		t0 := Start()
		t1 := Stop()
		sum += t1 - t0
	}
	fmt.Printf("Empty run: %.3f cycles\n", float32(sum)/float32(b.N))
}
