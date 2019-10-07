package jsonbench_test

import (
	"math"
	"testing"
)

var (
	SimpleTestValue32 int32
	SimpleTestValue64 int64
)

//
func Benchmark_simple_32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if SimpleTestValue32 < math.MaxInt32 {
			SimpleTestValue32++
		} else {
			SimpleTestValue32 = 0
		}
	}
}

//
func Benchmark_simple_64(b *testing.B) {
	// b.SetBytes(int64(len(StdExampleJsonBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if SimpleTestValue64 < math.MaxInt64 {
			SimpleTestValue64++
		} else {
			SimpleTestValue64 = 0
		}
	}
}
