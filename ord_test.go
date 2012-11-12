// Copyright 2012 Glenn Brown
// See LICENSE.

package ord

import (
	"fmt"
	"testing"
)

// Type slow implements the Slow interface.
type slow int
func (s slow) Less(i interface{}) bool { return s < i.(slow) }

// Type fast implements the Fast interface.
type fast int
func (f fast) Less(i interface{}) bool { return f < i.(fast) }
func (f fast) Score() float64 { return float64(f) }

// The pairs table holds test pairs {a,b} where a<b.
var pairs = [][]interface{} {
	// The Slow case is first, and is the only one with equal scores
	{ slow(0)		, slow(1)		},
	// The rest have distinquishing scores.
	{ fast(0)		, fast(1)		},
	{ ([]byte)("abcdef")	, ([]byte)("abcdeg")	},
	{ int(0)		, int(1)		},
	{ int8(0)		, int8(1)		},
	{ int16(0)		, int16(1)		},	    
	{ int32(0)		, int32(1)		},	    
	{ int64(0)		, int64(1)		},	    
	{ float32(0)		, float32(1)		},
	{ float64(0)		, float64(1)		},
	{ "abcdef"		, "abcdei"		},
	{ uint(0)		, uint(1)		},
	{  uint8(0)		, uint8(1)		},	    
	{ uint16(0)		, uint16(1)		},
	{ uint32(0)		, uint32(1)		},
	{ uint64(0)		, uint64(1)		},
	{ uintptr(0)		, uintptr(1)		},
}

func TestFns(t *testing.T) {
	for i := range pairs {
		a, b := pairs[i][0], pairs[i][1]
		less, score := Fns(a)
		if less(b, a) { t.Error(fmt.Sprintf("1:%T", a)) }
		if less(a, a) { t.Error(fmt.Sprintf("2:%T", a)) }
		if !less(a, b) { t.Error(fmt.Sprintf("3:%T", a)) }
		if 0 == i {
			if score(b) != score(a) { t.Error(fmt.Sprintf("4:%T", a)) }
		} else {
			if score(b) < score(a) { t.Error(fmt.Sprintf("5:%T", a)) }
			if score(a) < score(a) { t.Error(fmt.Sprintf("6:%T", a)) }
			if !(score(a) < score(b)) { t.Error(fmt.Sprintf("7:%T", a)) }
		}
	}
}

func TestFnsReversed(t *testing.T) {
	for i := range pairs {
		b, a := pairs[i][0], pairs[i][1]
		less, score := FnsReversed(a)
		if less(b, a) { t.Error(fmt.Sprintf("1:%T", a)) }
		if less(a, a) { t.Error(fmt.Sprintf("2:%T", a)) }
		if !less(a, b) { t.Error(fmt.Sprintf("3:%T", a)) }
		if 0 == i {
			if score(b) != score(a) { t.Error(fmt.Sprintf("4:%T", a)) }
		} else {
			if score(b) < score(a) { t.Error(fmt.Sprintf("5:%T", a)) }
			if score(a) < score(a) { t.Error(fmt.Sprintf("6:%T", a)) }
			if !(score(a) < score(b)) { t.Error(fmt.Sprintf("7:%T", a)) }
		}
	}
}

func TestFnScore(t *testing.T) {
	for i := range pairs {
		a, b := pairs[i][0], pairs[i][1]
		less, aScore := FnScore(a)
		_, bScore := FnScore(b)
		if less(b, a) { t.Error(fmt.Sprintf("1:%T", a)) }
		if less(a, a) { t.Error(fmt.Sprintf("2:%T", a)) }
		if !less(a, b) { t.Error(fmt.Sprintf("3:%T", a)) }
		if 0 == i {
			if bScore != aScore { t.Error(fmt.Sprintf("4:%T", a)) }
		} else {
			if bScore < aScore { t.Error(fmt.Sprintf("5:%T", a)) }
			if aScore < aScore { t.Error(fmt.Sprintf("6:%T", a)) }
			if !(aScore < bScore) { t.Error(fmt.Sprintf("7:%T", a)) }
		}
	}
}

func TestFnScoreReversed(t *testing.T) {
	for i := range pairs {
		b, a := pairs[i][0], pairs[i][1]
		less, aScore := FnScoreReversed(a)
		_, bScore := FnScoreReversed(b)
		if less(b, a) { t.Error(fmt.Sprintf("1:%T", a)) }
		if less(a, a) { t.Error(fmt.Sprintf("2:%T", a)) }
		if !less(a, b) { t.Error(fmt.Sprintf("3:%T", a)) }
		if 0 == i {
			if bScore != aScore { t.Error(fmt.Sprintf("4:%T", a)) }
		} else {
			if bScore < aScore { t.Error(fmt.Sprintf("5:%T", a)) }
			if aScore < aScore { t.Error(fmt.Sprintf("6:%T", a)) }
			if !(aScore < bScore) { t.Error(fmt.Sprintf("7:%T", a)) }
		}
	}
}

