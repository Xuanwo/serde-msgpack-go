package msgpack

import (
	"log"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

// serde: Deserialize
type Test struct {
	A string
	B string
	C int32
	D int64
}

func TestDe_DeserializeAny(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	ta := Test{
		A: "xxx",
		B: "yyy",
		C: -112323,
		D: 59583,
	}
	content, err := msgpack.Marshal(ta)
	if err != nil {
		t.Errorf("msgpack marshl: %v", err)
	}

	x := Test{}
	err = DeserializeFromBytes(content, &x)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%#+v", x)
}

func BenchmarkDe_DeserializeAny(b *testing.B) {
	ta := Test{
		A: "xxx",
		B: "yyy",
		C: -112323,
		D: 59583,
	}
	content, err := msgpack.Marshal(ta)
	if err != nil {
		b.Errorf("msgpack marshl: %v", err)
	}

	x := Test{}

	b.Run("serde msgpack", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DeserializeFromBytes(content, &x)
		}
	})
	b.Run("vmihailenco msgpack", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = msgpack.Unmarshal(content, &x)
		}
	})
}
