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

type serdeMapSerializer_int_int map[int]int

func (s serdeMapSerializer_int_int) Serialize(ser serde.Serializer) (err error) {
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

type serdeSliceVisitor_int struct {
	v *[]int

	serde.DummyVisitor
}

func serdeNewSliceVisitor_int(v *[]int) *serdeSliceVisitor_int {
	return &serdeSliceVisitor_int{
		v:            v,
		DummyVisitor: serde.NewDummyVisitor("[]int"),
	}
}

func (s *serdeSliceVisitor_int) VisitSlice(m serde.SliceAccess) (err error) {
	var value int
	for {
		ok, err := m.NextElement(serde.NewIntVisitor(&value))
		if !ok {
			break
		}
		if err != nil {
			return err
		}

		*s.v = append(*s.v, value)

	}
	return nil
}

type serdeSliceSerializer_int []int

func (s serdeSliceSerializer_int) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeSlice(len(s))
	if err != nil {
		return err
	}

	for _, v := range s {
		err = st.SerializeElement(serde.IntSerializer(v))
		if err != nil {
			return
		}
	}

	err = st.EndSlice()
	if err != nil {
		return
	}
	return nil
}

type serdeSliceVisitor_2_int struct {
	v *[2]int

	serde.DummyVisitor
}

func serdeNewSliceVisitor_2_int(v *[2]int) *serdeSliceVisitor_2_int {
	return &serdeSliceVisitor_2_int{
		v:            v,
		DummyVisitor: serde.NewDummyVisitor("[2]int"),
	}
}

func (s *serdeSliceVisitor_2_int) VisitSlice(m serde.SliceAccess) (err error) {
	var value int
	i := 0
	for {
		ok, err := m.NextElement(serde.NewIntVisitor(&value))
		if !ok {
			break
		}
		if err != nil {
			return err
		}

		(*s.v)[i] = value
		i += 1

	}
	return nil
}

type serdeSliceSerializer_2_int [2]int

func (s serdeSliceSerializer_2_int) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeSlice(len(s))
	if err != nil {
		return err
	}

	for _, v := range s {
		err = st.SerializeElement(serde.IntSerializer(v))
		if err != nil {
			return
		}
	}

	err = st.EndSlice()
	if err != nil {
		return
	}
	return nil
}

type serdePointerVisitor_int struct {
	serde.IntVisitor
}

func serdeNewPointerVisitor_int(v **int) serdePointerVisitor_int {
	// FIXME: nil is not handled correctly
	var tv int
	*v = &tv
	return serdePointerVisitor_int{serde.NewIntVisitor(*v)}
}

func serdePointerSerializer_int(v *int) serde.Serializable {
	if v == nil {
		return serde.NilSerializer{}
	}
	return serde.IntSerializer(*v)
}

type serdeStructEnum_Test = int

const (
	serdeStructEnum_Test_A serdeStructEnum_Test = iota + 1
	serdeStructEnum_Test_B
	serdeStructEnum_Test_C
	serdeStructEnum_Test_D
	serdeStructEnum_Test_M
	serdeStructEnum_Test_S
	serdeStructEnum_Test_varray
	serdeStructEnum_Test_vpointer
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
	case "S":
		s.e = serdeStructEnum_Test_S
	case "varray":
		s.e = serdeStructEnum_Test_varray
	case "vpointer":
		s.e = serdeStructEnum_Test_vpointer
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
		case serdeStructEnum_Test_S:
			v = serdeNewSliceVisitor_int(&s.v.S)
		case serdeStructEnum_Test_varray:
			v = serdeNewSliceVisitor_2_int(&s.v.varray)
		case serdeStructEnum_Test_vpointer:
			v = serdeNewPointerVisitor_int(&s.v.vpointer)
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
	st, err := ser.SerializeStruct("Test", 8)
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
		serdeMapSerializer_int_int(s.M),
	)
	if err != nil {
		return
	}
	err = st.SerializeField(
		serde.StringSerializer("S"),
		serdeSliceSerializer_int(s.S),
	)
	if err != nil {
		return
	}
	err = st.SerializeField(
		serde.StringSerializer("varray"),
		serdeSliceSerializer_2_int(s.varray),
	)
	if err != nil {
		return
	}
	err = st.SerializeField(
		serde.StringSerializer("vpointer"),
		serdePointerSerializer_int(s.vpointer),
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
