package lxn

import (
	"testing"

	"github.com/liblxn/lxn-go/internal/lxn"
)

func TestPluralTag(t *testing.T) {
	tests := map[string][]struct {
		num      number
		nf       lxn.NumberFormat
		plurals  []lxn.Plural
		expected lxn.PluralCategory
	}{
		"absolute values": {
			{
				num: Uint(7),
				nf: lxn.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.AbsoluteValue,
								Ranges:     []lxn.Range{{LowerBound: 7, UpperBound: 7}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Few,
			},
			{
				num: Int(7),
				nf: lxn.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.AbsoluteValue,
								Modulo:     5,
								Ranges:     []lxn.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Few,
			},
			{
				num: Float(7.5),
				nf: lxn.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.AbsoluteValue,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 6, UpperBound: 8}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Few,
			},
		},

		"integer digits": {
			{
				num: Float(7.5),
				nf: lxn.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Many,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Ranges:     []lxn.Range{{LowerBound: 7, UpperBound: 7}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Many,
			},
			{
				num: Float(0.5),
				nf: lxn.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Many,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Modulo:     10,
								Ranges:     []lxn.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Many,
			},
			{
				num: Float(75.25),
				nf: lxn.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Many,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Modulo:     10,
								Ranges:     []lxn.Range{{LowerBound: 5, UpperBound: 5}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Many,
			},
		},

		"number of fraction digits with trailing zeros": {
			{
				num: Uint(7),
				nf: lxn.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Two,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.NumFracDigits,
								Ranges:     []lxn.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Two,
			},
			{
				num: Int(7),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 4,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Two,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.NumFracDigits,
								Ranges:     []lxn.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Two,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 6,
					MaxFractionDigits: 6,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Two,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.NumFracDigits,
								Modulo:     4,
								Ranges:     []lxn.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Two,
			},
		},

		"number of fraction digits without trailing zeros": {
			{
				num: Int(7),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.One,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.NumFracDigitsNoZeros,
								Ranges:     []lxn.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.One,
			},
			{
				num: Float(7.5),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.One,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.NumFracDigitsNoZeros,
								Ranges:     []lxn.Range{{LowerBound: 1, UpperBound: 1}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.One,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 6,
					MaxFractionDigits: 6,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.One,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.NumFracDigitsNoZeros,
								Modulo:     4,
								Ranges:     []lxn.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.One,
			},
		},

		"fraction digits value with trailing zeros": {
			{
				num: Uint(7),
				nf: lxn.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Zero,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.FracDigits,
								Ranges:     []lxn.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Zero,
			},
			{
				num: Float(7.5),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Zero,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.FracDigits,
								Ranges:     []lxn.Range{{LowerBound: 50, UpperBound: 50}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Zero,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 6,
					MaxFractionDigits: 6,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Zero,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.FracDigits,
								Modulo:     100,
								Ranges:     []lxn.Range{{LowerBound: 92, UpperBound: 92}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Zero,
			},
		},

		"fraction digits value without trailing zeros": {
			{
				num: Int(7),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Zero,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.FracDigitsNoZeros,
								Ranges:     []lxn.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Zero,
			},
			{
				num: Float(7.5),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Zero,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.FracDigitsNoZeros,
								Ranges:     []lxn.Range{{LowerBound: 5, UpperBound: 5}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Zero,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 8,
					MaxFractionDigits: 8,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Zero,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.FracDigitsNoZeros,
								Modulo:     100,
								Ranges:     []lxn.Range{{LowerBound: 92, UpperBound: 92}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Zero,
			},
		},

		"conjunction": {
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Few,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Other,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Other,
			},
		},

		"disjunction": {
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Disjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Few,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Disjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Few,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Disjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Other,
			},
		},

		"conjunction and disjunction": {
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Many,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.Disjunction,
							},
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Many,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Many,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.Disjunction,
							},
							{
								Operand:    lxn.IntegerDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Many,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Many,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.Disjunction,
							},
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Many,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Many,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.Disjunction,
							},
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Many,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Many,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.Disjunction,
							},
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Many,
			},
			{
				num: Float(3.141592),
				nf: lxn.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []lxn.Plural{
					{
						Category: lxn.Many,
						Rules: []lxn.PluralRule{
							{
								Operand:    lxn.IntegerDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.Disjunction,
							},
							{
								Operand:    lxn.IntegerDigits,
								Negate:     true,
								Ranges:     []lxn.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: lxn.Conjunction,
							},
							{
								Operand:    lxn.FracDigits,
								Negate:     false,
								Ranges:     []lxn.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: lxn.None,
							},
						},
					},
				},
				expected: lxn.Other,
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
