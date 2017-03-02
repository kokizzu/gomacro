// this file was generated by gomacro command: import "flag"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "flag"
	. "reflect"
)

func init() {
	Binds["flag"] = map[string]Value{
		"Arg":	ValueOf(pkg.Arg),
		"Args":	ValueOf(pkg.Args),
		"Bool":	ValueOf(pkg.Bool),
		"BoolVar":	ValueOf(pkg.BoolVar),
		"CommandLine":	ValueOf(&pkg.CommandLine).Elem(),
		"ContinueOnError":	ValueOf(pkg.ContinueOnError),
		"Duration":	ValueOf(pkg.Duration),
		"DurationVar":	ValueOf(pkg.DurationVar),
		"ErrHelp":	ValueOf(&pkg.ErrHelp).Elem(),
		"ExitOnError":	ValueOf(pkg.ExitOnError),
		"Float64":	ValueOf(pkg.Float64),
		"Float64Var":	ValueOf(pkg.Float64Var),
		"Int":	ValueOf(pkg.Int),
		"Int64":	ValueOf(pkg.Int64),
		"Int64Var":	ValueOf(pkg.Int64Var),
		"IntVar":	ValueOf(pkg.IntVar),
		"Lookup":	ValueOf(pkg.Lookup),
		"NArg":	ValueOf(pkg.NArg),
		"NFlag":	ValueOf(pkg.NFlag),
		"NewFlagSet":	ValueOf(pkg.NewFlagSet),
		"PanicOnError":	ValueOf(pkg.PanicOnError),
		"Parse":	ValueOf(pkg.Parse),
		"Parsed":	ValueOf(pkg.Parsed),
		"PrintDefaults":	ValueOf(pkg.PrintDefaults),
		"Set":	ValueOf(pkg.Set),
		"String":	ValueOf(pkg.String),
		"StringVar":	ValueOf(pkg.StringVar),
		"Uint":	ValueOf(pkg.Uint),
		"Uint64":	ValueOf(pkg.Uint64),
		"Uint64Var":	ValueOf(pkg.Uint64Var),
		"UintVar":	ValueOf(pkg.UintVar),
		"UnquoteUsage":	ValueOf(pkg.UnquoteUsage),
		"Usage":	ValueOf(&pkg.Usage).Elem(),
		"Var":	ValueOf(pkg.Var),
		"Visit":	ValueOf(pkg.Visit),
		"VisitAll":	ValueOf(pkg.VisitAll),
	}
	Types["flag"] = map[string]Type{
		"ErrorHandling":	TypeOf((*pkg.ErrorHandling)(nil)).Elem(),
		"Flag":	TypeOf((*pkg.Flag)(nil)).Elem(),
		"FlagSet":	TypeOf((*pkg.FlagSet)(nil)).Elem(),
		"Getter":	TypeOf((*pkg.Getter)(nil)).Elem(),
		"Value":	TypeOf((*pkg.Value)(nil)).Elem(),
	}
}
