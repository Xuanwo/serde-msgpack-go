package msgpack

import (
	"log"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func TestSer(t *testing.T) {
	x := Test{
		A: "test",
		B: "qqq",
		C: -123213,
		D: 99855,
		M: map[int]int{
			1: 2,
		},
	}

	bs, err := SerializeToBytes(&x)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%v", bs)

	y := Test{}
	err = msgpack.Unmarshal(bs, &y)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%#+v", y)
}

func Benchmark_Serialize(b *testing.B) {
	x := Test{
		A: "test",
		B: "qqq",
		C: -123213,
		D: 99855,
	}

	b.Run("serde msgpack", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = SerializeToBytes(&x)
		}
	})

	b.Run("vmihailenco msgpack", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = msgpack.Marshal(&x)
		}
	})

}
