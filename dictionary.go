package lxn

import (
	"io"

	"github.com/mprot/msgpack-go"

	"github.com/liblxn/lxn-go/internal/lxn"
)

type Dictionary struct {
	loc *Locale
	cat *Catalog
}

func ReadDictionary(r io.Reader) (*Dictionary, error) {
	dic := &lxn.Dictionary{}
	if err := msgpack.Decode(r, dic); err != nil {
		return nil, err
	}

	return &Dictionary{
		loc: newLocale(dic.Locale),
		cat: newCatalog(dic.Locale.ID, dic.Messages),
	}, nil
}

func (d *Dictionary) Translate(section string, messageKey string, ctx Context) string {
	msg := d.cat.Message(section, messageKey)
	if msg == nil {
		return ""
	}
	return msg.Format(d.loc, ctx)
}
