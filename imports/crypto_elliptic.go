// this file was generated by gomacro command: import "crypto/elliptic"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "crypto/elliptic"
	. "reflect"
)

func init() {
	Binds["crypto/elliptic"] = map[string]Value{
		"GenerateKey":	ValueOf(pkg.GenerateKey),
		"Marshal":	ValueOf(pkg.Marshal),
		"P224":	ValueOf(pkg.P224),
		"P256":	ValueOf(pkg.P256),
		"P384":	ValueOf(pkg.P384),
		"P521":	ValueOf(pkg.P521),
		"Unmarshal":	ValueOf(pkg.Unmarshal),
	}
	Types["crypto/elliptic"] = map[string]Type{
		"Curve":	TypeOf((*pkg.Curve)(nil)).Elem(),
		"CurveParams":	TypeOf((*pkg.CurveParams)(nil)).Elem(),
	}
}
