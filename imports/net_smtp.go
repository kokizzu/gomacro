// this file was generated by gomacro command: import "net/smtp"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "net/smtp"
	. "reflect"
)

func init() {
	Binds["net/smtp"] = map[string]Value{
		"CRAMMD5Auth":	ValueOf(pkg.CRAMMD5Auth),
		"Dial":	ValueOf(pkg.Dial),
		"NewClient":	ValueOf(pkg.NewClient),
		"PlainAuth":	ValueOf(pkg.PlainAuth),
		"SendMail":	ValueOf(pkg.SendMail),
	}
	Types["net/smtp"] = map[string]Type{
		"Auth":	TypeOf((*pkg.Auth)(nil)).Elem(),
		"Client":	TypeOf((*pkg.Client)(nil)).Elem(),
		"ServerInfo":	TypeOf((*pkg.ServerInfo)(nil)).Elem(),
	}
}
