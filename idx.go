package idx

import (
	"encoding/binary"
	"errors"
	"github.com/Freeflow/matrix"
	"os"
)

/*
[offset] [type]          [value]          [description] 
0000     32 bit integer  0x00000803(2051) magic number 
0004     32 bit integer  60000            number of images 
0008     32 bit integer  28               number of rows 
0012     32 bit integer  28               number of columns 
0016     unsigned byte   ??               pixel 
0017     unsigned byte   ??               pixel 
........ 
xxxx     unsigned byte   ??               pixel
*/

const (
	magicNumber int32 = 2051
)

type Header struct {
	Magic, Count, Rows, Cols int32
}

var FormatError = errors.New("Failed to read magic number: incorrect format")
var IndexError = errors.New("Index out of range")

type Reader struct {
	Header *Header
	reader *bufio.Reader
	index  int64
}

func NewReader(r io.Reader) (rd *Reader, err error) {
	h := new(Header)
	err = binary.Read(r, binary.BigEndian, &h)
	if err != nil {
		return
	}
	if h.Magic != magicNumber {
		err = new(FormatError)
		return
	}
	rd = &Reader{bufio.NewReaderSize(r, h.Rows*h.Cols), h}
	return
}

func Read(rd *Reader) (el []uint8, err error) {
	rd.index++
	h := rd.Header
	el = make([]uint8, h.Rows*h.Cols)
	err = binary.Read(rd.reader, binary.BigEndian, &el)
	return
}

func ReadAll(rd *Reader) (els [][]uint8, err error) {
	for i := rd.index; i < rd.Header.Count; i++ {
		el, err = rd.Read()
		if err != nil {
			return
		}
		els = append(els, el)
	}
	return
}

type Writer struct {
	writer *bufio.Writer
	Header *Header
}

func NewWriter(w io.Writer, count, rows, cols int32) *Writer {
	h := &Header{magicNumber, count, rows, cols}
	return &Writer{bufio.NewWriterSize(w, h.Rows*h.Cols), h}
}

func (w *Writer) Flush() {
	w.writer.Flush()
}

func (w *Writer) Write(el []uint8) (err error) {
	err = binary.Write(w.writer, binary.BigEndian, el)
	return
}

func (w *Writer) WriteAll(els [][]uint8) (err error) {
	if len(els) != w.Header.Count {
		err = new(IndexError)
		return
	}
	var el []uint8
	for i := 0; i < w.Header.Count; i++ {
		el = els[i]
		w.Write(el)
	}
	return
}
