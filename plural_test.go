package lxn

import (
	"testing"

	schema "github.com/liblxn/lxn/schema/golang"
)

func TestPluralTag(t *testing.T) {
	tests := map[string][]struct {
		num      number
		nf       schema.NumberFormat
		plurals  []schema.Plural
		expected schema.PluralTag
	}{
		"absolute values": {
			{
				num: Uint(7),
				nf: schema.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.AbsoluteValue,
								Ranges:     []schema.Range{{LowerBound: 7, UpperBound: 7}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Few,
			},
			{
				num: Int(7),
				nf: schema.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.AbsoluteValue,
								Modulo:     5,
								Ranges:     []schema.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Few,
			},
			{
				num: Float(7.5),
				nf: schema.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.AbsoluteValue,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 6, UpperBound: 8}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Few,
			},
		},

		"integer digits": {
			{
				num: Float(7.5),
				nf: schema.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Many,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Ranges:     []schema.Range{{LowerBound: 7, UpperBound: 7}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Many,
			},
			{
				num: Float(0.5),
				nf: schema.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Many,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Modulo:     10,
								Ranges:     []schema.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Many,
			},
			{
				num: Float(75.25),
				nf: schema.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Many,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Modulo:     10,
								Ranges:     []schema.Range{{LowerBound: 5, UpperBound: 5}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Many,
			},
		},

		"number of fraction digits with trailing zeros": {
			{
				num: Uint(7),
				nf: schema.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Two,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.NumFracDigits,
								Ranges:     []schema.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Two,
			},
			{
				num: Int(7),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 4,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Two,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.NumFracDigits,
								Ranges:     []schema.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Two,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 6,
					MaxFractionDigits: 6,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Two,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.NumFracDigits,
								Modulo:     4,
								Ranges:     []schema.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Two,
			},
		},

		"number of fraction digits without trailing zeros": {
			{
				num: Int(7),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.One,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.NumFracDigitsNoZeros,
								Ranges:     []schema.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.One,
			},
			{
				num: Float(7.5),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.One,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.NumFracDigitsNoZeros,
								Ranges:     []schema.Range{{LowerBound: 1, UpperBound: 1}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.One,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 6,
					MaxFractionDigits: 6,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.One,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.NumFracDigitsNoZeros,
								Modulo:     4,
								Ranges:     []schema.Range{{LowerBound: 2, UpperBound: 2}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.One,
			},
		},

		"fraction digits value with trailing zeros": {
			{
				num: Uint(7),
				nf: schema.NumberFormat{
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Zero,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.FracDigits,
								Ranges:     []schema.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Zero,
			},
			{
				num: Float(7.5),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Zero,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.FracDigits,
								Ranges:     []schema.Range{{LowerBound: 50, UpperBound: 50}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Zero,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 6,
					MaxFractionDigits: 6,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Zero,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.FracDigits,
								Modulo:     100,
								Ranges:     []schema.Range{{LowerBound: 92, UpperBound: 92}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Zero,
			},
		},

		"fraction digits value without trailing zeros": {
			{
				num: Int(7),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Zero,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.FracDigitsNoZeros,
								Ranges:     []schema.Range{{LowerBound: 0, UpperBound: 0}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Zero,
			},
			{
				num: Float(7.5),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Zero,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.FracDigitsNoZeros,
								Ranges:     []schema.Range{{LowerBound: 5, UpperBound: 5}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Zero,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 8,
					MaxFractionDigits: 8,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Zero,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.FracDigitsNoZeros,
								Modulo:     100,
								Ranges:     []schema.Range{{LowerBound: 92, UpperBound: 92}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Zero,
			},
		},

		"conjunction": {
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Few,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Other,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Other,
			},
		},

		"disjunction": {
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Disjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Few,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Disjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Few,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Disjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Other,
			},
		},

		"conjunction and disjunction": {
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Many,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.Disjunction,
							},
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Many,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Many,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.Disjunction,
							},
							{
								Operand:    schema.IntegerDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Many,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Many,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.Disjunction,
							},
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Many,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Many,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.Disjunction,
							},
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Many,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Many,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.Disjunction,
							},
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Many,
			},
			{
				num: Float(3.141592),
				nf: schema.NumberFormat{
					MinFractionDigits: 2,
					MaxFractionDigits: 2,
				},
				plurals: []schema.Plural{
					{
						Tag: schema.Many,
						Rules: []schema.PluralRule{
							{
								Operand:    schema.IntegerDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.Disjunction,
							},
							{
								Operand:    schema.IntegerDigits,
								Negate:     true,
								Ranges:     []schema.Range{{LowerBound: 3, UpperBound: 3}},
								Connective: schema.Conjunction,
							},
							{
								Operand:    schema.FracDigits,
								Negate:     false,
								Ranges:     []schema.Range{{LowerBound: 14, UpperBound: 14}},
								Connective: schema.None,
							},
						},
					},
				},
				expected: schema.Other,
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
