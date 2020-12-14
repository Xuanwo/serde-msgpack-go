package msgpack

import (
	"log"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

//go:generate go run -tags tools github.com/Xuanwo/serde-go/cmd/serde ./...

// serde: Deserialize,Serialize
type Test struct {
	A        string
	B        string
	C        int32
	D        int64
	M        map[int]int
	S        []int
	varray   [2]int
	vpointer *int
}

func TestX(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	a := 10
	ta := Test{
		vpointer: &a,
	}
	content, err := SerializeToBytes(&ta)
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

func TestDe_DeserializeAny(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	ta := Test{
		A: "xxx",
		B: "yyy",
		C: -112323,
		D: 59583,
		M: map[int]int{
			1: 2,
		},
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
