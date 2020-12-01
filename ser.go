package msgpack

import (
	"bytes"

	"github.com/Xuanwo/serde-go"
	"github.com/vmihailenco/msgpack/v5"
)

func SerializeToBytes(v serde.Serializable) ([]byte, error) {
	mser := msgpack.GetEncoder()
	defer msgpack.PutEncoder(mser)

	var buf bytes.Buffer
	mser.Reset(&buf)

	ser := ser{ser: mser}

	err := v.Serialize(&ser)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type ser struct {
	ser *msgpack.Encoder
}

func (s *ser) SerializeBool(v bool) (err error) {
	return s.ser.EncodeBool(v)
}

func (s *ser) SerializeInt(v int) (err error) {
	return s.ser.EncodeInt(int64(v))
}

func (s *ser) SerializeInt8(v int8) (err error) {
	return s.ser.EncodeInt8(v)
}

func (s *ser) SerializeInt16(v int16) (err error) {
	return s.ser.EncodeInt16(v)
}

func (s *ser) SerializeInt32(v int32) (err error) {
	return s.ser.EncodeInt32(v)
}

func (s *ser) SerializeInt64(v int64) (err error) {
	return s.ser.EncodeInt64(v)
}

func (s *ser) SerializeUint(v uint) (err error) {
	return s.ser.EncodeUint(uint64(v))
}

func (s *ser) SerializeUint8(v uint8) (err error) {
	return s.ser.EncodeUint8(v)
}

func (s *ser) SerializeUint16(v uint16) (err error) {
	return s.ser.EncodeUint16(v)
}

func (s *ser) SerializeUint32(v uint32) (err error) {
	return s.ser.EncodeUint32(v)
}

func (s *ser) SerializeUint64(v uint64) (err error) {
	return s.ser.EncodeUint64(v)
}

func (s *ser) SerializeFloat32(v float32) (err error) {
	return s.ser.EncodeFloat32(v)
}

func (s *ser) SerializeFloat64(v float64) (err error) {
	return s.ser.EncodeFloat64(v)
}

func (s *ser) SerializeComplex64(v complex64) (err error) {
	panic("implement me")
}

func (s *ser) SerializeComplex128(v complex128) (err error) {
	panic("implement me")
}

func (s *ser) SerializeString(v string) (err error) {
	return s.ser.EncodeString(v)
}

func (s *ser) SerializeBytes(v []byte) (err error) {
	return s.ser.EncodeBytes(v)
}

func (s *ser) SerializeSlice(length int) (ss serde.SliceSerializer, err error) {
	err = s.ser.EncodeArrayLen(length)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *ser) SerializeMap(length int) (ms serde.MapSerializer, err error) {
	err = s.ser.EncodeMapLen(length)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *ser) SerializeStruct(name string, length int) (ss serde.StructSerializer, err error) {
	err = s.ser.EncodeMapLen(length)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *ser) SerializeElement(v serde.Serializable) (err error) {
	return v.Serialize(s)
}

func (s *ser) EndSlice() (err error) {
	return nil
}

func (s *ser) SerializeEntry(k, v serde.Serializable) (err error) {
	err = k.Serialize(s)
	if err != nil {
		return
	}
	return v.Serialize(s)
}

func (s *ser) EndMap() (err error) {
	return nil
}

func (s *ser) SerializeField(k, v serde.Serializable) (err error) {
	err = k.Serialize(s)
	if err != nil {
		return
	}
	return v.Serialize(s)
}

func (s *ser) EndStruct() (err error) {
	return nil
}
