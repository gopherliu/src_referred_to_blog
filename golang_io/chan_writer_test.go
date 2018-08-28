package golang_io_test

import (
	"testing"

	"src_referred_to_blog/golang_io"
)

func TestChanWriter_Chan(t *testing.T) {
	writer := golang_io.NewChanWriter()
	go func() {
		defer writer.Close()
		writer.Write([]byte("Hello "))
		writer.Write([]byte("World!"))
	}()
	for c := range writer.Chan() {
		t.Logf("%c", c)
	}
}
