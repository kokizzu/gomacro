// this file was generated by gomacro command: import "net/rpc"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "net/rpc"
	. "reflect"
)

func init() {
	Binds["net/rpc"] = map[string]Value{
		"Accept":	ValueOf(pkg.Accept),
		"DefaultDebugPath":	ValueOf(pkg.DefaultDebugPath),
		"DefaultRPCPath":	ValueOf(pkg.DefaultRPCPath),
		"DefaultServer":	ValueOf(&pkg.DefaultServer).Elem(),
		"Dial":	ValueOf(pkg.Dial),
		"DialHTTP":	ValueOf(pkg.DialHTTP),
		"DialHTTPPath":	ValueOf(pkg.DialHTTPPath),
		"ErrShutdown":	ValueOf(&pkg.ErrShutdown).Elem(),
		"HandleHTTP":	ValueOf(pkg.HandleHTTP),
		"NewClient":	ValueOf(pkg.NewClient),
		"NewClientWithCodec":	ValueOf(pkg.NewClientWithCodec),
		"NewServer":	ValueOf(pkg.NewServer),
		"Register":	ValueOf(pkg.Register),
		"RegisterName":	ValueOf(pkg.RegisterName),
		"ServeCodec":	ValueOf(pkg.ServeCodec),
		"ServeConn":	ValueOf(pkg.ServeConn),
		"ServeRequest":	ValueOf(pkg.ServeRequest),
	}
	Types["net/rpc"] = map[string]Type{
		"Call":	TypeOf((*pkg.Call)(nil)).Elem(),
		"Client":	TypeOf((*pkg.Client)(nil)).Elem(),
		"ClientCodec":	TypeOf((*pkg.ClientCodec)(nil)).Elem(),
		"Request":	TypeOf((*pkg.Request)(nil)).Elem(),
		"Response":	TypeOf((*pkg.Response)(nil)).Elem(),
		"Server":	TypeOf((*pkg.Server)(nil)).Elem(),
		"ServerCodec":	TypeOf((*pkg.ServerCodec)(nil)).Elem(),
		"ServerError":	TypeOf((*pkg.ServerError)(nil)).Elem(),
	}
}
