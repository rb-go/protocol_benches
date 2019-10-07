// go test -bench=. -benchmem -cpu 1,4,8 -benchtime 10s -count 1 -cpuprofile=./results/cpu.out -memprofile=./results/mem.out bench_test.go > ./results/old.txt
// go test -bench=. -benchmem -cpu 1,4,8 -benchtime 10s -count 1 -cpuprofile=./results/cpu.out -memprofile=./results/mem.out bench_test.go > ./results/new.txt
// Compare by golang.org/x/tools/cmd/benchcmp
// benchcmp ./results/old.txt ./results/new.txt > ./results/cmp.txt

package jsonbench_test

import (
	"fmt"
	"log"
	"runtime"
	"testing"

	"github.com/riftbit/protocol_benches/structs"
	"github.com/riftbit/protocol_benches/vars"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

var (
	EasyJSONExampleJsonBytes    []byte
	EasyJSONExampleJsonMinBytes []byte
	EasyJSONFilledEasyStruct    structs.EasedJSON
)

func init() {
	EasyJSONExampleJsonBytes = []byte(vars.ExampleJsonString)
	EasyJSONExampleJsonMinBytes = []byte(vars.ExampleJsonMinString)

	err := EasyJSONFilledEasyStruct.UnmarshalJSON(EasyJSONExampleJsonMinBytes)
	if err != nil {
		log.Fatal(err)
	}
	if EasyJSONFilledEasyStruct.Aliceblue != "#f0f8ff" {
		log.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, EasyJSONFilledEasyStruct.Aliceblue))
	}

	runtime.GC()
}

// Benchmark_Marshal_easyjson ...
func Benchmark_easyjson_Marshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, err := EasyJSONFilledEasyStruct.MarshalJSON()
		if err != nil {
			b.Fatal(err)
		}
		if len(buf) < 1000 {
			b.Fatal(fmt.Errorf(`wrong size of buf - got %d`, len(buf)))
		}
	}
}

// Benchmark_Marshal_easyjson_lexer ...
func Benchmark_easyjson_Marshal_lexer(b *testing.B) {
	w := jwriter.Writer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EasyJSONFilledEasyStruct.MarshalEasyJSON(&w)
		buf := w.Buffer.BuildBytes()
		if w.Error != nil {
			b.Fatal(w.Error)
		}
		if len(buf) < 1000 {
			b.Fatal(fmt.Errorf(`wrong size of buf - got %d`, len(buf)))
		}
	}
}

// Benchmark_Unmarshal_Full_Bytes_easyjson ...
func Benchmark_easyjson_Unmarshal_Full(b *testing.B) {
	// b.SetBytes(int64(len(EJUExampleJsonBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.EasedJSON
		if err := buf.UnmarshalJSON(EasyJSONExampleJsonBytes); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

// Benchmark_Unmarshal_Min_Bytes_easyjson ...
func Benchmark_easyjson_Unmarshal_Min(b *testing.B) {
	// b.SetBytes(int64(len(EJUExampleJsonMinBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.EasedJSON
		err := buf.UnmarshalJSON(EasyJSONExampleJsonMinBytes)
		if err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

// Benchmark_Unmarshal_Full_Bytes_easyjson ...
func Benchmark_easyjson_Unmarshal_Full_lexer(b *testing.B) {
	// b.SetBytes(int64(len(EJUExampleJsonBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.EasedJSON
		r := jlexer.Lexer{Data: EasyJSONExampleJsonBytes}
		buf.UnmarshalEasyJSON(&r)
		if r.Error() != nil {
			b.Fatal(r.Error())
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

// Benchmark_Unmarshal_Min_Bytes_easyjson ...
func Benchmark_easyjson_Unmarshal_Min_lexer(b *testing.B) {
	// b.SetBytes(int64(len(EJUExampleJsonMinBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.EasedJSON
		r := jlexer.Lexer{Data: EasyJSONExampleJsonMinBytes}
		buf.UnmarshalEasyJSON(&r)
		if r.Error() != nil {
			b.Fatal(r.Error())
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}

/*

func Benchmark_Marshal_By_Lexer(b *testing.B) {
	b.ResetTimer()
	w := jwriter.Writer{}
	for i := 0; i < b.N; i++ {
		messageBaseJson.MarshalEasyJSON(&w)
	}
}
*/

/*

func Benchmark_Marshal_By_Lexer(b *testing.B) {
	b.ResetTimer()
	w := jwriter.Writer{}
	for i := 0; i < b.N; i++ {
		messageBaseJson.MarshalEasyJSON(&w)
	}
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Snippet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComRiftbitJsonBenches3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Snippet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComRiftbitJsonBenches3(l, v)
}

*/
