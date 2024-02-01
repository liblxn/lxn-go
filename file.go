package lxn

import (
	"io"
	"os"
)

func FromFile[T any](read func(io.Reader) (*T, error), filename string) (*T, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return read(f)

}
