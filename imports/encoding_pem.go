// this file was generated by gomacro command: import "encoding/pem"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "encoding/pem"
	. "reflect"
)

func init() {
	Binds["encoding/pem"] = map[string]Value{
		"Decode":	ValueOf(pkg.Decode),
		"Encode":	ValueOf(pkg.Encode),
		"EncodeToMemory":	ValueOf(pkg.EncodeToMemory),
	}
	Types["encoding/pem"] = map[string]Type{
		"Block":	TypeOf((*pkg.Block)(nil)).Elem(),
	}
}
