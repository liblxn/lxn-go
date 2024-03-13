package lxn

import (
	"io"

	"github.com/mprot/msgpack-go"

	"github.com/liblxn/lxn-go/internal/lxn"
)

// Dictionary is a container that holds messages for a locale and the
// locale data necessary to format these messages. Each message is
// identified by a unique key which consists of the section and the
// message key within the lxn file.
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

func (d *Dictionary) Locale() *Locale {
	return d.loc
}

func (d *Dictionary) Translate(section string, messageKey string, ctx Context) string {
	msg := d.cat.Message(section, messageKey)
	if msg == nil {
		return ""
	}
	return msg.Format(d.loc, ctx)
}
