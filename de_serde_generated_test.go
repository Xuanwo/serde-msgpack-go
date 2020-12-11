package msgpack

import (
	"errors"

	"github.com/Xuanwo/serde-go"
)

type serdeMapVisitor_int_int struct {
	v *map[int]int

	serde.DummyVisitor
}

func serdeNewMapVisitor_int_int(v *map[int]int) *serdeMapVisitor_int_int {
	if *v == nil {
		*v = make(map[int]int)
	}
	return &serdeMapVisitor_int_int{
		v:            v,
		DummyVisitor: serde.NewDummyVisitor("map[int]int"),
	}
}

func (s *serdeMapVisitor_int_int) VisitMap(m serde.MapAccess) (err error) {
	var field int
	var value int
	for {
		ok, err := m.NextKey(serde.NewIntVisitor(&field))
		if !ok {
			break
		}
		if err != nil {
			return err
		}
		err = m.NextValue(serde.NewIntVisitor(&value))
		if err != nil {
			return err
		}
		(*s.v)[field] = value
	}
	return nil
}

type serdeSerializer_int_int map[int]int

func (s serdeSerializer_int_int) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeMap(len(s))
	if err != nil {
		return err
	}

	for k, v := range s {
		err = st.SerializeEntry(
			serde.IntSerializer(k),
			serde.IntSerializer(v),
		)
		if err != nil {
			return
		}
	}

	err = st.EndMap()
	if err != nil {
		return
	}
	return nil
}

type serdeStructEnum_Test = int

const (
	serdeStructEnum_Test_A serdeStructEnum_Test = iota + 1
	serdeStructEnum_Test_B
	serdeStructEnum_Test_C
	serdeStructEnum_Test_D
	serdeStructEnum_Test_M
)

type serdeStructFieldVisitor_Test struct {
	e serdeStructEnum_Test

	serde.DummyVisitor
}

func serdeNewStructFieldVisitor_Test() *serdeStructFieldVisitor_Test {
	return &serdeStructFieldVisitor_Test{
		DummyVisitor: serde.NewDummyVisitor("Test Field"),
	}
}

func (s *serdeStructFieldVisitor_Test) VisitString(v string) (err error) {
	switch v {
	case "A":
		s.e = serdeStructEnum_Test_A
	case "B":
		s.e = serdeStructEnum_Test_B
	case "C":
		s.e = serdeStructEnum_Test_C
	case "D":
		s.e = serdeStructEnum_Test_D
	case "M":
		s.e = serdeStructEnum_Test_M
	default:
		return errors.New("invalid field")
	}
	return nil
}

type serdeStructVisitor_Test struct {
	v *Test

	serde.DummyVisitor
}

func serdeNewStructVisitor_Test(v *Test) *serdeStructVisitor_Test {
	return &serdeStructVisitor_Test{
		v:            v,
		DummyVisitor: serde.NewDummyVisitor("Test"),
	}
}

func (s *serdeStructVisitor_Test) VisitMap(m serde.MapAccess) (err error) {
	field := serdeNewStructFieldVisitor_Test()
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
		case serdeStructEnum_Test_A:
			v = serde.NewStringVisitor(&s.v.A)
		case serdeStructEnum_Test_B:
			v = serde.NewStringVisitor(&s.v.B)
		case serdeStructEnum_Test_C:
			v = serde.NewInt32Visitor(&s.v.C)
		case serdeStructEnum_Test_D:
			v = serde.NewInt64Visitor(&s.v.D)
		case serdeStructEnum_Test_M:
			v = serdeNewMapVisitor_int_int(&s.v.M)
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
	return de.DeserializeStruct("Test", nil, serdeNewStructVisitor_Test(s))
}

func (s *Test) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeStruct("Test", 5)
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
	err = st.SerializeField(
		serde.StringSerializer("M"),
		serdeSerializer_int_int(s.M),
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
