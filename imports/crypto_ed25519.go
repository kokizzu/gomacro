// this file was generated by gomacro command: import _b "crypto/ed25519"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	ed25519 "crypto/ed25519"
)

// reflection: allow interpreted code to import "crypto/ed25519"
func init() {
	Packages["crypto/ed25519"] = Package{
		Name: "ed25519",
		Binds: map[string]Value{
			"GenerateKey":	ValueOf(ed25519.GenerateKey),
			"NewKeyFromSeed":	ValueOf(ed25519.NewKeyFromSeed),
			"PrivateKeySize":	ValueOf(ed25519.PrivateKeySize),
			"PublicKeySize":	ValueOf(ed25519.PublicKeySize),
			"SeedSize":	ValueOf(ed25519.SeedSize),
			"Sign":	ValueOf(ed25519.Sign),
			"SignatureSize":	ValueOf(ed25519.SignatureSize),
			"Verify":	ValueOf(ed25519.Verify),
		}, Types: map[string]Type{
			"PrivateKey":	TypeOf((*ed25519.PrivateKey)(nil)).Elem(),
			"PublicKey":	TypeOf((*ed25519.PublicKey)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"PrivateKeySize":	"int:64",
			"PublicKeySize":	"int:32",
			"SeedSize":	"int:32",
			"SignatureSize":	"int:64",
		}, 
	}
}
