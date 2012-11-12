// Copyright 2012 Glenn Brown
// See LICENSE.
//
// Portions of this file are licensed as follows:
//
// > Copyright (c) 2011 Huan Du
// > 
// > Permission is hereby granted, free of charge, to any person obtaining a copy
// > of this software and associated documentation files (the "Software"), to deal
// > in the Software without restriction, including without limitation the rights
// > to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// > copies of the Software, and to permit persons to whom the Software is
// > furnished to do so, subject to the following conditions:
// > 
// > The above copyright notice and this permission notice shall be included in
// > all copies or substantial portions of the Software.
// > 
// > THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// > IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// > FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// > AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// > LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// > OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// > THE SOFTWARE.
	
// Package ord provides utility functions for fast ordered containers.
//
package ord

import (
	"bytes"
	"fmt"
)

// Any type implementing the Slow interface is supported by this package, but the Fast interface
// is faster.
type Slow interface { Less(interface{}) bool }

// Any type impementing the Fast interface is well-supported by this package.
//
type Fast interface {
	Less(interface{}) bool
	Score() float64
}

// IncreasingFns returns functions ordered containers can use to sort values in increasing order.
// If the container cannot conveniently cache the results, consider FnScore instead.
//
func Fns(a interface{}) (less func(a, b interface{}) bool, score func(a interface{}) float64) {
	switch a.(type) {
	case Fast:
		return func(a, b interface{}) bool { return a.(Fast).Less(b) },
		func(t interface{}) float64 { return t.(Fast).Score() }
	case Slow:
		return func(a, b interface{}) bool { return a.(Slow).Less(b) },
		func(t interface{}) float64 { return float64(0.0) }
	case []byte:
		return func(a, b interface{}) bool { return bytes.Compare (a.([]byte), b.([]byte)) < 0 },
		func(t interface{}) float64 { return scoreBytes(t.([]byte)) }
	case float32:
		return func(a, b interface{}) bool { return a.(float32) < b.(float32) },
		func(t interface{}) float64 { return float64(t.(float32)) }
	case float64:
		return func(a, b interface{}) bool { return a.(float64) < b.(float64) },
		func(t interface{}) float64 { return float64(t.(float64)) }
	case int:
		return func(a, b interface{}) bool { return a.(int) < b.(int) },
		func(t interface{}) float64 { return float64(t.(int)) }
	case int16:
		return func(a, b interface{}) bool { return a.(int16) < b.(int16) },
		func(t interface{}) float64 { return float64(t.(int16)) }
	case int32:
		return func(a, b interface{}) bool { return a.(int32) < b.(int32) },
		func(t interface{}) float64 { return float64(t.(int32)) }
	case int64:
		return func(a, b interface{}) bool { return a.(int64) < b.(int64) },
		func(t interface{}) float64 { return float64(t.(int64)) }
	case int8:
		return func(a, b interface{}) bool { return a.(int8) < b.(int8) },
		func(t interface{}) float64 { return float64(t.(int8)) }
	case string:
		return func(a, b interface{}) bool { return a.(string) < b.(string) },
		func(t interface{}) float64 { return scoreBytes([]byte(t.(string))) }
	case uint:
		return func(a, b interface{}) bool { return a.(uint) < b.(uint) },
		func(t interface{}) float64 { return float64(t.(uint)) }
	case uint16:
		return func(a, b interface{}) bool { return a.(uint16) < b.(uint16) },
		func(t interface{}) float64 { return float64(t.(uint16)) }
	case uint32:
		return func(a, b interface{}) bool { return a.(uint32) < b.(uint32) },
		func(t interface{}) float64 { return float64(t.(uint32)) }
	case uint64:
		return func(a, b interface{}) bool { return a.(uint64) < b.(uint64) },
		func(t interface{}) float64 { return float64(t.(uint64)) }
	case uint8:
		return func(a, b interface{}) bool { return a.(uint8) < b.(uint8) },
		func(t interface{}) float64 { return float64(t.(uint8)) }
	case uintptr:
		return func(a, b interface{}) bool { return a.(uintptr) < b.(uintptr) },
		func(t interface{}) float64 { return float64(t.(uintptr)) }
	}
	panic(fmt.Sprintf("cannot order type %T: no method Less(a, b interface{}) bool", a))
}

// FnsReversed is like Fns, but sorts in reverse order.
//
func FnsReversed(a interface{}) (less func(a, b interface{}) bool, score func(interface{}) float64) {
	switch a.(type) {
	case Fast:
		return func(a, b interface{}) bool { return b.(Fast).Less(a) },
		func(t interface{}) float64 { return -t.(Fast).Score() }
	case Slow:
		return func(a, b interface{}) bool { return b.(Slow).Less(a) },
		func(t interface{}) float64 { return -float64(0.0) }
	case []byte:
		return func(a, b interface{}) bool { return bytes.Compare (a.([]byte), b.([]byte)) > 0 },
		func(t interface{}) float64 { return -scoreBytes(t.([]byte)) }
	case float32:
		return func(a, b interface{}) bool { return a.(float32) > b.(float32) },
		func(t interface{}) float64 { return -float64(t.(float32)) }
	case float64:
		return func(a, b interface{}) bool { return a.(float64) > b.(float64) },
		func(t interface{}) float64 { return -float64(t.(float64)) }
	case int:
		return func(a, b interface{}) bool { return a.(int) > b.(int) },
		func(t interface{}) float64 { return -float64(t.(int)) }
	case int16:
		return func(a, b interface{}) bool { return a.(int16) > b.(int16) },
		func(t interface{}) float64 { return -float64(t.(int16)) }
	case int32:
		return func(a, b interface{}) bool { return a.(int32) > b.(int32) },
		func(t interface{}) float64 { return -float64(t.(int32)) }
	case int64:
		return func(a, b interface{}) bool { return a.(int64) > b.(int64) },
		func(t interface{}) float64 { return -float64(t.(int64)) }
	case int8:
		return func(a, b interface{}) bool { return a.(int8) > b.(int8) },
		func(t interface{}) float64 { return -float64(t.(int8)) }
	case string:
		return func(a, b interface{}) bool { return a.(string) > b.(string) },
		func(t interface{}) float64 { return -scoreBytes([]byte(t.(string))) }
	case uint:
		return func(a, b interface{}) bool { return a.(uint) > b.(uint) },
		func(t interface{}) float64 { return -float64(t.(uint)) }
	case uint16:
		return func(a, b interface{}) bool { return a.(uint16) > b.(uint16) },
		func(t interface{}) float64 { return -float64(t.(uint16)) }
	case uint32:
		return func(a, b interface{}) bool { return a.(uint32) > b.(uint32) },
		func(t interface{}) float64 { return -float64(t.(uint32)) }
	case uint64:
		return func(a, b interface{}) bool { return a.(uint64) > b.(uint64) },
		func(t interface{}) float64 { return -float64(t.(uint64)) }
	case uint8:
		return func(a, b interface{}) bool { return a.(uint8) > b.(uint8) },
		func(t interface{}) float64 { return -float64(t.(uint8)) }
	case uintptr:
		return func(a, b interface{}) bool { return a.(uintptr) > b.(uintptr) },
		func(t interface{}) float64 { return -float64(t.(uintptr)) }
	}
	panic(fmt.Sprintf("cannot order type %T: no method Less(a, b interface{}) bool", a))
}

// FnScore returns a score for a single value and a comparison
// function that can be used to sort similar values.  It is typically
// used by container methods for containers that do not cache the less
// and score functions.
// 
func FnScore(value interface{}) (less func(a, b interface{}) bool, score float64) {
	switch v := value.(type) {
	case Fast:
		return func(a, b interface{}) bool { return a.(Fast).Less(b) },
		v.Score()
	case Slow:
		return func(a, b interface{}) bool { return a.(Slow).Less(b) },
		float64(0.0)
	case []byte:
		return func(a, b interface{}) bool { return bytes.Compare (a.([]byte), b.([]byte)) < 0 },
		scoreBytes(v)
	case float32:
		return func(a, b interface{}) bool { return a.(float32) < b.(float32) },
		float64(v)
	case float64:
		return func(a, b interface{}) bool { return a.(float64) < b.(float64) },
		float64(v)
	case int:
		return func(a, b interface{}) bool { return a.(int) < b.(int) },
		float64(v)
	case int16:
		return func(a, b interface{}) bool { return a.(int16) < b.(int16) },
		float64(v)
	case int32:
		return func(a, b interface{}) bool { return a.(int32) < b.(int32) },
		float64(v)
	case int64:
		return func(a, b interface{}) bool { return a.(int64) < b.(int64) },
		float64(v)
	case int8:
		return func(a, b interface{}) bool { return a.(int8) < b.(int8) },
		float64(v)
	case string:
		return func(a, b interface{}) bool { return a.(string) < b.(string) },
		scoreBytes([]byte(v))
	case uint:
		return func(a, b interface{}) bool { return a.(uint) < b.(uint) },
		float64(v)
	case uint16:
		return func(a, b interface{}) bool { return a.(uint16) < b.(uint16) },
		float64(v)
	case uint32:
		return func(a, b interface{}) bool { return a.(uint32) < b.(uint32) },
		float64(v)
	case uint64:
		return func(a, b interface{}) bool { return a.(uint64) < b.(uint64) },
		float64(v)
	case uint8:
		return func(a, b interface{}) bool { return a.(uint8) < b.(uint8) },
		float64(v)
	case uintptr:
		return func(a, b interface{}) bool { return a.(uintptr) < b.(uintptr) },
		float64(v)
	}
	panic(fmt.Sprintf("cannot order type %T: no method Less(a, b interface{}) bool", value))
}

// FnScoreReversed is like FnScore but for sorting in reverse order.
// 
func FnScoreReversed(value interface{}) (less func(a, b interface{}) bool, score float64) {
	switch v := value.(type) {
	case Fast:
		return func(a, b interface{}) bool { return b.(Fast).Less(a) },
		-v.Score()
	case Slow:
		return func(a, b interface{}) bool { return b.(Slow).Less(a) },
		-float64(0.0)
	case []byte:
		return func(a, b interface{}) bool { return bytes.Compare (a.([]byte), b.([]byte)) > 0 },
		-scoreBytes(v)
	case float32:
		return func(a, b interface{}) bool { return a.(float32) > b.(float32) },
		-float64(v)
	case float64:
		return func(a, b interface{}) bool { return a.(float64) > b.(float64) },
		-float64(v)
	case int:
		return func(a, b interface{}) bool { return a.(int) > b.(int) },
		-float64(v)
	case int16:
		return func(a, b interface{}) bool { return a.(int16) > b.(int16) },
		-float64(v)
	case int32:
		return func(a, b interface{}) bool { return a.(int32) > b.(int32) },
		-float64(v)
	case int64:
		return func(a, b interface{}) bool { return a.(int64) > b.(int64) },
		-float64(v)
	case int8:
		return func(a, b interface{}) bool { return a.(int8) > b.(int8) },
		-float64(v)
	case string:
		return func(a, b interface{}) bool { return a.(string) > b.(string) },
		-scoreBytes([]byte(v))
	case uint:
		return func(a, b interface{}) bool { return a.(uint) > b.(uint) },
		-float64(v)
	case uint16:
		return func(a, b interface{}) bool { return a.(uint16) > b.(uint16) },
		-float64(v)
	case uint32:
		return func(a, b interface{}) bool { return a.(uint32) > b.(uint32) },
		-float64(v)
	case uint64:
		return func(a, b interface{}) bool { return a.(uint64) > b.(uint64) },
		-float64(v)
	case uint8:
		return func(a, b interface{}) bool { return a.(uint8) > b.(uint8) },
		-float64(v)
	case uintptr:
		return func(a, b interface{}) bool { return a.(uintptr) > b.(uintptr) },
		-float64(v)
	}
	panic(fmt.Sprintf("cannot order type %T: no method Less(a, b interface{}) bool", value))
}

// Use up to 6 bytes to generate a score.
func scoreBytes(data []byte) float64 {
	l := uint(len(data))
	if l > 6 {
		l = 6
	}
	var result uint64
	for i := uint(0); i < l; i++ {
		result |= uint64(data[i]) << ((5 - i) * 8)
	}
	return float64(result)
}
