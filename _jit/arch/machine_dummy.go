// +build !amd64,!arm64

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
 * machine_dummy.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

const SUPPORTED = false

type Op0 struct{}
type Op1 struct{}
type Op2 struct{}
type Op3 struct{}
type Op4 struct{}

func (asm *Asm) Op0(op Op0) *Asm {
	return asm
}

func (asm *Asm) Op1(op Op1, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op2(op Op2, dst Arg, src Arg) *Asm {
	return asm
}

func (asm *Asm) Op3(op Op3, dst Arg, a Arg, b Arg) *Asm {
	return asm
}

func (asm *Asm) Op4(op Op4, dst Arg, a Arg, b Arg, c Arg) *Asm {
	return asm
}

const (
	NoRegId RegId = iota
	RLo           = NoRegId
	RHi           = NoRegId
)

func (r RegId) Valid() bool {
	return false
}

var alwaysLiveHwRegs = Regs{}
