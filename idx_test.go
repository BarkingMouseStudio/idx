package idx

import (
	"testing"
)

func TestIdxNewReader(t *testing.T) {
}

func TestIdxRead(t *testing.T) {
}

func TestIdxReadAll(t *testing.T) {
}

func TestIdxNewWriter(t *testing.T) {
}

func TestIdxWrite(t *testing.T) {
}

func TestIdxWriteAll(t *testing.T) {
}

func ExampleReader(t *testing.T) {
  f = os.Open('some_path')
	rd = NewReader(f)
  ch := make(chan, rd.Header.Count)
  go func() {
    for i := 0; i < rd.Header.Count; i++ {
      el, err := rd.Read()
      ch <- el
    }
    close(ch)
    f.Close()
  }()
}
