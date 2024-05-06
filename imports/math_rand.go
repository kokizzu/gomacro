// this file was generated by gomacro command: import _b "math/rand"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	rand "math/rand"
)

// reflection: allow interpreted code to import "math/rand"
func init() {
	Packages["math/rand"] = Package{
		Name: "rand",
		Binds: map[string]Value{
			"ExpFloat64":	ValueOf(rand.ExpFloat64),
			"Float32":	ValueOf(rand.Float32),
			"Float64":	ValueOf(rand.Float64),
			"Int":	ValueOf(rand.Int),
			"Int31":	ValueOf(rand.Int31),
			"Int31n":	ValueOf(rand.Int31n),
			"Int63":	ValueOf(rand.Int63),
			"Int63n":	ValueOf(rand.Int63n),
			"Intn":	ValueOf(rand.Intn),
			"New":	ValueOf(rand.New),
			"NewSource":	ValueOf(rand.NewSource),
			"NewZipf":	ValueOf(rand.NewZipf),
			"NormFloat64":	ValueOf(rand.NormFloat64),
			"Perm":	ValueOf(rand.Perm),
			"Read":	ValueOf(rand.Read),
			"Seed":	ValueOf(rand.Seed),
			"Shuffle":	ValueOf(rand.Shuffle),
			"Uint32":	ValueOf(rand.Uint32),
			"Uint64":	ValueOf(rand.Uint64),
		}, Types: map[string]Type{
			"Rand":	TypeOf((*rand.Rand)(nil)).Elem(),
			"Source":	TypeOf((*rand.Source)(nil)).Elem(),
			"Source64":	TypeOf((*rand.Source64)(nil)).Elem(),
			"Zipf":	TypeOf((*rand.Zipf)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"Source":	TypeOf((*P_math_rand_Source)(nil)).Elem(),
			"Source64":	TypeOf((*P_math_rand_Source64)(nil)).Elem(),
		}, 
	}
}

// --------------- proxy for math/rand.Source ---------------
type P_math_rand_Source struct {
	Object	interface{}
	Int63_	func(interface{}) int64
	Seed_	func(_proxy_obj_ interface{}, seed int64) 
}
func (P *P_math_rand_Source) Int63() int64 {
	return P.Int63_(P.Object)
}
func (P *P_math_rand_Source) Seed(seed int64)  {
	P.Seed_(P.Object, seed)
}

// --------------- proxy for math/rand.Source64 ---------------
type P_math_rand_Source64 struct {
	Object	interface{}
	Int63_	func(interface{}) int64
	Seed_	func(_proxy_obj_ interface{}, seed int64) 
	Uint64_	func(interface{}) uint64
}
func (P *P_math_rand_Source64) Int63() int64 {
	return P.Int63_(P.Object)
}
func (P *P_math_rand_Source64) Seed(seed int64)  {
	P.Seed_(P.Object, seed)
}
func (P *P_math_rand_Source64) Uint64() uint64 {
	return P.Uint64_(P.Object)
}
