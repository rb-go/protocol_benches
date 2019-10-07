package jsonbench_test

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"testing"

	"github.com/riftbit/protocol_benches/structs"
	"github.com/riftbit/protocol_benches/vars"
)

var (
	StdExampleJsonBytes    []byte
	StdExampleJsonMinBytes []byte
	StdFilledClassicStruct structs.ClassicJSON
)

func init() {
	StdExampleJsonBytes = []byte(vars.ExampleJsonString)
	StdExampleJsonMinBytes = []byte(vars.ExampleJsonMinString)

	err := json.Unmarshal(StdExampleJsonMinBytes, &StdFilledClassicStruct)
	if err != nil {
		log.Fatal(err)
	}

	runtime.GC()
}

// Benchmark_Marshal_Bytes_encode ...
func Benchmark_std_Marshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, err := json.Marshal(StdFilledClassicStruct)
		if err != nil {
			b.Fatal(err)
		}
		if len(buf) < 1000 {
			b.Fatal(fmt.Errorf(`wrong size of buf - got %d`, len(buf)))
		}
	}
}

// Benchmark_Unmarshal_Full_Bytes_encode ...
func Benchmark_std_Unmarshal_Full(b *testing.B) {
	// b.SetBytes(int64(len(StdExampleJsonBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.ClassicJSON
		if err := json.Unmarshal(StdExampleJsonBytes, &buf); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

// Benchmark_Unmarshal_Min_Bytes_encode ...
func Benchmark_std_Unmarshal_Min(b *testing.B) {
	// b.SetBytes(int64(len(StdExampleJsonMinBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.ClassicJSON
		if err := json.Unmarshal(StdExampleJsonMinBytes, &buf); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}
