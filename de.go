package msgpack

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/Xuanwo/serde-go"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
)

func DeserializeFromReader(r io.Reader, v serde.Deserializable) error {
	mde := msgpack.GetDecoder()
	defer msgpack.PutDecoder(mde)

	de := de{de: mde}

	if br, ok := r.(bufReader); ok {
		de.br = br
	} else {
		de.br = bufio.NewReader(r)
	}

	mde.Reset(de.br)

	return v.Deserialize(&de)
}

func DeserializeFromBytes(s []byte, v serde.Deserializable) error {
	return DeserializeFromReader(bytes.NewReader(s), v)
}

type bufReader interface {
	io.Reader
	io.ByteReader
}

type de struct {
	br bufReader
	de *msgpack.Decoder
}

func (d *de) DeserializeAny(v serde.Visitor) (err error) {
	code, err := d.de.PeekCode()
	if err != nil {
		return
	}

	if msgpcode.IsFixedNum(code) {
		// Make sure this code has been consumed.
		_, _ = d.br.ReadByte()
		return v.VisitInt8(int8(code))
	}
	if msgpcode.IsFixedMap(code) {
		return d.DeserializeMap(v)
	}
	if msgpcode.IsFixedArray(code) {
		return d.DeserializeSlice(v)
	}
	if msgpcode.IsFixedString(code) {
		return d.DeserializeString(v)
	}

	switch code {
	case msgpcode.Nil:
		err = d.DeserializeNil(v)
	case msgpcode.False, msgpcode.True:
		err = d.DeserializeBool(v)
	case msgpcode.Float:
		err = d.DeserializeFloat32(v)
	case msgpcode.Double:
		err = d.DeserializeFloat64(v)
	case msgpcode.Uint8:
		err = d.DeserializeUint8(v)
	case msgpcode.Uint16:
		err = d.DeserializeUint16(v)
	case msgpcode.Uint32:
		err = d.DeserializeUint32(v)
	case msgpcode.Uint64:
		err = d.DeserializeUint64(v)
	case msgpcode.Int8:
		err = d.DeserializeInt8(v)
	case msgpcode.Int16:
		err = d.DeserializeInt16(v)
	case msgpcode.Int32:
		err = d.DeserializeInt32(v)
	case msgpcode.Int64:
		err = d.DeserializeInt64(v)
	case msgpcode.Bin8, msgpcode.Bin16, msgpcode.Bin32:
		err = d.DeserializeBytes(v)
	case msgpcode.Str8, msgpcode.Str16, msgpcode.Str32:
		err = d.DeserializeString(v)
	case msgpcode.Array16, msgpcode.Array32:
		err = d.DeserializeSlice(v)
	case msgpcode.Map16, msgpcode.Map32:
		err = d.DeserializeMap(v)
	default:
		err = fmt.Errorf("not supported msgpcode: %v", code)
	}

	return err
}

func (d *de) DeserializeNil(v serde.Visitor) (err error) {
	err = d.de.DecodeNil()
	if err != nil {
		return
	}
	return v.VisitNil()
}

func (d *de) DeserializeBool(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeBool()
	if err != nil {
		return err
	}
	return v.VisitBool(vv)
}

func (d *de) DeserializeInt(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeInt()
	if err != nil {
		return err
	}
	return v.VisitInt(vv)
}

func (d *de) DeserializeInt8(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeInt8()
	if err != nil {
		return err
	}
	return v.VisitInt8(vv)
}

func (d *de) DeserializeInt16(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeInt16()
	if err != nil {
		return err
	}
	return v.VisitInt16(vv)
}

func (d *de) DeserializeInt32(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeInt32()
	if err != nil {
		return err
	}
	return v.VisitInt32(vv)
}

func (d *de) DeserializeInt64(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeInt64()
	if err != nil {
		return err
	}
	return v.VisitInt64(vv)
}

func (d *de) DeserializeUint(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeUint()
	if err != nil {
		return err
	}
	return v.VisitUint(vv)
}

func (d *de) DeserializeUint8(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeUint8()
	if err != nil {
		return err
	}
	return v.VisitUint8(vv)
}

func (d *de) DeserializeUint16(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeUint16()
	if err != nil {
		return err
	}
	return v.VisitUint16(vv)
}

func (d *de) DeserializeUint32(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeUint32()
	if err != nil {
		return err
	}
	return v.VisitUint32(vv)
}

func (d *de) DeserializeUint64(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeUint64()
	if err != nil {
		return err
	}
	return v.VisitUint64(vv)
}

func (d *de) DeserializeFloat32(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeFloat32()
	if err != nil {
		return err
	}
	return v.VisitFloat32(vv)
}

func (d *de) DeserializeFloat64(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeFloat64()
	if err != nil {
		return err
	}
	return v.VisitFloat64(vv)
}

func (d *de) DeserializeComplex64(v serde.Visitor) (err error) {
	panic("implement me")
}

func (d *de) DeserializeComplex128(v serde.Visitor) (err error) {
	panic("implement me")
}

func (d *de) DeserializeString(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeString()
	if err != nil {
		return err
	}
	return v.VisitString(vv)
}

func (d *de) DeserializeBytes(v serde.Visitor) (err error) {
	vv, err := d.de.DecodeBytes()
	if err != nil {
		return err
	}
	return v.VisitBytes(vv)
}

func (d *de) DeserializeSlice(v serde.Visitor) (err error) {
	length, err := d.de.DecodeArrayLen()
	if err != nil {
		return err
	}

	return v.VisitSlice(&containerAccess{
		d:      d,
		length: length,
	})
}

func (d *de) DeserializeMap(v serde.Visitor) (err error) {
	length, err := d.de.DecodeMapLen()
	if err != nil {
		return err
	}

	return v.VisitMap(&containerAccess{
		d:      d,
		length: length,
	})
}

func (d *de) DeserializeStruct(name string, fields []string, v serde.Visitor) (err error) {
	return d.DeserializeMap(v)
}

type containerAccess struct {
	d *de

	length int
	idx    int
}

func (ca *containerAccess) NextElement(v serde.Visitor) (ok bool, err error) {
	if ca.idx >= ca.length {
		return false, nil
	}

	err = ca.d.DeserializeAny(v)
	if err != nil {
		return
	}
	ca.idx += 1
	return true, nil
}

func (ca *containerAccess) NextKey(v serde.Visitor) (ok bool, err error) {
	if ca.idx >= ca.length {
		return false, nil
	}

	err = ca.d.DeserializeAny(v)
	if err != nil {
		return
	}
	return true, nil
}

func (ca *containerAccess) NextValue(v serde.Visitor) (err error) {
	defer func() {
		ca.idx += 1
	}()

	if ca.idx >= ca.length {
		return nil
	}
	return ca.d.DeserializeAny(v)
}
