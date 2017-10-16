package lxn

import (
	"testing"

	"github.com/liblxn/lxn-go/internal"
)

func TestFormatMsg(t *testing.T) {
	tests := []struct {
		msg      internal.Message
		ctx      Context
		loc      internal.Locale
		expected string
	}{
		// no replacement
		{
			msg: internal.Message{
				Text: []string{"foobar"},
			},
			expected: "foobar",
		},
		{
			msg: internal.Message{
				Text: []string{"foo", "bar"},
			},
			expected: "foobar",
		},

		// string replacement
		{
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.StringReplacement,
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.NumberReplacement,
					},
				},
			},
			loc: internal.Locale{
				DecimalFormat: internal.NumberFormat{
					Symbols: internal.Symbols{Zero: '0'},
				},
			},
			ctx: Context{
				"replkey": Int(7),
			},
			expected: "foo 7 bar",
		},

		// percent replacement
		{
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.PercentReplacement,
					},
				},
			},
			loc: internal.Locale{
				PercentFormat: internal.NumberFormat{
					Symbols:        internal.Symbols{Zero: '0', Percent: "%"},
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.MoneyReplacement,
						Details: internal.ReplacementDetails{
							internal.MoneyDetails{Currency: "currkey"},
						},
					},
				},
			},
			loc: internal.Locale{
				MoneyFormat: internal.NumberFormat{
					Symbols:        internal.Symbols{Zero: '0'},
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.PluralReplacement,
						Details: internal.ReplacementDetails{
							internal.PluralDetails{
								Type: internal.Cardinal,
								Variants: map[internal.PluralTag]internal.Message{
									internal.Other: {
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.PluralReplacement,
						Details: internal.ReplacementDetails{
							internal.PluralDetails{
								Type: internal.Cardinal,
								Variants: map[internal.PluralTag]internal.Message{
									internal.Few: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			loc: internal.Locale{
				CardinalPlurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand: internal.AbsoluteValue,
								Ranges:  []internal.Range{{LowerBound: 7, UpperBound: 7}},
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.PluralReplacement,
						Details: internal.ReplacementDetails{
							internal.PluralDetails{
								Type: internal.Ordinal,
								Variants: map[internal.PluralTag]internal.Message{
									internal.Few: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			loc: internal.Locale{
				OrdinalPlurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand: internal.AbsoluteValue,
								Ranges:  []internal.Range{{LowerBound: 7, UpperBound: 7}},
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.PluralReplacement,
						Details: internal.ReplacementDetails{
							internal.PluralDetails{
								Type: internal.Ordinal,
								Variants: map[internal.PluralTag]internal.Message{
									internal.Other: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			loc: internal.Locale{
				OrdinalPlurals: []internal.Plural{
					{
						Tag: internal.Few,
						Rules: []internal.PluralRule{
							{
								Operand: internal.AbsoluteValue,
								Ranges:  []internal.Range{{LowerBound: 7, UpperBound: 7}},
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.PluralReplacement,
						Details: internal.ReplacementDetails{
							internal.PluralDetails{
								Type: internal.Ordinal,
								Custom: map[int64]internal.Message{
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.SelectReplacement,
						Details: internal.ReplacementDetails{
							internal.SelectDetails{
								Cases: map[string]internal.Message{
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.SelectReplacement,
						Details: internal.ReplacementDetails{
							internal.SelectDetails{
								Cases: map[string]internal.Message{
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 0,
						Type:    internal.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("abc"),
			},
			expected: "abcfoo  bar",
		},
		{
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 2,
						Type:    internal.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("abc"),
			},
			expected: "foo  barabc",
		},
		{
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey1",
						TextPos: 0,
						Type:    internal.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 0,
						Type:    internal.StringReplacement,
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey1",
						TextPos: 1,
						Type:    internal.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 1,
						Type:    internal.StringReplacement,
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey1",
						TextPos: 2,
						Type:    internal.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 2,
						Type:    internal.StringReplacement,
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey1",
						TextPos: 0,
						Type:    internal.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 2,
						Type:    internal.StringReplacement,
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
		msg      internal.Message
		ctx      Context
		loc      internal.Locale
		expected string
	}{
		// missing variable
		{
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.StringReplacement,
					},
				},
			},
			expected: "foo %!(MISSING:replkey) bar",
		},

		// missing currency variable
		{
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.MoneyReplacement,
						Details: internal.ReplacementDetails{
							internal.MoneyDetails{Currency: "currkey"},
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.ReplacementType(99),
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.NumberReplacement,
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
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.MoneyReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("foo"),
			},
			expected: "foo %!(CORRUPTED:replkey) bar",
		},
		{
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.PluralReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("foo"),
			},
			expected: "foo %!(CORRUPTED:replkey) bar",
		},
		{
			msg: internal.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []internal.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    internal.SelectReplacement,
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
