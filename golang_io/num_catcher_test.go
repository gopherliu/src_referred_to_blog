package golang_io_test

import (
	"io"
	"testing"

	"src_referred_to_blog/golang_io"
)

func TestNumCatcher_Read(t *testing.T) {
		reader := golang_io.NewNumCatcher("123W3333e2")
		p := make([]byte, 4)
		i := 0
		for {
			i++
			n, err := reader.Read(p)
			if err == io.EOF {
				break
			}
			t.Log(string(p[:n]))
		}
}

