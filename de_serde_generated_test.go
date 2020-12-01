package msgpack

import (
	"errors"

	"github.com/Xuanwo/serde-go"
)

type serdeTestEnum = int

const (
	serdeTestEnumA serdeTestEnum = iota + 1
	serdeTestEnumB
	serdeTestEnumC
	serdeTestEnumD
)

type serdeTestFieldVisitor struct {
	e serdeTestEnum

	serde.DummyVisitor
}

func newSerdeTestFieldVisitor() *serdeTestFieldVisitor {
	return &serdeTestFieldVisitor{
		DummyVisitor: serde.NewDummyVisitor("Test Field"),
	}
}

func (s *serdeTestFieldVisitor) VisitString(v string) (err error) {
	switch v {
	case "A":
		s.e = serdeTestEnumA
	case "B":
		s.e = serdeTestEnumB
	case "C":
		s.e = serdeTestEnumC
	case "D":
		s.e = serdeTestEnumD
	default:
		return errors.New("invalid field")
	}
	return nil
}

type serdeTestVisitor struct {
	v *Test

	serde.DummyVisitor
}

func newTestVisitor(v *Test) *serdeTestVisitor {
	return &serdeTestVisitor{
		v:            v,
		DummyVisitor: serde.NewDummyVisitor("Test"),
	}
}

func (s *serdeTestVisitor) VisitMap(m serde.MapAccess) (err error) {
	field := newSerdeTestFieldVisitor()
	for {
		ok, err := m.NextKey(field)
		if !ok {
			break
		}
		if err != nil {
			return err
		}

		var v serde.Visitor
		switch field.e {
		case serdeTestEnumA:
			v = serde.NewStringVisitor(&s.v.A)
		case serdeTestEnumB:
			v = serde.NewStringVisitor(&s.v.B)
		case serdeTestEnumC:
			v = serde.NewInt32Visitor(&s.v.C)
		case serdeTestEnumD:
			v = serde.NewInt64Visitor(&s.v.D)
		default:
			return errors.New("invalid field")
		}
		err = m.NextValue(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Test) Deserialize(de serde.Deserializer) (err error) {
	return de.DeserializeStruct("Test", nil, newTestVisitor(s))
}

func (s *Test) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeStruct("Test", 4)
	if err != nil {
		return err
	}
	err = st.SerializeField(
		serde.StringSerializer("A"),
		serde.StringSerializer(s.A),
	)
	if err != nil {
		return
	}
	err = st.SerializeField(
		serde.StringSerializer("B"),
		serde.StringSerializer(s.B),
	)
	if err != nil {
		return
	}
	err = st.SerializeField(
		serde.StringSerializer("C"),
		serde.Int32Serializer(s.C),
	)
	if err != nil {
		return
	}
	err = st.SerializeField(
		serde.StringSerializer("D"),
		serde.Int64Serializer(s.D),
	)
	if err != nil {
		return
	}
	err = st.EndStruct()
	if err != nil {
		return
	}
	return nil
}
