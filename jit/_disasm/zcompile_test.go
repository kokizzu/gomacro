/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * zcompile_test.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"testing"

	. "github.com/cosmos72/gomacro/jit"
	"github.com/cosmos72/gomacro/jit/asm"
)

const (
	S0 SoftRegId = iota
	S1
)

func SameCode(actual Code, expected Code) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := range actual {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}

func TestCompileExpr1(t *testing.T) {
	var c Comp
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		r := MakeReg(c.RLo, Uint64)
		c.Expr(
			NewExpr1(
				NEG, NewExpr1(NOT, r),
			),
		)
		actual := c.Code()
		t.Log(actual...)

		expected := Code{
			asm.ALLOC, S0, Uint64,
			asm.NOT2, r, S0,
			asm.NEG2, S0, S0,
		}

		if !SameCode(actual, expected) {
			t.Errorf("miscompiled code:\n\texpected %v\n\tactual   %v",
				expected, actual)
		}

		// assemble
		a := c.NewAsm()
		a.Asm(c.Code()...)
		a.Epilogue()
		PrintDisasm(t, c.ArchId(), a.Code())
	}
}

func TestCompileExpr2(t *testing.T) {
	var c Comp
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		a := c.NewAsm()

		c7 := MakeConst(7, Uint64)
		c9 := MakeConst(9, Uint64)
		r1 := a.RegAlloc(Uint64)
		r2 := a.RegAlloc(Uint64)
		// compile
		c.Expr(
			NewExpr2(
				ADD, NewExpr2(MUL, c7, r1), NewExpr2(SUB, c9, r2),
			),
		)
		actual := c.Code()
		t.Log(actual...)

		expected := Code{
			asm.ALLOC, S0, Uint64,
			asm.MUL3, c7, r1, S0,
			asm.ALLOC, S1, Uint64,
			asm.SUB3, c9, r2, S1,
			asm.ADD3, S0, S1, S0,
			asm.FREE, S1, asm.Uint64,
		}

		if !SameCode(actual, expected) {
			t.Errorf("miscompiled code:\n\texpected %v\n\tactual   %v",
				expected, actual)
		}

		// assemble
		a.Asm(c.Code()...)
		a.Epilogue()
		PrintDisasm(t, c.ArchId(), a.Code())
	}
}