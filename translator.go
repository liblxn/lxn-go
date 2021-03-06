package lxn

import (
	"io"
	"math"
	"os"

	schema "github.com/liblxn/lxn/schema/golang"
	msgpack "github.com/mprot/msgpack-go"
)

const noCurrency = string(currencyPlaceholder)

// Translator represents the function type for translating messages with
// the given context.
type Translator func(key string, ctx Context) string

// ReadCatalog reads a catalog from the given binary stream and returns
// the corresponding translation function.
func ReadCatalog(r io.Reader) (Translator, error) {
	var cat schema.Catalog
	if err := msgpack.Decode(r, &cat); err != nil {
		return nil, err
	}

	msgs := make(map[string]schema.Message, len(cat.Messages)) // key => message
	for _, m := range cat.Messages {
		key := m.Key
		if m.Section != "" {
			key = m.Section + "." + key
		}
		msgs[key] = m
	}

	return func(key string, ctx Context) string {
		m, has := msgs[key]
		if !has {
			return ""
		}

		var w writer
		formatMsg(&w, &m, ctx, &cat.Locale)
		return w.String()
	}, nil
}

func ReadCatalogFile(filename string) (Translator, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadCatalog(f)
}

func formatMsg(w *writer, m *schema.Message, ctx Context, loc *schema.Locale) {
	off := 0
	for i, t := range m.Text {
		for off < len(m.Replacements) && m.Replacements[off].TextPos <= i {
			replace(w, &m.Replacements[off], ctx, loc)
			off++
		}
		w.WriteString(t)
	}
	for _, r := range m.Replacements[off:] {
		replace(w, &r, ctx, loc)
	}
}

func replace(w *writer, r *schema.Replacement, ctx Context, loc *schema.Locale) {
	v, has := ctx[r.Key]
	if !has {
		w.MissingVar(r.Key)
		return
	}

	switch r.Type {
	case schema.StringReplacement:
		w.WriteString(v.String())

	case schema.NumberReplacement:
		replaceNumber(w, v, r.Key, &loc.DecimalFormat, noCurrency)

	case schema.PercentReplacement:
		replaceNumber(w, v, r.Key, &loc.PercentFormat, noCurrency)

	case schema.MoneyReplacement:
		details, ok := r.Details.Value.(schema.MoneyDetails)
		if !ok {
			w.Corrupted(r.Key)
		} else if curr, has := ctx[details.Currency]; has {
			replaceNumber(w, v, r.Key, &loc.MoneyFormat, curr.String())
		} else {
			w.MissingVar(details.Currency)
		}

	case schema.PluralReplacement:
		details, ok := r.Details.Value.(schema.PluralDetails)
		if !ok {
			w.Corrupted(r.Key)
		} else {
			replacePlural(w, v, ctx, &details, loc)
		}

	case schema.SelectReplacement:
		details, ok := r.Details.Value.(schema.SelectDetails)
		if !ok {
			w.Corrupted(r.Key)
		} else {
			replaceSelect(w, v, ctx, &details, loc)
		}

	default:
		w.UnsupportedReplType(r.Type)
	}
}

func replaceNumber(w *writer, v Variable, key string, nf *schema.NumberFormat, currency string) {
	if num, isNum := v.(number); isNum {
		num.format(w, nf, currency)
	} else {
		w.InvalidType(key)
	}
}

func replacePlural(w *writer, v Variable, ctx Context, details *schema.PluralDetails, loc *schema.Locale) {
	tag := schema.Other
	if num, isNum := v.(number); isNum {
		if i, ok := intval(num); ok {
			if msg, has := details.Custom[i]; has {
				formatMsg(w, &msg, ctx, loc)
				return
			}
		}
		plurals := loc.CardinalPlurals
		if details.Type == schema.Ordinal {
			plurals = loc.OrdinalPlurals
		}
		tag = pluralTag(num, &loc.DecimalFormat, plurals)
	}

	msg, has := details.Variants[tag]
	if !has {
		if msg, has = details.Variants[schema.Other]; !has {
			return
		}
	}
	formatMsg(w, &msg, ctx, loc)
}

func replaceSelect(w *writer, v Variable, ctx Context, details *schema.SelectDetails, loc *schema.Locale) {
	msg, has := details.Cases[v.String()]
	if !has {
		if msg, has = details.Cases[details.Fallback]; !has {
			return
		}
	}
	formatMsg(w, &msg, ctx, loc)
}

func intval(num number) (int64, bool) {
	switch num := num.(type) {
	case Int:
		return int64(num), true
	case Uint:
		return int64(num), uint64(num) <= math.MaxInt64
	}
	return 0, false
}
