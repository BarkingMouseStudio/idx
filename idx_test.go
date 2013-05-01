package idx

import (
	"bytes"
	"testing"
)

func TestNewReader(t *testing.T) {
	sample := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x08, 0x03,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x01, 0x02, 0x03, 0x04,
	})
	rd, err := NewReader(sample)
	if err != nil {
		t.Fatal(err)
	}
	if rd.Count != 1 {
		t.Fatal("Got unexpect count")
	}
	if rd.Dimensions[0] != 1 || rd.Dimensions[1] != 2 || rd.Dimensions[2] != 2 {
		t.Fatal("Got unexpect dimensions")
	}
}

func TestRead(t *testing.T) {
	sample := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x08, 0x03,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x01, 0x02, 0x03, 0x04,
	})
	rd, err := NewReader(sample)
	if err != nil {
		t.Fatal(err)
	}
	el, err := rd.Read()
	if err != nil {
		t.Fatal(err)
	}
	if el[0] != 1 || el[1] != 2 || el[2] != 3 || el[3] != 4 {
		t.Fatal("Got unexpected results", el)
	}
}

func TestReadUint8(t *testing.T) {
	sample := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x08, 0x03,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x01, 0x02, 0x03, 0x04,
	})
	rd, err := NewReader(sample)
	if err != nil {
		t.Fatal(err)
	}
	el, err := rd.ReadUint8()
	if err != nil {
		t.Fatal(err)
	}
	if el[0] != 1 || el[1] != 2 || el[2] != 3 || el[3] != 4 {
		t.Fatal("Got unexpected results", el)
	}
}

func TestReadInt8(t *testing.T) {
	sample := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x09, 0x03,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x01, 0x02, 0x03, 0x04,
	})
	rd, err := NewReader(sample)
	if err != nil {
		t.Fatal(err)
	}
	el, err := rd.ReadInt8()
	if err != nil {
		t.Fatal(err)
	}
	if el[0] != 1 || el[1] != 2 || el[2] != 3 || el[3] != 4 {
		t.Fatal("Got unexpected results", el)
	}
}

func TestReadInt16(t *testing.T) {
	sample := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x0B, 0x03,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x01, 0x00, 0x02,
		0x00, 0x03, 0x00, 0x04,
	})
	rd, err := NewReader(sample)
	if err != nil {
		t.Fatal(err)
	}
	el, err := rd.ReadInt16()
	if err != nil {
		t.Fatal(err)
	}
	if el[0] != 1 || el[1] != 2 || el[2] != 3 || el[3] != 4 {
		t.Fatal("Got unexpected results", el)
	}
}

func TestReadInt32(t *testing.T) {
	sample := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x0C, 0x03,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00, 0x04,
	})
	rd, err := NewReader(sample)
	if err != nil {
		t.Fatal(err)
	}
	el, err := rd.ReadInt32()
	if err != nil {
		t.Fatal(err)
	}
	if el[0] != 1 || el[1] != 2 || el[2] != 3 || el[3] != 4 {
		t.Fatal("Got unexpected results", el)
	}
}

func TestReadFloat32(t *testing.T) {
	sample := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x0D, 0x03,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x3f, 0x80, 0x00, 0x00,
		0x40, 0x00, 0x00, 0x00,
		0x40, 0x40, 0x00, 0x00,
		0x40, 0x80, 0x00, 0x00,
	})
	rd, err := NewReader(sample)
	if err != nil {
		t.Fatal(err)
	}
	el, err := rd.ReadFloat32()
	if err != nil {
		t.Fatal(err)
	}
	if el[0] != 1 || el[1] != 2 || el[2] != 3 || el[3] != 4 {
		t.Fatal("Got unexpected results", el)
	}
}

func TestReadFloat64(t *testing.T) {
	sample := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x0E, 0x03,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x3f, 0xf0, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x40, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x40, 0x08, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x40, 0x10, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	})
	rd, err := NewReader(sample)
	if err != nil {
		t.Fatal(err)
	}
	el, err := rd.ReadFloat64()
	if err != nil {
		t.Fatal(err)
	}
	if el[0] != 1 || el[1] != 2 || el[2] != 3 || el[3] != 4 {
		t.Fatal("Got unexpected results", el)
	}
}
