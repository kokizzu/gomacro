/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * call.go
 *
 *  Created on: Feb 15, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	r "reflect"
)

func (env *Env) evalCall(node *ast.CallExpr) (r.Value, []r.Value) {
	fun, _ := env.evalExpr(node.Fun)
	if fun.Kind() != r.Func {
		return env.Errorf("call of non-function %#v", node)
	}
	// TODO support the special case fooAcceptsMultipleArgs( barReturnsMultipleValues() )
	args := env.evalExprs(node.Args)
	if !fun.Type().IsVariadic() {
		argTypes := fun.Type()
		for i, arg := range args {
			args[i] = env.toType(arg, argTypes.In(i))
		}
	}
	rets := fun.Call(args)
	return unpackValues(rets)
}
