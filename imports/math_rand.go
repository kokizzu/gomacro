// this file was generated by gomacro command: import "math/rand"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "math/rand"
	. "reflect"
)

func init() {
	Binds["math/rand"] = map[string]Value{
		"ExpFloat64":	ValueOf(pkg.ExpFloat64),
		"Float32":	ValueOf(pkg.Float32),
		"Float64":	ValueOf(pkg.Float64),
		"Int":	ValueOf(pkg.Int),
		"Int31":	ValueOf(pkg.Int31),
		"Int31n":	ValueOf(pkg.Int31n),
		"Int63":	ValueOf(pkg.Int63),
		"Int63n":	ValueOf(pkg.Int63n),
		"Intn":	ValueOf(pkg.Intn),
		"New":	ValueOf(pkg.New),
		"NewSource":	ValueOf(pkg.NewSource),
		"NewZipf":	ValueOf(pkg.NewZipf),
		"NormFloat64":	ValueOf(pkg.NormFloat64),
		"Perm":	ValueOf(pkg.Perm),
		"Read":	ValueOf(pkg.Read),
		"Seed":	ValueOf(pkg.Seed),
		"Uint32":	ValueOf(pkg.Uint32),
		"Uint64":	ValueOf(pkg.Uint64),
	}
	Types["math/rand"] = map[string]Type{
		"Rand":	TypeOf((*pkg.Rand)(nil)).Elem(),
		"Source":	TypeOf((*pkg.Source)(nil)).Elem(),
		"Source64":	TypeOf((*pkg.Source64)(nil)).Elem(),
		"Zipf":	TypeOf((*pkg.Zipf)(nil)).Elem(),
	}
}
