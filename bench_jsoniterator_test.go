package jsonbench_test

import (
	"fmt"
	"log"
	"runtime"
	"testing"

	"github.com/riftbit/protocol_benches/structs"
	"github.com/riftbit/protocol_benches/vars"

	jsoniter "github.com/json-iterator/go"
)

var (
	JSONIteratorExampleJsonBytes    = []byte(vars.ExampleJsonString)
	JSONIteratorExampleJsonMinBytes []byte
	JSONIteratorFilledClassicStruct structs.ClassicJSON
	JSONIteratorjsoniterator        = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONIteratorjsoniteratorfast    = jsoniter.ConfigFastest
)

func init() {
	JSONIteratorExampleJsonBytes = []byte(vars.ExampleJsonString)
	JSONIteratorExampleJsonMinBytes = []byte(vars.ExampleJsonMinString)

	err := JSONIteratorjsoniterator.Unmarshal(JSONIteratorExampleJsonMinBytes, &JSONIteratorFilledClassicStruct)
	if err != nil {
		log.Fatal(err)
	}

	runtime.GC()
}

//
func Benchmark_jsoniterator_Marshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, err := JSONIteratorjsoniterator.Marshal(JSONIteratorFilledClassicStruct)
		if err != nil {
			b.Fatal(err)
		}
		if len(buf) < 1000 {
			b.Fatal(fmt.Errorf(`wrong size of buf - got %d`, len(buf)))
		}
	}
}

//
func Benchmark_jsoniterator_Marshal_fast(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, err := JSONIteratorjsoniteratorfast.Marshal(JSONIteratorFilledClassicStruct)
		if err != nil {
			b.Fatal(err)
		}
		if len(buf) < 1000 {
			b.Fatal(fmt.Errorf(`wrong size of buf - got %d`, len(buf)))
		}
	}
}

//
func Benchmark_jsoniterator_Unmarshal_Full(b *testing.B) {
	// b.SetBytes(int64(len(JSONIteratorExampleJsonBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.ClassicJSON
		if err := JSONIteratorjsoniterator.Unmarshal(JSONIteratorExampleJsonBytes, &buf); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

//
func Benchmark_jsoniterator_Unmarshal_Min(b *testing.B) {
	// b.SetBytes(int64(len(JSONIteratorExampleJsonMinBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.ClassicJSON
		if err := JSONIteratorjsoniterator.Unmarshal(JSONIteratorExampleJsonMinBytes, &buf); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

//
func Benchmark_jsoniterator_Unmarshal_Full_fast(b *testing.B) {
	// b.SetBytes(int64(len(JSONIteratorExampleJsonBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.ClassicJSON
		if err := JSONIteratorjsoniteratorfast.Unmarshal(JSONIteratorExampleJsonBytes, &buf); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

//
func Benchmark_jsoniterator_Unmarshal_Min_fast(b *testing.B) {
	// b.SetBytes(int64(len(JSONIteratorExampleJsonMinBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.ClassicJSON
		if err := JSONIteratorjsoniteratorfast.Unmarshal(JSONIteratorExampleJsonMinBytes, &buf); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}
