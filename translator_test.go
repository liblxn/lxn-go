package lxn

import (
	"testing"

	"github.com/liblxn/lxn-go/internal/lxn"
)

func TestFormatMsg(t *testing.T) {
	tests := []struct {
		msg      lxn.Message
		ctx      Context
		loc      lxn.Locale
		expected string
	}{
		// no replacement
		{
			msg: lxn.Message{
				Text: []string{"foobar"},
			},
			expected: "foobar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo", "bar"},
			},
			expected: "foobar",
		},

		// string replacement
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("replval"),
			},
			expected: "foo replval bar",
		},

		// number replacement
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.NumberReplacement,
					},
				},
			},
			loc: lxn.Locale{
				DecimalFormat: lxn.NumberFormat{
					Symbols: lxn.Symbols{Zero: '0'},
				},
			},
			ctx: Context{
				"replkey": Int(7),
			},
			expected: "foo 7 bar",
		},

		// percent replacement
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.PercentReplacement,
					},
				},
			},
			loc: lxn.Locale{
				PercentFormat: lxn.NumberFormat{
					Symbols:        lxn.Symbols{Zero: '0', Percent: "%"},
					PositiveSuffix: "%",
				},
			},
			ctx: Context{
				"replkey": Uint(7),
			},
			expected: "foo 7% bar",
		},

		// money replacement
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.MoneyReplacement,
						Details: lxn.ReplacementDetails{
							lxn.MoneyDetails{Currency: "currkey"},
						},
					},
				},
			},
			loc: lxn.Locale{
				MoneyFormat: lxn.NumberFormat{
					Symbols:        lxn.Symbols{Zero: '0'},
					PositiveSuffix: string(currencyPlaceholder),
				},
			},
			ctx: Context{
				"replkey": Uint(7),
				"currkey": String("EUR"),
			},
			expected: "foo 7EUR bar",
		},

		// plural replacement
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.PluralReplacement,
						Details: lxn.ReplacementDetails{
							Value: lxn.PluralDetails{
								Type: lxn.Cardinal,
								Variants: map[lxn.PluralCategory]lxn.Message{
									lxn.Other: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			ctx: Context{
				"replkey": String("nan"),
			},
			expected: "foo plural bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.PluralReplacement,
						Details: lxn.ReplacementDetails{
							Value: lxn.PluralDetails{
								Type: lxn.Cardinal,
								Variants: map[lxn.PluralCategory]lxn.Message{
									lxn.Few: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			loc: lxn.Locale{
				CardinalPlurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand: lxn.AbsoluteValue,
								Ranges:  []lxn.Range{{LowerBound: 7, UpperBound: 7}},
							},
						},
					},
				},
			},
			ctx: Context{
				"replkey": Int(7),
			},
			expected: "foo plural bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.PluralReplacement,
						Details: lxn.ReplacementDetails{
							Value: lxn.PluralDetails{
								Type: lxn.Ordinal,
								Variants: map[lxn.PluralCategory]lxn.Message{
									lxn.Few: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			loc: lxn.Locale{
				OrdinalPlurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand: lxn.AbsoluteValue,
								Ranges:  []lxn.Range{{LowerBound: 7, UpperBound: 7}},
							},
						},
					},
				},
			},
			ctx: Context{
				"replkey": Uint(7),
			},
			expected: "foo plural bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.PluralReplacement,
						Details: lxn.ReplacementDetails{
							Value: lxn.PluralDetails{
								Type: lxn.Ordinal,
								Variants: map[lxn.PluralCategory]lxn.Message{
									lxn.Other: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			loc: lxn.Locale{
				OrdinalPlurals: []lxn.Plural{
					{
						Category: lxn.Few,
						Rules: []lxn.PluralRule{
							{
								Operand: lxn.AbsoluteValue,
								Ranges:  []lxn.Range{{LowerBound: 7, UpperBound: 7}},
							},
						},
					},
				},
			},
			ctx: Context{
				"replkey": Int(7),
			},
			expected: "foo plural bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.PluralReplacement,
						Details: lxn.ReplacementDetails{
							Value: lxn.PluralDetails{
								Type: lxn.Ordinal,
								Custom: map[int64]lxn.Message{
									7: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			ctx: Context{
				"replkey": Int(7),
			},
			expected: "foo plural bar",
		},

		// select replacement
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.SelectReplacement,
						Details: lxn.ReplacementDetails{
							Value: lxn.SelectDetails{
								Cases: map[string]lxn.Message{
									"abc": {
										Text: []string{"select"},
									},
								},
							},
						},
					},
				},
			},
			ctx: Context{
				"replkey": String("abc"),
			},
			expected: "foo select bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.SelectReplacement,
						Details: lxn.ReplacementDetails{
							lxn.SelectDetails{
								Cases: map[string]lxn.Message{
									"abc": {
										Text: []string{"select"},
									},
								},
								Fallback: "abc",
							},
						},
					},
				},
			},
			ctx: Context{
				"replkey": String("def"),
			},
			expected: "foo select bar",
		},

		// replacement positioning
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 0,
						Type:    lxn.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("abc"),
			},
			expected: "abcfoo  bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 2,
						Type:    lxn.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("abc"),
			},
			expected: "foo  barabc",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey1",
						TextPos: 0,
						Type:    lxn.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 0,
						Type:    lxn.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey1": String("abc"),
				"replkey2": String("def"),
			},
			expected: "abcdeffoo  bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey1",
						TextPos: 1,
						Type:    lxn.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 1,
						Type:    lxn.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey1": String("abc"),
				"replkey2": String("def"),
			},
			expected: "foo abcdef bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey1",
						TextPos: 2,
						Type:    lxn.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 2,
						Type:    lxn.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey1": String("abc"),
				"replkey2": String("def"),
			},
			expected: "foo  barabcdef",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey1",
						TextPos: 0,
						Type:    lxn.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 2,
						Type:    lxn.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey1": String("abc"),
				"replkey2": String("def"),
			},
			expected: "abcfoo  bardef",
		},
	}

	for _, test := range tests {
		var w writer
		formatMsg(&w, &test.msg, test.ctx, &test.loc)
		if s := w.String(); s != test.expected {
			t.Errorf("unexpected message format for %q: %q", test.expected, s)
		}
	}
}

func TestFormatMsgWithIncompleteInput(t *testing.T) {
	tests := []struct {
		msg      lxn.Message
		ctx      Context
		loc      lxn.Locale
		expected string
	}{
		// missing variable
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.StringReplacement,
					},
				},
			},
			expected: "foo %!(MISSING:replkey) bar",
		},

		// missing currency variable
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.MoneyReplacement,
						Details: lxn.ReplacementDetails{
							lxn.MoneyDetails{Currency: "currkey"},
						},
					},
				},
			},
			ctx: Context{
				"replkey": Uint(7),
			},
			expected: "foo %!(MISSING:currkey) bar",
		},

		// unsupported replacement type
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.ReplacementType(99),
					},
				},
			},
			ctx: Context{
				"replkey": String("foo"),
			},
			expected: "foo %!(UNSUPPORTED:ReplType-99) bar",
		},

		// invalid number type
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.NumberReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("foo"),
			},
			expected: "foo %!(INVALID:replkey) bar",
		},

		// corrupted details
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.MoneyReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("foo"),
			},
			expected: "foo %!(CORRUPTED:replkey) bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.PluralReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("foo"),
			},
			expected: "foo %!(CORRUPTED:replkey) bar",
		},
		{
			msg: lxn.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []lxn.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    lxn.SelectReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("foo"),
			},
			expected: "foo %!(CORRUPTED:replkey) bar",
		},
	}

	for _, test := range tests {
		var w writer
		formatMsg(&w, &test.msg, test.ctx, &test.loc)
		if s := w.String(); s != test.expected {
			t.Errorf("unexpected message format for %q: %q", test.expected, s)
		}
	}
}
