// this file was generated by gomacro command: import _b "debug/gosym"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	gosym "debug/gosym"
)

// reflection: allow interpreted code to import "debug/gosym"
func init() {
	Packages["debug/gosym"] = Package{
		Name: "gosym",
		Binds: map[string]Value{
			"NewLineTable":	ValueOf(gosym.NewLineTable),
			"NewTable":	ValueOf(gosym.NewTable),
		}, Types: map[string]Type{
			"DecodingError":	TypeOf((*gosym.DecodingError)(nil)).Elem(),
			"Func":	TypeOf((*gosym.Func)(nil)).Elem(),
			"LineTable":	TypeOf((*gosym.LineTable)(nil)).Elem(),
			"Obj":	TypeOf((*gosym.Obj)(nil)).Elem(),
			"Sym":	TypeOf((*gosym.Sym)(nil)).Elem(),
			"Table":	TypeOf((*gosym.Table)(nil)).Elem(),
			"UnknownFileError":	TypeOf((*gosym.UnknownFileError)(nil)).Elem(),
			"UnknownLineError":	TypeOf((*gosym.UnknownLineError)(nil)).Elem(),
		}, Wrappers: map[string][]string{
			"Func":	[]string{"BaseName","PackageName","ReceiverName","Static",},
		}, 
	}
}
