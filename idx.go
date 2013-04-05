package idx

import (
	"encoding/binary"
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

type header struct {
	Magic, Count, Rows, Cols int32
}

type idxError int

func (err *idxError) Error() string {
	return "Failed to read magic number: file not in correct format"
}

func (err *idxError) String() string {
	return err.Error()
}

func Read(fp string, ch chan<- *matrix.Matrix) (err error) {
	f, err := os.Open(fp)
	if err != nil {
		return
	}

	var h header
	err = binary.Read(f, binary.BigEndian, &h)
	if err != nil {
		return
	}

	if h.Magic != 2051 {
		err = new(idxError)
		return
	}

  count := int(h.Count)
  count = 100
  rows := int(h.Rows)
  cols := int(h.Cols)
  length := rows * cols
	slice := make([]uint8, length)

	go func() {
		for i := 0; i < count; i++ {
			err = binary.Read(f, binary.BigEndian, &slice)
			if err != nil {
				return
			}
			elements := make([]float64, length)
			for j, x := range slice {
				elements[j] = float64(x) / 255
			}
			var m *matrix.Matrix
			m, err = matrix.New(elements, rows, cols)
			if err != nil {
				return
			}
			ch <- m
		}
		close(ch)
		f.Close()
	}()

	return
}
