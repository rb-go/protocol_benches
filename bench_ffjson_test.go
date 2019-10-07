package jsonbench_test

import (
	"fmt"
	"log"
	"runtime"
	"testing"

	"github.com/riftbit/protocol_benches/structs"
	"github.com/riftbit/protocol_benches/vars"

	"github.com/pquerna/ffjson/ffjson"
)

var (
	FFJSONExampleJsonBytes    = []byte(vars.ExampleJsonString)
	FFJSONExampleJsonMinBytes []byte
	FFJSONFilledClassicStruct structs.ClassicJSON
	FFJSONFilledFFStruct      structs.FFJSON
)

func init() {
	FFJSONExampleJsonMinBytes = []byte(vars.ExampleJsonMinString)

	err := FFJSONFilledFFStruct.UnmarshalJSON(FFJSONExampleJsonMinBytes)
	if err != nil {
		log.Fatal(err)
	}
	if FFJSONFilledFFStruct.Aliceblue != "#f0f8ff" {
		log.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, FFJSONFilledFFStruct.Aliceblue))
	}

	err = ffjson.Unmarshal(FFJSONExampleJsonMinBytes, &FFJSONFilledClassicStruct)
	if err != nil {
		log.Fatal(err)
	}

	runtime.GC()
}

//
func Benchmark_ffjson_Marshal_simple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, err := ffjson.Marshal(FFJSONFilledClassicStruct)
		if err != nil {
			b.Fatal(err)
		}
		if len(buf) < 1000 {
			b.Fatal(fmt.Errorf(`wrong size of buf - got %d`, len(buf)))
		}
	}
}

//
func Benchmark_ffjson_Marshal_simple_pooling(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, err := ffjson.Marshal(FFJSONFilledClassicStruct)
		if err != nil {
			b.Fatal(err)
		}
		if len(buf) < 1000 {
			b.Fatal(fmt.Errorf(`wrong size of buf - got %d`, len(buf)))
		}
		ffjson.Pool(buf)
	}
}

//
func Benchmark_ffjson_Marshal_generated(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, err := FFJSONFilledFFStruct.MarshalJSON()
		if err != nil {
			b.Fatal(err)
		}
		if len(buf) < 1000 {
			b.Fatal(fmt.Errorf(`wrong size of buf - got %d`, len(buf)))
		}
	}
}

// Benchmark_Unmarshal_Full_Bytes_encode ...
func Benchmark_ffjson_Unmarshal_Full_simple(b *testing.B) {
	// b.SetBytes(int64(len(FFJSONExampleJsonBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.ClassicJSON
		if err := ffjson.Unmarshal(FFJSONExampleJsonBytes, &buf); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

// Benchmark_Unmarshal_Min_Bytes_encode ...
func Benchmark_ffjson_Unmarshal_Min_simple(b *testing.B) {
	// b.SetBytes(int64(len(FFJSONExampleJsonMinBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.ClassicJSON
		if err := ffjson.Unmarshal(FFJSONExampleJsonMinBytes, &buf); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

//
func Benchmark_ffjson_Unmarshal_Full_generated(b *testing.B) {
	// b.SetBytes(int64(len(FFJSONExampleJsonBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.FFJSON
		if err := buf.UnmarshalJSON(FFJSONExampleJsonBytes); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

//
func Benchmark_ffjson_Unmarshal_Min_generated(b *testing.B) {
	// b.SetBytes(int64(len(FFJSONExampleJsonMinBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.FFJSON
		if err := buf.UnmarshalJSON(FFJSONExampleJsonMinBytes); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}
