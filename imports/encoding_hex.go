// this file was generated by gomacro command: import "encoding/hex"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "encoding/hex"
	. "reflect"
)

func init() {
	Binds["encoding/hex"] = map[string]Value{
		"Decode":	ValueOf(pkg.Decode),
		"DecodeString":	ValueOf(pkg.DecodeString),
		"DecodedLen":	ValueOf(pkg.DecodedLen),
		"Dump":	ValueOf(pkg.Dump),
		"Dumper":	ValueOf(pkg.Dumper),
		"Encode":	ValueOf(pkg.Encode),
		"EncodeToString":	ValueOf(pkg.EncodeToString),
		"EncodedLen":	ValueOf(pkg.EncodedLen),
		"ErrLength":	ValueOf(&pkg.ErrLength).Elem(),
	}
	Types["encoding/hex"] = map[string]Type{
		"InvalidByteError":	TypeOf((*pkg.InvalidByteError)(nil)).Elem(),
	}
}
