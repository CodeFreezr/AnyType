// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// from go/types/type.go

// BasicKind describes the kind of basic type.
type BasicKind int

// BasicInfo is a set of flags describing properties of a basic type.
type BasicInfo int

// Properties of basic types.
const (
	IsBoolean BasicInfo = 1 << iota
	IsInteger
	IsUnsigned
	IsFloat
	IsComplex
	IsString
	IsUntyped

	IsOrdered   = IsInteger | IsFloat | IsString
	IsNumeric   = IsInteger | IsFloat | IsComplex
	IsConstType = IsBoolean | IsNumeric | IsString
)

var _ BasicInfo = IsOrdered // allows to use "Merge"
// token.LSS, token.LEQ, token.GTR, token.GEQ

/* Quote from "The Go Programming Language" Chapter 3 - page 52:
The type rune is a synonym for int32 and conventionally indicates that a value is a Unicode code point. The two names may be used interchangeably.

Similarly, the type byte is a synonym for uint8, and emphasizes that the value is a piece of raw data rather than a small numeric quantity.

Finally, there is an unsigned integer type uintptr, whose width is not specified but is sufficient to hold all the bits of a pointer value.
The uintptr type is used only for low-level programming, such as at the boundary of a Go program with a C library or an operating system.
*/