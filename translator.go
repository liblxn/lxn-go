package lxn

import (
	"io"
	"math"

	msgpack "github.com/mprot/msgpack-go"

	"github.com/liblxn/lxn-go/internal"
)

const noCurrency = string(currencyPlaceholder)

// Translator represents the function type for translating messages with
// the given context.
type Translator func(key string, ctx Context) string

// ReadCatalog reads a catalog from the given binary stream and returns
// the corresponding translation function.
func ReadCatalog(r io.Reader) (Translator, error) {
	var cat internal.Catalog
	if err := msgpack.Decode(r, &cat); err != nil {
		return nil, err
	}

	msgs := make(map[string]internal.Message, len(cat.Messages)) // key => message
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

func formatMsg(w *writer, m *internal.Message, ctx Context, loc *internal.Locale) {
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

func replace(w *writer, r *internal.Replacement, ctx Context, loc *internal.Locale) {
	v, has := ctx[r.Key]
	if !has {
		w.MissingVar(r.Key)
		return
	}

	switch r.Type {
	case internal.StringReplacement:
		w.WriteString(v.String())

	case internal.NumberReplacement:
		replaceNumber(w, v, r.Key, &loc.DecimalFormat, noCurrency)

	case internal.PercentReplacement:
		replaceNumber(w, v, r.Key, &loc.PercentFormat, noCurrency)

	case internal.MoneyReplacement:
		details, ok := r.Details.Value.(internal.MoneyDetails)
		if !ok {
			w.Corrupted(r.Key)
		} else if curr, has := ctx[details.Currency]; has {
			replaceNumber(w, v, r.Key, &loc.MoneyFormat, curr.String())
		} else {
			w.MissingVar(details.Currency)
		}

	case internal.PluralReplacement:
		details, ok := r.Details.Value.(internal.PluralDetails)
		if !ok {
			w.Corrupted(r.Key)
		} else {
			replacePlural(w, v, ctx, &details, loc)
		}

	case internal.SelectReplacement:
		details, ok := r.Details.Value.(internal.SelectDetails)
		if !ok {
			w.Corrupted(r.Key)
		} else {
			replaceSelect(w, v, ctx, &details, loc)
		}

	default:
		w.UnsupportedReplType(r.Type)
	}
}

func replaceNumber(w *writer, v Variable, key string, nf *internal.NumberFormat, currency string) {
	if num, isNum := v.(number); isNum {
		num.format(w, nf, currency)
	} else {
		w.InvalidType(key)
	}
}

func replacePlural(w *writer, v Variable, ctx Context, details *internal.PluralDetails, loc *internal.Locale) {
	tag := internal.Other
	if num, isNum := v.(number); isNum {
		if i, ok := intval(num); ok {
			if msg, has := details.Custom[i]; has {
				formatMsg(w, &msg, ctx, loc)
				return
			}
		}
		plurals := loc.CardinalPlurals
		if details.Type == internal.Ordinal {
			plurals = loc.OrdinalPlurals
		}
		tag = pluralTag(num, &loc.DecimalFormat, plurals)
	}

	msg, has := details.Variants[tag]
	if !has {
		if msg, has = details.Variants[internal.Other]; !has {
			return
		}
	}
	formatMsg(w, &msg, ctx, loc)
}

func replaceSelect(w *writer, v Variable, ctx Context, details *internal.SelectDetails, loc *internal.Locale) {
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
