// +build amd64

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
 * asm_amd64_op1.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) Op1(op Op1, a Arg) *Asm {
	switch a := a.(type) {
	case Reg:
		asm.op1Reg(op, a)
	case Mem:
		asm.op1Mem(op, a)
	case Const:
		errorf("destination cannot be a constant: %v %v", op, a)
	default:
		errorf("unsupported destination type %T, expecting Reg or Mem: %v %v", a, op, a)
	}
	return asm
}

// OP %reg_dst
func (asm *Asm) op1Reg(op Op1, r Reg) *Asm {
	rlo, rhi := r.lohi()
	oplo, ophi := op.lohi()

	switch SizeOf(r) {
	case 1:
		if r.id >= RSP {
			asm.Bytes(0x40 | rhi)
		}
		asm.Bytes(0xF6|ophi, 0xC0|oplo|rlo)
	case 2:
		asm.Bytes(0x66)
		fallthrough
	case 4:
		if rhi != 0 {
			asm.Bytes(0x41)
		}
		asm.Bytes(0xF7|ophi, 0xC0|oplo|rlo)
	case 8:
		asm.Bytes(0x48|rhi, 0xF7|ophi, 0xC0|oplo|rlo)
	default:
		errorf("unsupported register size %v, expecting 1,2,4 or 8 bytes: %v %v", SizeOf(r), op, r)
	}
	return asm
}

// OP off_m(%reg_m)
func (asm *Asm) op1Mem(op Op1, m Mem) *Asm {

	r := m.reg
	rlo, dhi := r.lohi()
	oplo, ophi := op.lohi()

	offlen, offbit := m.offlen(r.id)

	switch SizeOf(m) {
	case 1:
		if dhi != 0 {
			asm.Bytes(0x41)
		}
		asm.Bytes(0xF6|ophi, offbit|oplo|rlo)
	case 2:
		asm.Bytes(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Bytes(0x41)
		}
		asm.Bytes(0xF7|ophi, offbit|oplo|rlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xF7|ophi, offbit|oplo|rlo)
	default:
		errorf("unsupported memory size %v, expecting 1,2,4 or 8 bytes: %v %v", SizeOf(m), op, m)
	}
	asm.quirk24(r)
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}
	return asm
}