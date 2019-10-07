package jsonbench_test

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"testing"

	"github.com/riftbit/protocol_benches/structs"
	"github.com/riftbit/protocol_benches/vars"

	"github.com/vmihailenco/msgpack"
)

var (
	MSGPackExampleJSONBytes    []byte
	MSGPackExampleMsgPackBytes []byte
	MSGPackFilledClassicStruct structs.ClassicJSON
)

func init() {
	MSGPackExampleJSONBytes = []byte(vars.ExampleJsonMinString)
	err := json.Unmarshal(MSGPackExampleJSONBytes, &MSGPackFilledClassicStruct)
	if err != nil {
		log.Fatal(err)
	}
	MSGPackExampleMsgPackBytes, _ = msgpack.Marshal(&MSGPackFilledClassicStruct)
	runtime.GC()
}

//
func Benchmark_msgpack_vmihailenco_Marshal(b *testing.B) {
	// b.SetBytes(int64(len(MSGPackExampleMsgPackBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, err := msgpack.Marshal(MSGPackExampleMsgPackBytes)
		if err != nil {
			b.Fatal(err)
		}
		if len(buf) < 1000 {
			b.Fatal(fmt.Errorf(`wrong size of buf - got %d`, len(buf)))
		}
	}
}

//
func Benchmark_msgpack_vmihailenco_Unmarshal(b *testing.B) {
	// b.SetBytes(int64(len(MSGPackExampleMsgPackBytes)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf structs.ClassicJSON
		if err := msgpack.Unmarshal(MSGPackExampleMsgPackBytes, &buf); err != nil {
			b.Fatal(err)
		}
		if buf.Aliceblue != "#f0f8ff" {
			b.Fatal(fmt.Errorf(`wrong data in Aliceblue - got "%s" but expected "#f0f8ff"`, buf.Aliceblue))
		}
	}
}
