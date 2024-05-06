// this file was generated by gomacro command: import _b "log"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	log "log"
)

// reflection: allow interpreted code to import "log"
func init() {
	Packages["log"] = Package{
		Name: "log",
		Binds: map[string]Value{
			"Default":	ValueOf(log.Default),
			"Fatal":	ValueOf(log.Fatal),
			"Fatalf":	ValueOf(log.Fatalf),
			"Fatalln":	ValueOf(log.Fatalln),
			"Flags":	ValueOf(log.Flags),
			"LUTC":	ValueOf(log.LUTC),
			"Ldate":	ValueOf(log.Ldate),
			"Llongfile":	ValueOf(log.Llongfile),
			"Lmicroseconds":	ValueOf(log.Lmicroseconds),
			"Lmsgprefix":	ValueOf(log.Lmsgprefix),
			"Lshortfile":	ValueOf(log.Lshortfile),
			"LstdFlags":	ValueOf(log.LstdFlags),
			"Ltime":	ValueOf(log.Ltime),
			"New":	ValueOf(log.New),
			"Output":	ValueOf(log.Output),
			"Panic":	ValueOf(log.Panic),
			"Panicf":	ValueOf(log.Panicf),
			"Panicln":	ValueOf(log.Panicln),
			"Prefix":	ValueOf(log.Prefix),
			"Print":	ValueOf(log.Print),
			"Printf":	ValueOf(log.Printf),
			"Println":	ValueOf(log.Println),
			"SetFlags":	ValueOf(log.SetFlags),
			"SetOutput":	ValueOf(log.SetOutput),
			"SetPrefix":	ValueOf(log.SetPrefix),
			"Writer":	ValueOf(log.Writer),
		}, Types: map[string]Type{
			"Logger":	TypeOf((*log.Logger)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"LUTC":	"int:32",
			"Ldate":	"int:1",
			"Llongfile":	"int:8",
			"Lmicroseconds":	"int:4",
			"Lmsgprefix":	"int:64",
			"Lshortfile":	"int:16",
			"LstdFlags":	"int:3",
			"Ltime":	"int:2",
		}, 
	}
}
