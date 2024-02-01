package lxn

import (
	"io"

	"github.com/mprot/msgpack-go"

	"github.com/liblxn/lxn-go/internal/lxn"
)

type Locale struct {
	loc lxn.Locale
}

// ReadLocale reads the locale information from the given binary
// stream.
func ReadLocale(r io.Reader) (*Locale, error) {
	var loc lxn.Locale
	if err := msgpack.Decode(r, &loc); err != nil {
		return nil, err
	}
	return newLocale(loc), nil
}

func newLocale(loc lxn.Locale) *Locale {
	return &Locale{loc: loc}
}

func (l *Locale) ID() string {
	return l.loc.ID
}
