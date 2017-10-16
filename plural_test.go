package lxn

import (
	"testing"

	"github.com/liblxn/lxn-go/internal"
)

func TestPluralTag(t *testing.T) {
	tests := map[string][]struct {
		num      number
		nf       internal.NumberFormat
		plurals  []internal.Plural
		expected internal.PluralTag
	}{
		"absolute values": {
			{
				num: Uint(7),
				nf: internal.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.AbsoluteValue,
								Ranges:     []internal.Range{{LowerBound: 7, UpperBound: 7}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Few,
			},
			{
				num: Int(7),
				nf: internal.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.AbsoluteValue,
								Modulo:     5,
								Ranges:     []internal.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Few,
			},
			{
				num: Float(7.5),
				nf: internal.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.AbsoluteValue,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 6, UpperBound: 8}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Few,
			},
		},

		"integer digits": {
			{
				num: Float(7.5),
				nf: internal.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Many,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Ranges:     []internal.Range{{LowerBound: 7, UpperBound: 7}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Many,
			},
			{
				num: Float(0.5),
				nf: internal.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Many,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Modulo:     10,
								Ranges:     []internal.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Many,
			},
			{
				num: Float(75.25),
				nf: internal.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Many,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Modulo:     10,
								Ranges:     []internal.Range{{LowerBound: 5, UpperBound: 5}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Many,
			},
		},

		"number of fraction digits with trailing zeros": {
			{
				num: Uint(7),
				nf: internal.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Two,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.NumFracDigits,
								Ranges:     []internal.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Two,
			},
			{
				num: Int(7),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 4,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Two,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.NumFracDigits,
								Ranges:     []internal.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Two,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 6,
					MaxFractionDigits: 6,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Two,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.NumFracDigits,
								Modulo:     4,
								Ranges:     []internal.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Two,
			},
		},

		"number of fraction digits without trailing zeros": {
			{
				num: Int(7),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.One,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.NumFracDigitsNoZeros,
								Ranges:     []internal.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.One,
			},
			{
				num: Float(7.5),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.One,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.NumFracDigitsNoZeros,
								Ranges:     []internal.Range{{LowerBound: 1, UpperBound: 1}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.One,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 6,
					MaxFractionDigits: 6,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.One,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.NumFracDigitsNoZeros,
								Modulo:     4,
								Ranges:     []internal.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.One,
			},
		},

		"fraction digits value with trailing zeros": {
			{
				num: Uint(7),
				nf: internal.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Zero,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.FracDigits,
								Ranges:     []internal.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Zero,
			},
			{
				num: Float(7.5),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Zero,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.FracDigits,
								Ranges:     []internal.Range{{LowerBound: 50, UpperBound: 50}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Zero,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 6,
					MaxFractionDigits: 6,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Zero,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.FracDigits,
								Modulo:     100,
								Ranges:     []internal.Range{{LowerBound: 92, UpperBound: 92}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Zero,
			},
		},

		"fraction digits value without trailing zeros": {
			{
				num: Int(7),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Zero,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.FracDigitsNoZeros,
								Ranges:     []internal.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Zero,
			},
			{
				num: Float(7.5),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Zero,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.FracDigitsNoZeros,
								Ranges:     []internal.Range{{LowerBound: 5, UpperBound: 5}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Zero,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 8,
					MaxFractionDigits: 8,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Zero,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.FracDigitsNoZeros,
								Modulo:     100,
								Ranges:     []internal.Range{{LowerBound: 92, UpperBound: 92}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Zero,
			},
		},

		"conjunction": {
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Few,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Other,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Other,
			},
		},

		"disjunction": {
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Disjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Few,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Disjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Few,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Disjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Other,
			},
		},

		"conjunction and disjunction": {
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Many,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.Disjunction,
							},
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Many,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Many,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.Disjunction,
							},
							{
								Operand:    internal.IntegerDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Many,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Many,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.Disjunction,
							},
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Many,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Many,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.Disjunction,
							},
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Many,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Many,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.Disjunction,
							},
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Many,
			},
			{
				num: Float(3.141592),
				nf: internal.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []internal.Plural{
					{
						Tag: internal.Many,
						Rules: []internal.PluralRule{
							{
								Operand:    internal.IntegerDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.Disjunction,
							},
							{
								Operand:    internal.IntegerDigits,
								Negate:     true,
								Ranges:     []internal.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: internal.Conjunction,
							},
							{
								Operand:    internal.FracDigits,
								Negate:     false,
								Ranges:     []internal.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: internal.None,
							},
						},
					},
				},
				expected: internal.Other,
			},
		},
	}

	for name := range tests {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests[name] {
				tag := pluralTag(test.num, &test.nf, test.plurals)
				if tag != test.expected {
					t.Errorf("unexpected plural tag: %v", tag)
				}
			}
		})
	}
}
