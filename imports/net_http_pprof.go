// this file was generated by gomacro command: import _b "net/http/pprof"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	pprof "net/http/pprof"
)

// reflection: allow interpreted code to import "net/http/pprof"
func init() {
	Packages["net/http/pprof"] = Package{
		Name: "pprof",
		Binds: map[string]Value{
			"Cmdline":	ValueOf(pprof.Cmdline),
			"Handler":	ValueOf(pprof.Handler),
			"Index":	ValueOf(pprof.Index),
			"Profile":	ValueOf(pprof.Profile),
			"Symbol":	ValueOf(pprof.Symbol),
			"Trace":	ValueOf(pprof.Trace),
		}, 
	}
}
