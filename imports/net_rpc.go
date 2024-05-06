// this file was generated by gomacro command: import _b "net/rpc"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	rpc "net/rpc"
)

// reflection: allow interpreted code to import "net/rpc"
func init() {
	Packages["net/rpc"] = Package{
		Name: "rpc",
		Binds: map[string]Value{
			"Accept":	ValueOf(rpc.Accept),
			"DefaultDebugPath":	ValueOf(rpc.DefaultDebugPath),
			"DefaultRPCPath":	ValueOf(rpc.DefaultRPCPath),
			"DefaultServer":	ValueOf(&rpc.DefaultServer).Elem(),
			"Dial":	ValueOf(rpc.Dial),
			"DialHTTP":	ValueOf(rpc.DialHTTP),
			"DialHTTPPath":	ValueOf(rpc.DialHTTPPath),
			"ErrShutdown":	ValueOf(&rpc.ErrShutdown).Elem(),
			"HandleHTTP":	ValueOf(rpc.HandleHTTP),
			"NewClient":	ValueOf(rpc.NewClient),
			"NewClientWithCodec":	ValueOf(rpc.NewClientWithCodec),
			"NewServer":	ValueOf(rpc.NewServer),
			"Register":	ValueOf(rpc.Register),
			"RegisterName":	ValueOf(rpc.RegisterName),
			"ServeCodec":	ValueOf(rpc.ServeCodec),
			"ServeConn":	ValueOf(rpc.ServeConn),
			"ServeRequest":	ValueOf(rpc.ServeRequest),
		}, Types: map[string]Type{
			"Call":	TypeOf((*rpc.Call)(nil)).Elem(),
			"Client":	TypeOf((*rpc.Client)(nil)).Elem(),
			"ClientCodec":	TypeOf((*rpc.ClientCodec)(nil)).Elem(),
			"Request":	TypeOf((*rpc.Request)(nil)).Elem(),
			"Response":	TypeOf((*rpc.Response)(nil)).Elem(),
			"Server":	TypeOf((*rpc.Server)(nil)).Elem(),
			"ServerCodec":	TypeOf((*rpc.ServerCodec)(nil)).Elem(),
			"ServerError":	TypeOf((*rpc.ServerError)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"ClientCodec":	TypeOf((*P_net_rpc_ClientCodec)(nil)).Elem(),
			"ServerCodec":	TypeOf((*P_net_rpc_ServerCodec)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"DefaultDebugPath":	"string:/debug/rpc",
			"DefaultRPCPath":	"string:/_goRPC_",
		}, 
	}
}

// --------------- proxy for net/rpc.ClientCodec ---------------
type P_net_rpc_ClientCodec struct {
	Object	interface{}
	Close_	func(interface{}) error
	ReadResponseBody_	func(interface{}, any) error
	ReadResponseHeader_	func(interface{}, *rpc.Response) error
	WriteRequest_	func(interface{}, *rpc.Request, any) error
}
func (P *P_net_rpc_ClientCodec) Close() error {
	return P.Close_(P.Object)
}
func (P *P_net_rpc_ClientCodec) ReadResponseBody(unnamed0 any) error {
	return P.ReadResponseBody_(P.Object, unnamed0)
}
func (P *P_net_rpc_ClientCodec) ReadResponseHeader(unnamed0 *rpc.Response) error {
	return P.ReadResponseHeader_(P.Object, unnamed0)
}
func (P *P_net_rpc_ClientCodec) WriteRequest(unnamed0 *rpc.Request, unnamed1 any) error {
	return P.WriteRequest_(P.Object, unnamed0, unnamed1)
}

// --------------- proxy for net/rpc.ServerCodec ---------------
type P_net_rpc_ServerCodec struct {
	Object	interface{}
	Close_	func(interface{}) error
	ReadRequestBody_	func(interface{}, any) error
	ReadRequestHeader_	func(interface{}, *rpc.Request) error
	WriteResponse_	func(interface{}, *rpc.Response, any) error
}
func (P *P_net_rpc_ServerCodec) Close() error {
	return P.Close_(P.Object)
}
func (P *P_net_rpc_ServerCodec) ReadRequestBody(unnamed0 any) error {
	return P.ReadRequestBody_(P.Object, unnamed0)
}
func (P *P_net_rpc_ServerCodec) ReadRequestHeader(unnamed0 *rpc.Request) error {
	return P.ReadRequestHeader_(P.Object, unnamed0)
}
func (P *P_net_rpc_ServerCodec) WriteResponse(unnamed0 *rpc.Response, unnamed1 any) error {
	return P.WriteResponse_(P.Object, unnamed0, unnamed1)
}
