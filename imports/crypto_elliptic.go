// this file was generated by gomacro command: import _b "crypto/elliptic"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	big "math/big"
	elliptic "crypto/elliptic"
)

// reflection: allow interpreted code to import "crypto/elliptic"
func init() {
	Packages["crypto/elliptic"] = Package{
		Name: "elliptic",
		Binds: map[string]Value{
			"GenerateKey":	ValueOf(elliptic.GenerateKey),
			"Marshal":	ValueOf(elliptic.Marshal),
			"MarshalCompressed":	ValueOf(elliptic.MarshalCompressed),
			"P224":	ValueOf(elliptic.P224),
			"P256":	ValueOf(elliptic.P256),
			"P384":	ValueOf(elliptic.P384),
			"P521":	ValueOf(elliptic.P521),
			"Unmarshal":	ValueOf(elliptic.Unmarshal),
			"UnmarshalCompressed":	ValueOf(elliptic.UnmarshalCompressed),
		}, Types: map[string]Type{
			"Curve":	TypeOf((*elliptic.Curve)(nil)).Elem(),
			"CurveParams":	TypeOf((*elliptic.CurveParams)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"Curve":	TypeOf((*P_crypto_elliptic_Curve)(nil)).Elem(),
		}, 
	}
}

// --------------- proxy for crypto/elliptic.Curve ---------------
type P_crypto_elliptic_Curve struct {
	Object	interface{}
	Add_	func(_proxy_obj_ interface{}, x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (x *big.Int, y *big.Int)
	Double_	func(_proxy_obj_ interface{}, x1 *big.Int, y1 *big.Int) (x *big.Int, y *big.Int)
	IsOnCurve_	func(_proxy_obj_ interface{}, x *big.Int, y *big.Int) bool
	Params_	func(interface{}) *elliptic.CurveParams
	ScalarBaseMult_	func(_proxy_obj_ interface{}, k []byte) (x *big.Int, y *big.Int)
	ScalarMult_	func(_proxy_obj_ interface{}, x1 *big.Int, y1 *big.Int, k []byte) (x *big.Int, y *big.Int)
}
func (P *P_crypto_elliptic_Curve) Add(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (x *big.Int, y *big.Int) {
	return P.Add_(P.Object, x1, y1, x2, y2)
}
func (P *P_crypto_elliptic_Curve) Double(x1 *big.Int, y1 *big.Int) (x *big.Int, y *big.Int) {
	return P.Double_(P.Object, x1, y1)
}
func (P *P_crypto_elliptic_Curve) IsOnCurve(x *big.Int, y *big.Int) bool {
	return P.IsOnCurve_(P.Object, x, y)
}
func (P *P_crypto_elliptic_Curve) Params() *elliptic.CurveParams {
	return P.Params_(P.Object)
}
func (P *P_crypto_elliptic_Curve) ScalarBaseMult(k []byte) (x *big.Int, y *big.Int) {
	return P.ScalarBaseMult_(P.Object, k)
}
func (P *P_crypto_elliptic_Curve) ScalarMult(x1 *big.Int, y1 *big.Int, k []byte) (x *big.Int, y *big.Int) {
	return P.ScalarMult_(P.Object, x1, y1, k)
}
