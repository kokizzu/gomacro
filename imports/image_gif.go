// this file was generated by gomacro command: import "image/gif"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "image/gif"
	. "reflect"
)

func init() {
	Binds["image/gif"] = map[string]Value{
		"Decode":	ValueOf(pkg.Decode),
		"DecodeAll":	ValueOf(pkg.DecodeAll),
		"DecodeConfig":	ValueOf(pkg.DecodeConfig),
		"DisposalBackground":	ValueOf(pkg.DisposalBackground),
		"DisposalNone":	ValueOf(pkg.DisposalNone),
		"DisposalPrevious":	ValueOf(pkg.DisposalPrevious),
		"Encode":	ValueOf(pkg.Encode),
		"EncodeAll":	ValueOf(pkg.EncodeAll),
	}
	Types["image/gif"] = map[string]Type{
		"GIF":	TypeOf((*pkg.GIF)(nil)).Elem(),
		"Options":	TypeOf((*pkg.Options)(nil)).Elem(),
	}
}
