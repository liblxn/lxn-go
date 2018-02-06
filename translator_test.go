package lxn

import (
	"testing"

	schema "github.com/liblxn/lxn/schema/golang"
)

func TestFormatMsg(t *testing.T) {
	tests := []struct {
		msg      schema.Message
		ctx      Context
		loc      schema.Locale
		expected string
	}{
		// no replacement
		{
			msg: schema.Message{
				Text: []string{"foobar"},
			},
			expected: "foobar",
		},
		{
			msg: schema.Message{
				Text: []string{"foo", "bar"},
			},
			expected: "foobar",
		},

		// string replacement
		{
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.StringReplacement,
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.NumberReplacement,
					},
				},
			},
			loc: schema.Locale{
				DecimalFormat: schema.NumberFormat{
					Symbols: schema.Symbols{Zero: '0'},
				},
			},
			ctx: Context{
				"replkey": Int(7),
			},
			expected: "foo 7 bar",
		},

		// percent replacement
		{
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.PercentReplacement,
					},
				},
			},
			loc: schema.Locale{
				PercentFormat: schema.NumberFormat{
					Symbols:        schema.Symbols{Zero: '0', Percent: "%"},
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.MoneyReplacement,
						Details: schema.ReplacementDetails{
							schema.MoneyDetails{Currency: "currkey"},
						},
					},
				},
			},
			loc: schema.Locale{
				MoneyFormat: schema.NumberFormat{
					Symbols:        schema.Symbols{Zero: '0'},
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.PluralReplacement,
						Details: schema.ReplacementDetails{
							schema.PluralDetails{
								Type: schema.Cardinal,
								Variants: map[schema.PluralTag]schema.Message{
									schema.Other: {
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.PluralReplacement,
						Details: schema.ReplacementDetails{
							schema.PluralDetails{
								Type: schema.Cardinal,
								Variants: map[schema.PluralTag]schema.Message{
									schema.Few: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			loc: schema.Locale{
				CardinalPlurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand: schema.AbsoluteValue,
								Ranges:  []schema.Range{{LowerBound: 7, UpperBound: 7}},
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.PluralReplacement,
						Details: schema.ReplacementDetails{
							schema.PluralDetails{
								Type: schema.Ordinal,
								Variants: map[schema.PluralTag]schema.Message{
									schema.Few: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			loc: schema.Locale{
				OrdinalPlurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand: schema.AbsoluteValue,
								Ranges:  []schema.Range{{LowerBound: 7, UpperBound: 7}},
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.PluralReplacement,
						Details: schema.ReplacementDetails{
							schema.PluralDetails{
								Type: schema.Ordinal,
								Variants: map[schema.PluralTag]schema.Message{
									schema.Other: {
										Text: []string{"plural"},
									},
								},
							},
						},
					},
				},
			},
			loc: schema.Locale{
				OrdinalPlurals: []schema.Plural{
					{
						Tag: schema.Few,
						Rules: []schema.PluralRule{
							{
								Operand: schema.AbsoluteValue,
								Ranges:  []schema.Range{{LowerBound: 7, UpperBound: 7}},
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.PluralReplacement,
						Details: schema.ReplacementDetails{
							schema.PluralDetails{
								Type: schema.Ordinal,
								Custom: map[int64]schema.Message{
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.SelectReplacement,
						Details: schema.ReplacementDetails{
							schema.SelectDetails{
								Cases: map[string]schema.Message{
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.SelectReplacement,
						Details: schema.ReplacementDetails{
							schema.SelectDetails{
								Cases: map[string]schema.Message{
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 0,
						Type:    schema.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("abc"),
			},
			expected: "abcfoo  bar",
		},
		{
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 2,
						Type:    schema.StringReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("abc"),
			},
			expected: "foo  barabc",
		},
		{
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey1",
						TextPos: 0,
						Type:    schema.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 0,
						Type:    schema.StringReplacement,
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey1",
						TextPos: 1,
						Type:    schema.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 1,
						Type:    schema.StringReplacement,
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey1",
						TextPos: 2,
						Type:    schema.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 2,
						Type:    schema.StringReplacement,
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey1",
						TextPos: 0,
						Type:    schema.StringReplacement,
					},
					{
						Key:     "replkey2",
						TextPos: 2,
						Type:    schema.StringReplacement,
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
		msg      schema.Message
		ctx      Context
		loc      schema.Locale
		expected string
	}{
		// missing variable
		{
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.StringReplacement,
					},
				},
			},
			expected: "foo %!(MISSING:replkey) bar",
		},

		// missing currency variable
		{
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.MoneyReplacement,
						Details: schema.ReplacementDetails{
							schema.MoneyDetails{Currency: "currkey"},
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.ReplacementType(99),
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.NumberReplacement,
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
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.MoneyReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("foo"),
			},
			expected: "foo %!(CORRUPTED:replkey) bar",
		},
		{
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.PluralReplacement,
					},
				},
			},
			ctx: Context{
				"replkey": String("foo"),
			},
			expected: "foo %!(CORRUPTED:replkey) bar",
		},
		{
			msg: schema.Message{
				Text: []string{"foo ", " bar"},
				Replacements: []schema.Replacement{
					{
						Key:     "replkey",
						TextPos: 1,
						Type:    schema.SelectReplacement,
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
