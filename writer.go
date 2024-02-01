package lxn

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/liblxn/lxn-go/internal/lxn"
)

const (
	currencyPlaceholder = '¤'
	minusPlaceholder    = '-'
	percentPlaceholder  = '%'
)

type writer struct {
	strings.Builder
}

func (w *writer) WriteRunes(runes []rune) {
	for _, r := range runes {
		w.WriteRune(r)
	}
}

func (w *writer) WriteAffix(affix string, symb *lxn.Symbols, currency string) {
	for affix != "" {
		ch, n := utf8.DecodeRuneInString(affix)
		switch ch {
		case currencyPlaceholder:
			w.WriteString(currency)
		case minusPlaceholder:
			w.WriteString(symb.Minus)
		case percentPlaceholder:
			w.WriteString(symb.Percent)
		default:
			w.WriteRune(ch)
		}
		affix = affix[n:]
	}
}

func (w *writer) WriteInt(digits []rune, nf *lxn.NumberFormat) {
	if nf.PrimaryIntegerGrouping > 0 {
		// secondary groups
		lead := (len(digits) - nf.PrimaryIntegerGrouping) % nf.SecondaryIntegerGrouping
		if lead > 0 {
			w.WriteRunes(digits[:lead])
			w.WriteString(nf.Symbols.Group)
			digits = digits[lead:]
		}

		for len(digits) > nf.PrimaryIntegerGrouping {
			w.WriteRunes(digits[:nf.SecondaryIntegerGrouping])
			w.WriteString(nf.Symbols.Group)
			digits = digits[nf.SecondaryIntegerGrouping:]
		}
	}

	// primary group
	w.WriteRunes(digits)
}

func (w *writer) WriteFrac(digits []rune, nf *lxn.NumberFormat) {
	if len(digits) == 0 {
		return
	}

	w.WriteString(nf.Symbols.Decimal)
	if nf.FractionGrouping > 0 {
		for len(digits) > nf.FractionGrouping {
			w.WriteRunes(digits[:nf.FractionGrouping])
			w.WriteString(nf.Symbols.Group)
			digits = digits[nf.FractionGrouping:]
		}
	}
	w.WriteRunes(digits)
}

func (w *writer) MissingVar(key string) {
	w.WriteString("%!(MISSING:" + key + ")")
}

func (w *writer) InvalidType(key string) {
	w.WriteString("%!(INVALID:" + key + ")")
}

func (w *writer) UnsupportedReplType(typ lxn.ReplacementType) {
	t := strconv.FormatInt(int64(typ), 10)
	w.WriteString("%!(UNSUPPORTED:ReplType-" + t + ")")
}

func (w *writer) Corrupted(key string) {
	w.WriteString("%!(CORRUPTED:" + key + ")")
}
