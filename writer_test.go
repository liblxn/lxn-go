package lxn

import (
	"strconv"
	"testing"

	schema "github.com/liblxn/lxn/schema/golang"
)

func TestWriterWriteRunes(t *testing.T) {
	runes := []rune{
		'a', 'ä', '¤',
	}

	var w writer
	w.WriteRunes(runes)
	if s := w.String(); s != string(runes) {
		t.Errorf("unexpected runes: %q", s)
	}
}

func TestWriterWriteAffix(t *testing.T) {
	tests := []struct {
		affix    string
		symbols  schema.Symbols
		currency string
		expected string
	}{
		{
			affix:    "foobar",
			expected: "foobar",
		},
		{
			affix:    "foo_" + string(currencyPlaceholder) + "_bar",
			currency: "$$",
			expected: "foo_$$_bar",
		},
		{
			affix:    "foo_" + string(minusPlaceholder) + "_bar",
			symbols:  schema.Symbols{Minus: "minus"},
			expected: "foo_minus_bar",
		},
		{
			affix:    "foo_" + string(percentPlaceholder) + "_bar",
			symbols:  schema.Symbols{Percent: "percent"},
			expected: "foo_percent_bar",
		},
		{
			affix: string(currencyPlaceholder) + "_foo_" + string(minusPlaceholder) + "_bar_" + string(percentPlaceholder),
			symbols: schema.Symbols{
				Minus:   "minus",
				Percent: "percent",
			},
			currency: "$$",
			expected: "$$_foo_minus_bar_percent",
		},
	}

	for _, test := range tests {
		var w writer
		w.WriteAffix(test.affix, &test.symbols, test.currency)
		if s := w.String(); s != test.expected {
			t.Errorf("unexpected affix for %q: %s", test.affix, s)
		}
	}
}

func TestWriterWriteInt(t *testing.T) {
	tests := []struct {
		digits   []rune
		nf       schema.NumberFormat
		expected string
	}{
		{
			digits: []rune{'1', '2', '3'},
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{Group: "#"},
			},
			expected: "123",
		},
		{
			digits: []rune{'1', '2', '3'},
			nf: schema.NumberFormat{
				Symbols:                  schema.Symbols{Group: "#"},
				PrimaryIntegerGrouping:   2,
				SecondaryIntegerGrouping: 2,
			},
			expected: "1#23",
		},
		{
			digits: []rune{'1', '2', '3', '4'},
			nf: schema.NumberFormat{
				Symbols:                  schema.Symbols{Group: "#"},
				PrimaryIntegerGrouping:   2,
				SecondaryIntegerGrouping: 2,
			},
			expected: "12#34",
		},
		{
			digits: []rune{'1', '2', '3', '4', '5', '6'},
			nf: schema.NumberFormat{
				Symbols:                  schema.Symbols{Group: "#"},
				PrimaryIntegerGrouping:   3,
				SecondaryIntegerGrouping: 2,
			},
			expected: "1#23#456",
		},
		{
			digits: []rune{'1', '2', '3', '4', '5', '6', '7'},
			nf: schema.NumberFormat{
				Symbols:                  schema.Symbols{Group: "#"},
				PrimaryIntegerGrouping:   3,
				SecondaryIntegerGrouping: 2,
			},
			expected: "12#34#567",
		},
	}

	for _, test := range tests {
		var w writer
		w.WriteInt(test.digits, &test.nf)
		if s := w.String(); s != test.expected {
			t.Errorf("unexpected integer for %q: %s", string(test.expected), s)
		}
	}
}

func TestWriterWriteFrac(t *testing.T) {
	tests := []struct {
		digits   []rune
		nf       schema.NumberFormat
		expected string
	}{
		{
			digits:   nil,
			expected: "",
		},
		{
			digits: []rune{'1', '2', '3'},
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Decimal: "_",
					Group:   "#",
				},
			},
			expected: "_123",
		},
		{
			digits: []rune{'1', '2', '3'},
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Decimal: "_",
					Group:   "#",
				},
				FractionGrouping: 2,
			},
			expected: "_12#3",
		},
		{
			digits: []rune{'1', '2', '3', '4'},
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Decimal: "_",
					Group:   "#",
				},
				FractionGrouping: 2,
			},
			expected: "_12#34",
		},
	}

	for _, test := range tests {
		var w writer
		w.WriteFrac(test.digits, &test.nf)
		if s := w.String(); s != test.expected {
			t.Errorf("unexpected fraction for %q: %s", string(test.expected), s)
		}
	}
}

func TestWriterMissingVar(t *testing.T) {
	var w writer
	w.MissingVar("varkey")
	if s := w.String(); s != "%!(MISSING:varkey)" {
		t.Errorf("unexpected message for missing variable: %s", s)
	}
}

func TestWriterInvalidType(t *testing.T) {
	var w writer
	w.InvalidType("varkey")
	if s := w.String(); s != "%!(INVALID:varkey)" {
		t.Errorf("unexpected message for invalid type: %s", s)
	}
}

func TestWriterUnsupportedReplType(t *testing.T) {
	const typ = schema.SelectReplacement
	typstr := strconv.FormatInt(int64(typ), 10)

	var w writer
	w.UnsupportedReplType(typ)
	if s := w.String(); s != "%!(UNSUPPORTED:ReplType-"+typstr+")" {
		t.Errorf("unexpected message for invalid type: %s", s)
	}
}
