// this file was generated by gomacro command: import "expvar"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "expvar"
	. "reflect"
)

func init() {
	Binds["expvar"] = map[string]Value{
		"Do":	ValueOf(pkg.Do),
		"Get":	ValueOf(pkg.Get),
		"Handler":	ValueOf(pkg.Handler),
		"NewFloat":	ValueOf(pkg.NewFloat),
		"NewInt":	ValueOf(pkg.NewInt),
		"NewMap":	ValueOf(pkg.NewMap),
		"NewString":	ValueOf(pkg.NewString),
		"Publish":	ValueOf(pkg.Publish),
	}
	Types["expvar"] = map[string]Type{
		"Float":	TypeOf((*pkg.Float)(nil)).Elem(),
		"Func":	TypeOf((*pkg.Func)(nil)).Elem(),
		"Int":	TypeOf((*pkg.Int)(nil)).Elem(),
		"KeyValue":	TypeOf((*pkg.KeyValue)(nil)).Elem(),
		"Map":	TypeOf((*pkg.Map)(nil)).Elem(),
		"String":	TypeOf((*pkg.String)(nil)).Elem(),
		"Var":	TypeOf((*pkg.Var)(nil)).Elem(),
	}
}
