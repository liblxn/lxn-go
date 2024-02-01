package lxn

import (
	"fmt"
	"io"

	"github.com/mprot/msgpack-go"

	"github.com/liblxn/lxn-go/internal/lxn"
)

type Catalog struct {
	localeID string
	msgs     map[string]*Message // message's unique key => message
}

// ReadCatalog reads a catalog from the given binary stream.
func ReadCatalog(r io.Reader) (*Catalog, error) {
	cat := &lxn.Catalog{}
	if err := msgpack.Decode(r, cat); err != nil {
		return nil, err
	}
	return newCatalog(cat.LocaleID, cat.Messages), nil
}

func newCatalog(localeID string, messages []lxn.Message) *Catalog {
	msgs := make(map[string]*Message, len(messages))
	for _, m := range messages {
		msg := newMessage(m)
		msgs[uniqMessageKey(msg.Section(), msg.Key())] = msg
	}

	return &Catalog{
		localeID: localeID,
		msgs:     msgs,
	}
}

// MergeCatalogs merges all catalogs into a single one. All given
// catalogs must have the same locale id. If there are duplicates
// in the message key, they will be overwritten.
//
// Note: This function panics if no catalogs are specified.
func MergeCatalogs(catalogs ...*Catalog) (*Catalog, error) {
	if len(catalogs) == 0 {
		panic("no catalogs to merge")
	}

	localeID := catalogs[0].localeID
	msgs := map[string]*Message{}
	for _, cat := range catalogs {
		if localeID != cat.localeID {
			return nil, fmt.Errorf("multiple locale ids detected: %s and %s", localeID, cat.localeID)
		}
		for key, msg := range cat.msgs {
			msgs[key] = msg
		}
	}

	return &Catalog{
		localeID: localeID,
		msgs:     msgs,
	}, nil
}

func (c *Catalog) LocaleID() string {
	return c.localeID
}

// Message returns the message for the given section and message key.
// For messages that do not live within a section, the first argument
// needs to be empty. If the message with the given key cannot be found,
// nil will be returned.
func (c *Catalog) Message(section string, key string) *Message {
	return c.msgs[uniqMessageKey(section, key)]
}

func uniqMessageKey(section, key string) string {
	return section + "." + key
}
