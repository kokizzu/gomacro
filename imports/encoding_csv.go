// this file was generated by gomacro command: import "encoding/csv"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "encoding/csv"
	. "reflect"
)

func init() {
	Binds["encoding/csv"] = map[string]Value{
		"ErrBareQuote":	ValueOf(&pkg.ErrBareQuote).Elem(),
		"ErrFieldCount":	ValueOf(&pkg.ErrFieldCount).Elem(),
		"ErrQuote":	ValueOf(&pkg.ErrQuote).Elem(),
		"ErrTrailingComma":	ValueOf(&pkg.ErrTrailingComma).Elem(),
		"NewReader":	ValueOf(pkg.NewReader),
		"NewWriter":	ValueOf(pkg.NewWriter),
	}
	Types["encoding/csv"] = map[string]Type{
		"ParseError":	TypeOf((*pkg.ParseError)(nil)).Elem(),
		"Reader":	TypeOf((*pkg.Reader)(nil)).Elem(),
		"Writer":	TypeOf((*pkg.Writer)(nil)).Elem(),
	}
}
