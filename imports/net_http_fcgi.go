// this file was generated by gomacro command: import "net/http/fcgi"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "net/http/fcgi"
	. "reflect"
)

func init() {
	Binds["net/http/fcgi"] = map[string]Value{
		"ErrConnClosed":	ValueOf(&pkg.ErrConnClosed).Elem(),
		"ErrRequestAborted":	ValueOf(&pkg.ErrRequestAborted).Elem(),
		"Serve":	ValueOf(pkg.Serve),
	}
	Types["net/http/fcgi"] = map[string]Type{
	}
}
