package lxn

import (
	"github.com/liblxn/lxn-go/internal/lxn"
)

func pluralTag(num number, nf *lxn.NumberFormat, plurals []lxn.Plural) lxn.PluralCategory {
	var buf [maxFloatDigits]rune

	intDigits, fracDigits := num.digits(buf[:], nf, 0)
	op := newOperands(intDigits, fracDigits)

	for _, p := range plurals {
		match := false
		for i := 0; i < len(p.Rules); i++ {
			match = op.matchRule(&p.Rules[i], fracDigits)
			if !match {
				// We found an operand in the conjunction that is false, i.e. our
				// whole condition evaluates to false. So we skip the conjunction.
				for i < len(p.Rules) && p.Rules[i].Connective == lxn.Conjunction {
					i++
				}
			} else if p.Rules[i].Connective == lxn.Disjunction {
				// We found an operand in the disjunction that is true, i.e. our
				// whole condition evaluates to true.
				break
			}
		}
		if match {
			return p.Category
		}
	}
	return lxn.Other
}

type operands struct {
	i int64
	v int64
	f int64
	w int64
	t int64
}

func newOperands(intDigits []rune, fracDigits []rune) operands {
	digitsValue := func(digits []rune) (val int64) {
		for _, d := range digits {
			val *= 10
			val += int64(d)
		}
		return val
	}

	return operands{
		i: digitsValue(intDigits),
		f: digitsValue(fracDigits),
		v: int64(len(fracDigits)),
		w: -1, // lazily calculated
		t: -1, // lazily calculated
	}
}

func (op *operands) matchRule(r *lxn.PluralRule, fracDigits []rune) bool {
	var x int64
	switch r.Operand {
	case lxn.AbsoluteValue: // n
		// Since the ranges contain integer values only, we do not match
		// if n has non-zero fractional digits.
		if op.f != 0 {
			return r.Negate
		}
		fallthrough

	case lxn.IntegerDigits: // i
		x = op.i

	case lxn.NumFracDigits: // v
		x = op.v

	case lxn.NumFracDigitsNoZeros: // w
		if op.w < 0 {
			op.w = op.v
			for op.w > 0 && fracDigits[op.w-1] == 0 {
				op.w--
			}
		}
		x = op.w

	case lxn.FracDigits: // f
		x = op.f

	case lxn.FracDigitsNoZeros: // t
		if op.t < 0 {
			op.t = op.f
			for op.t > 0 && op.t%10 == 0 {
				op.t /= 10
			}
		}
		x = op.t

	default:
		return r.Negate // ignore unknown operands
	}

	if r.Modulo > 0 {
		x %= int64(r.Modulo)
	}

	for _, rng := range r.Ranges {
		if int64(rng.LowerBound) <= x && x <= int64(rng.UpperBound) {
			return !r.Negate // match!
		}
	}
	return r.Negate
}
