// Do not edit. Bootstrap copy of /usr/local/google/buildbot/src/android/build-tools/out/obj/go/src/cmd/internal/obj/fmt.go

//line /usr/local/google/buildbot/src/android/build-tools/out/obj/go/src/cmd/internal/obj/fmt.go:1
/*
 * The authors of this software are Rob Pike and Ken Thompson.
 *              Copyright (c) 2002 by Lucent Technologies.
 * Permission to use, copy, modify, and distribute this software for any
 * purpose without fee is hereby granted, provided that this entire notice
 * is included in all copies of any software which is or includes a copy
 * or modification of this software and in all copies of the supporting
 * documentation for such software.
 * THIS SOFTWARE IS BEING PROVIDED "AS IS", WITHOUT ANY EXPRESS OR IMPLIED
 * WARRANTY.  IN PARTICULAR, NEITHER THE AUTHORS NOR LUCENT TECHNOLOGIES MAKE ANY
 * REPRESENTATION OR WARRANTY OF ANY KIND CONCERNING THE MERCHANTABILITY
 * OF THIS SOFTWARE OR ITS FITNESS FOR ANY PARTICULAR PURPOSE.
 */

package obj

const (
	FmtWidth = 1 << iota
	FmtLeft
	FmtPrec
	FmtSharp
	FmtSpace
	FmtSign
	FmtApost
	FmtZero
	FmtUnsigned
	FmtShort
	FmtLong
	FmtVLong
	FmtComma
	FmtByte
	FmtLDouble
	FmtFlag
)