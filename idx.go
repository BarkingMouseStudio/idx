package idx

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
)

var FormatError = errors.New("Incorrect format")
var IndexError = errors.New("Index out of range")
var SizeError = errors.New("Dimensions too small")
var TypeError = errors.New("Type mismatch")

func sliceSize(dims []int32) (size int) {
	size = 1
	if len(dims) < 2 {
		return
	}
	for _, x := range dims[1:] {
		size *= int(x)
	}
	return
}

var dataTypes = map[int8]int{
	0x08: 1, // uint8
	0x09: 1, // int8
	0x0B: 2, // int16
	0x0C: 4, // int32
	0x0D: 4, // float32
	0x0E: 8, // float64
}

type Reader struct {
	DataType   int8
	Count      int32
	Dimensions []int32
	reader     *bufio.Reader
	size       int
	index      int32
}

func NewReader(r io.Reader) (rr *Reader, err error) {
	var zeros int16
	err = binary.Read(r, binary.BigEndian, &zeros)
	if err != nil {
		return
	}
	if zeros != 0 {
		err = FormatError
		return
	}

	var dataType int8
	err = binary.Read(r, binary.BigEndian, &dataType)
	if err != nil {
		return
	}
	_, isValid := dataTypes[dataType]
	if !isValid {
		err = FormatError
		return
	}

	var numDimensions int8
	err = binary.Read(r, binary.BigEndian, &numDimensions)
	if err != nil {
		return
	}

	dimensions := make([]int32, numDimensions)
	err = binary.Read(r, binary.BigEndian, &dimensions)
	if err != nil {
		return
	}

	if numDimensions < 1 || int(numDimensions) != len(dimensions) {
		err = FormatError
		return
	}

	count := dimensions[0]
	size := sliceSize(dimensions)
	br := bufio.NewReaderSize(r, size)
	rr = &Reader{dataType, count, dimensions, br, size, 0}
	return
}

func (rr *Reader) Read() (el []byte, err error) {
	if rr.index >= rr.Count {
		err = IndexError
		return
	}
	rr.index++
	el = make([]byte, rr.size*dataTypes[rr.DataType])
	rr.reader.Read(el)
	return
}

func (rr *Reader) ReadUint8() (el []uint8, err error) {
	if rr.index >= rr.Count {
		err = IndexError
		return
	}
	rr.index++
	el = make([]uint8, rr.size)
	binary.Read(rr.reader, binary.BigEndian, el)
	return
}

func (rr *Reader) ReadInt8() (el []int8, err error) {
	if rr.index >= rr.Count {
		err = IndexError
		return
	}
	rr.index++
	el = make([]int8, rr.size)
	binary.Read(rr.reader, binary.BigEndian, el)
	return
}

func (rr *Reader) ReadInt16() (el []int16, err error) {
	if rr.index >= rr.Count {
		err = IndexError
		return
	}
	rr.index++
	el = make([]int16, rr.size)
	binary.Read(rr.reader, binary.BigEndian, el)
	return
}

func (rr *Reader) ReadInt32() (el []int32, err error) {
	if rr.index >= rr.Count {
		err = IndexError
		return
	}
	rr.index++
	el = make([]int32, rr.size)
	binary.Read(rr.reader, binary.BigEndian, el)
	return
}

func (rr *Reader) ReadFloat32() (el []float32, err error) {
	if rr.index >= rr.Count {
		err = IndexError
		return
	}
	rr.index++
	el = make([]float32, rr.size)
	binary.Read(rr.reader, binary.BigEndian, el)
	return
}

func (rr *Reader) ReadFloat64() (el []float64, err error) {
	if rr.index >= rr.Count {
		err = IndexError
		return
	}
	rr.index++
	el = make([]float64, rr.size)
	binary.Read(rr.reader, binary.BigEndian, el)
	return
}

/* TODO:
type Writer struct {
	Dimensions []int32
	writer     *bufio.Writer
	index      int32
	dataType   int8
}

func NewWriter(w io.Writer, dataType int8, dimensions []int32) (ww *Writer, err error) {
	if len(dimensions) < 1 {
		err = SizeError
		return
	}
	bw := bufio.NewWriterSize(w, sliceSize(dimensions))
	ww = &Writer{dimensions, bw, 0, dataType}
	return
}

func (ww *Writer) Write(el []byte) (err error) {
	if ww.index >= ww.Dimensions[0] {
		err = IndexError
		return
	}
	ww.index++
	err = binary.Write(ww.writer, binary.BigEndian, el)
	return
}

func (ww *Writer) Flush() {
	ww.writer.Flush()
} */
