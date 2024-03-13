package lxn

import (
	"math"
	"strconv"

	"github.com/liblxn/lxn-go/internal/lxn"
)

type number interface {
	Variable
	digits(buf []rune, nf *lxn.NumberFormat, zero rune) ([]rune, []rune) // (integer, fraction)
	format(w *writer, nf *lxn.NumberFormat, currency string)
}

var (
	_ number = Int(0)
	_ number = Uint(0)
	_ number = Float(0)
)

const (
	maxIntDigits   = 32 + 32 // integer + fraction digits
	maxFloatDigits = 256
)

// Int is a signed integer variable which can be passed to message replacements.
type Int int64

// String implements the Variable interface.
func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// returns (integer digits, fraction digits)
func (i Int) digits(buf []rune, nf *lxn.NumberFormat, zero rune) ([]rune, []rune) {
	if i < 0 {
		i = -i
	}
	return Uint(i).digits(buf, nf, zero)
}

func (i Int) format(w *writer, nf *lxn.NumberFormat, currency string) {
	negative := i < 0
	if negative {
		i = -i
	}
	Uint(i).fmt(w, nf, currency, negative)
}

// Uint is an unsigned integer variable which can be passed to message replacements.
type Uint uint64

// String implements the Variable interface.
func (ui Uint) String() string {
	return strconv.FormatUint(uint64(ui), 10)
}

// returns (integer digits, fraction digits)
func (ui Uint) digits(buf []rune, nf *lxn.NumberFormat, zero rune) ([]rune, []rune) {
	// fractional digits
	fracidx := len(buf)
	for i := 0; i < nf.MinFractionDigits; i++ {
		fracidx--
		buf[fracidx] = zero
	}

	// integer digits
	intidx := fracidx
	for ui >= 10 {
		intidx--
		rest := ui / 10
		buf[intidx] = zero + rune(ui-rest*10)
		ui = rest
	}
	intidx--
	buf[intidx] = zero + rune(ui)

	for i := fracidx - intidx; i < nf.MinIntegerDigits; i++ {
		intidx--
		buf[intidx] = zero
	}

	return buf[intidx:fracidx], buf[fracidx:]
}

func (ui Uint) format(w *writer, nf *lxn.NumberFormat, currency string) {
	ui.fmt(w, nf, currency, false)
}

func (ui Uint) fmt(w *writer, nf *lxn.NumberFormat, currency string, negative bool) {
	prefix, suffix := nf.PositivePrefix, nf.PositiveSuffix
	if negative {
		prefix, suffix = nf.NegativePrefix, nf.NegativeSuffix
	}

	var buf [maxIntDigits]rune
	intDigits, fracDigits := ui.digits(buf[:], nf, rune(nf.Symbols.Zero))

	w.WriteAffix(prefix, &nf.Symbols, currency)
	w.WriteInt(intDigits, nf)
	w.WriteFrac(fracDigits, nf)
	w.WriteAffix(suffix, &nf.Symbols, currency)
}

// Float is a floating-point variable which can be passed to message replacements.
type Float float64

// String implements the Variable interface.
func (f Float) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

// returns (integer digits, fraction digits)
func (f Float) digits(buf []rune, nf *lxn.NumberFormat, zero rune) ([]rune, []rune) {
	if f < 0 {
		f = -f
	}

	var fmtbuf [maxFloatDigits]byte
	fmt := strconv.AppendFloat(fmtbuf[:0], float64(f), 'f', nf.MaxFractionDigits, 64)

	// trim trailing zeros
	fmtidx := len(fmt)
	for diff := nf.MaxFractionDigits - nf.MinFractionDigits; diff > 0 && fmt[fmtidx-1] == '0'; diff-- {
		fmtidx--
	}

	// fractional digits
	fracidx := len(buf)
	if nf.MaxFractionDigits > 0 {
		for fmtidx > 0 {
			fmtidx--
			d := fmt[fmtidx]
			if d == '.' {
				break
			}
			fracidx--
			buf[fracidx] = zero + rune(d-'0')
		}
	}

	// integer digits
	intidx := fracidx
	for fmtidx > 0 {
		fmtidx--
		intidx--
		buf[intidx] = zero + rune(fmt[fmtidx]-'0')
	}
	for i := fracidx - intidx; i < nf.MinIntegerDigits; i++ {
		intidx--
		buf[intidx] = zero
	}

	return buf[intidx:fracidx], buf[fracidx:]
}

func (f Float) format(w *writer, nf *lxn.NumberFormat, currency string) {
	switch {
	case math.IsNaN(float64(f)):
		w.WriteString(nf.Symbols.Nan)
		return

	case math.IsInf(float64(f), 0):
		prefix, suffix := nf.PositivePrefix, nf.PositiveSuffix
		if math.IsInf(float64(f), -1) {
			prefix, suffix = nf.NegativePrefix, nf.NegativeSuffix
		}

		w.WriteAffix(prefix, &nf.Symbols, currency)
		w.WriteString(nf.Symbols.Inf)
		w.WriteAffix(suffix, &nf.Symbols, currency)
		return
	}

	var buf [maxFloatDigits]rune
	intDigits, fracDigits := f.digits(buf[:], nf, rune(nf.Symbols.Zero))

	prefix, suffix := nf.PositivePrefix, nf.PositiveSuffix
	if f < 0 {
		prefix, suffix = nf.NegativePrefix, nf.NegativeSuffix
	}

	w.WriteAffix(prefix, &nf.Symbols, currency)
	w.WriteInt(intDigits, nf)
	w.WriteFrac(fracDigits, nf)
	w.WriteAffix(suffix, &nf.Symbols, currency)
}
