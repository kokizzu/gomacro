// this file was generated by gomacro command: import "container/ring"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "container/ring"
	. "reflect"
)

func init() {
	Binds["container/ring"] = map[string]Value{
		"New":	ValueOf(pkg.New),
	}
	Types["container/ring"] = map[string]Type{
		"Ring":	TypeOf((*pkg.Ring)(nil)).Elem(),
	}
}
