package awk

import (
	"fmt"
	"strconv"
)

const convFmt = "%.6g"

// A Value represents an immutable datum that can be converted to various
// types in best-effort fashion (i.e., never returning an error).
type Value struct {
	i64 int64   // Value converted to an int64
	f64 float64 // Value converted to a float64
	s   string  // Value converted to a string

	i64_ok bool // true: i64 is valid; false: invalid
	f64_ok bool // true: f64 is valid; false: invalid
	s_ok   bool // true: s is valid; false: invalid

	script *Script // Pointer to the script that produced this value
}

// NewInt64 creates a Value from an int64.
func (s *Script) NewInt64(i int64) *Value {
	return &Value{
		i64:    i,
		i64_ok: true,
		script: s,
	}
}

// NewFloat64 creates a Value from a float64.
func (s *Script) NewFloat64(f float64) *Value {
	return &Value{
		f64:    f,
		f64_ok: true,
		script: s,
	}
}

// NewString creates a Value from a string.
func (s *Script) NewString(str string) *Value {
	return &Value{
		s:      str,
		s_ok:   true,
		script: s,
	}
}

// Int64 converts a Value to an int64.
func (v *Value) Int64() int64 {
	switch {
	case v.f64_ok:
		v.i64 = int64(v.f64)
		v.i64_ok = true
	case v.s_ok:
		v.i64, _ = strconv.ParseInt(v.s, 10, 64)
		v.i64_ok = true
	}
	return v.i64
}

// Float64 converts a Value to a float64.
func (v *Value) Float64() float64 {
	switch {
	case v.i64_ok:
		v.f64 = float64(v.i64)
		v.f64_ok = true
	case v.s_ok:
		v.f64, _ = strconv.ParseFloat(v.s, 64)
		v.f64_ok = true
	}
	return v.f64
}

// String converts a Value to a string.
func (v *Value) String() string {
	switch {
	case v.i64_ok:
		v.s = strconv.FormatInt(v.i64, 10)
		v.s_ok = true
	case v.f64_ok:
		v.s = fmt.Sprintf(v.script.ConvFmt, v.f64)
		v.s_ok = true
	}
	return v.s
}
