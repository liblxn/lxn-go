package lxn

import (
	"math"
	"testing"

	schema "github.com/liblxn/lxn/schema/golang"
)

func TestIntDigits(t *testing.T) {
	tests := []struct {
		val          Int
		nf           schema.NumberFormat
		expectedInt  string
		expectedFrac string
	}{
		{
			val:          0,
			expectedInt:  "0",
			expectedFrac: "",
		},
		{
			val:          -123,
			expectedInt:  "123",
			expectedFrac: "",
		},
		{
			val: -123,
			nf: schema.NumberFormat{
				MinFractionDigits: 2,
			},
			expectedInt:  "123",
			expectedFrac: "00",
		},
		{
			val: 123,
			nf: schema.NumberFormat{
				MinIntegerDigits: 5,
			},
			expectedInt:  "00123",
			expectedFrac: "",
		},
		{
			val: -123,
			nf: schema.NumberFormat{
				MinIntegerDigits:  5,
				MinFractionDigits: 5,
			},
			expectedInt:  "00123",
			expectedFrac: "00000",
		},
	}

	var buf [maxIntDigits]rune
	for _, test := range tests {
		intDigits, fracDigits := test.val.digits(buf[:], &test.nf, '0')
		if string(intDigits) != test.expectedInt {
			t.Errorf("unexpected integer digits: %s", string(intDigits))
		}
		if string(fracDigits) != test.expectedFrac {
			t.Errorf("unexpected fractional digits: %s", string(fracDigits))
		}
	}
}

func TestIntFormat(t *testing.T) {
	tests := []struct {
		val      Int
		nf       schema.NumberFormat
		expected string
	}{
		{
			val: 0,
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Zero:    '0',
					Decimal: ":",
				},
				PositivePrefix:    "p",
				PositiveSuffix:    "s",
				MinFractionDigits: 2,
			},
			expected: "p0:00s",
		},
		{
			val: 123,
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Zero:    '0',
					Decimal: ":",
				},
				PositivePrefix:    "p",
				PositiveSuffix:    "s",
				MinFractionDigits: 1,
				MinIntegerDigits:  4,
			},
			expected: "p0123:0s",
		},
		{
			val: -123,
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Zero:    '0',
					Decimal: ":",
				},
				NegativePrefix:    "np",
				NegativeSuffix:    "ns",
				MinFractionDigits: 2,
			},
			expected: "np123:00ns",
		},
	}

	for _, test := range tests {
		var w writer
		test.val.format(&w, &test.nf, noCurrency)
		if s := w.String(); s != test.expected {
			t.Errorf("unexpected Int format for %q: %s", test.expected, s)
		}
	}
}

func TestUintDigits(t *testing.T) {
	tests := []struct {
		val          Uint
		nf           schema.NumberFormat
		expectedInt  string
		expectedFrac string
	}{
		{
			val:          0,
			expectedInt:  "0",
			expectedFrac: "",
		},
		{
			val:          123,
			expectedInt:  "123",
			expectedFrac: "",
		},
		{
			val: 123,
			nf: schema.NumberFormat{
				MinFractionDigits: 2,
			},
			expectedInt:  "123",
			expectedFrac: "00",
		},
		{
			val: 123,
			nf: schema.NumberFormat{
				MinIntegerDigits: 5,
			},
			expectedInt:  "00123",
			expectedFrac: "",
		},
		{
			val: 123,
			nf: schema.NumberFormat{
				MinIntegerDigits:  5,
				MinFractionDigits: 5,
			},
			expectedInt:  "00123",
			expectedFrac: "00000",
		},
	}

	var buf [maxIntDigits]rune
	for _, test := range tests {
		intDigits, fracDigits := test.val.digits(buf[:], &test.nf, '0')
		if string(intDigits) != test.expectedInt {
			t.Errorf("unexpected integer digits: %s", string(intDigits))
		}
		if string(fracDigits) != test.expectedFrac {
			t.Errorf("unexpected fractional digits: %s", string(fracDigits))
		}
	}
}

func TestUintFormat(t *testing.T) {
	tests := []struct {
		val      Uint
		nf       schema.NumberFormat
		expected string
	}{
		{
			val: 0,
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Zero:    '0',
					Decimal: ":",
				},
				PositivePrefix:    "p",
				PositiveSuffix:    "s",
				MinFractionDigits: 2,
			},
			expected: "p0:00s",
		},
		{
			val: 123,
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Zero:    '0',
					Decimal: ":",
				},
				PositivePrefix:    "p",
				PositiveSuffix:    "s",
				MinFractionDigits: 1,
				MinIntegerDigits:  4,
			},
			expected: "p0123:0s",
		},
	}

	for _, test := range tests {
		var w writer
		test.val.format(&w, &test.nf, noCurrency)
		if s := w.String(); s != test.expected {
			t.Errorf("unexpected Uint format for %q: %s", test.expected, s)
		}
	}
}

func TestFloatDigits(t *testing.T) {
	tests := []struct {
		val          Float
		nf           schema.NumberFormat
		expectedInt  string
		expectedFrac string
	}{
		{
			val:          0,
			expectedInt:  "0",
			expectedFrac: "",
		},
		{
			val:          -123,
			expectedInt:  "123",
			expectedFrac: "",
		},
		{
			val: 123,
			nf: schema.NumberFormat{
				MinFractionDigits: 2,
				MaxFractionDigits: 2,
			},
			expectedInt:  "123",
			expectedFrac: "00",
		},
		{
			val: -123,
			nf: schema.NumberFormat{
				MinIntegerDigits: 5,
			},
			expectedInt:  "00123",
			expectedFrac: "",
		},
		{
			val: 123,
			nf: schema.NumberFormat{
				MinIntegerDigits:  5,
				MinFractionDigits: 3,
				MaxFractionDigits: 5,
			},
			expectedInt:  "00123",
			expectedFrac: "000",
		},
		{
			val: -123.12,
			nf: schema.NumberFormat{
				MinFractionDigits: 1,
				MaxFractionDigits: 1,
			},
			expectedInt:  "123",
			expectedFrac: "1",
		},
		{
			val: 123.12,
			nf: schema.NumberFormat{
				MinFractionDigits: 3,
				MaxFractionDigits: 5,
			},
			expectedInt:  "123",
			expectedFrac: "120",
		},
	}

	var buf [maxFloatDigits]rune
	for _, test := range tests {
		intDigits, fracDigits := test.val.digits(buf[:], &test.nf, '0')
		if string(intDigits) != test.expectedInt {
			t.Errorf("unexpected integer digits: %s", string(intDigits))
		}
		if string(fracDigits) != test.expectedFrac {
			t.Errorf("unexpected fractional digits: %s", string(fracDigits))
		}
	}
}

func TestFloatFormat(t *testing.T) {
	tests := []struct {
		val      Float
		nf       schema.NumberFormat
		expected string
	}{
		{
			val: Float(math.NaN()),
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Nan: "not-a-number",
				},
			},
			expected: "not-a-number",
		},
		{
			val: Float(math.Inf(+1)),
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Inf: "infinity",
				},
				PositivePrefix: "p_",
				PositiveSuffix: "_s",
			},
			expected: "p_infinity_s",
		},
		{
			val: Float(math.Inf(-1)),
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Inf: "infinity",
				},
				NegativePrefix: "np_",
				NegativeSuffix: "_ns",
			},
			expected: "np_infinity_ns",
		},
		{
			val: 0,
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Zero:    '0',
					Decimal: ":",
				},
				PositivePrefix:    "p",
				PositiveSuffix:    "s",
				MinFractionDigits: 2,
				MaxFractionDigits: 5,
			},
			expected: "p0:00s",
		},
		{
			val: 123.12,
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Zero:    '0',
					Decimal: ":",
				},
				PositivePrefix:    "p",
				PositiveSuffix:    "s",
				MinFractionDigits: 1,
				MaxFractionDigits: 3,
			},
			expected: "p123:12s",
		},
		{
			val: -123.12,
			nf: schema.NumberFormat{
				Symbols: schema.Symbols{
					Zero:    '0',
					Decimal: ":",
				},
				NegativePrefix:    "np",
				NegativeSuffix:    "ns",
				MinFractionDigits: 3,
				MaxFractionDigits: 5,
			},
			expected: "np123:120ns",
		},
	}

	for _, test := range tests {
		var w writer
		test.val.format(&w, &test.nf, noCurrency)
		if s := w.String(); s != test.expected {
			t.Errorf("unexpected Float format for %q: %s", test.expected, s)
		}
	}
}
