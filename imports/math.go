// this file was generated by gomacro command: import _b "math"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	math "math"
)

// reflection: allow interpreted code to import "math"
func init() {
	Packages["math"] = Package{
		Name: "math",
		Binds: map[string]Value{
			"Abs":	ValueOf(math.Abs),
			"Acos":	ValueOf(math.Acos),
			"Acosh":	ValueOf(math.Acosh),
			"Asin":	ValueOf(math.Asin),
			"Asinh":	ValueOf(math.Asinh),
			"Atan":	ValueOf(math.Atan),
			"Atan2":	ValueOf(math.Atan2),
			"Atanh":	ValueOf(math.Atanh),
			"Cbrt":	ValueOf(math.Cbrt),
			"Ceil":	ValueOf(math.Ceil),
			"Copysign":	ValueOf(math.Copysign),
			"Cos":	ValueOf(math.Cos),
			"Cosh":	ValueOf(math.Cosh),
			"Dim":	ValueOf(math.Dim),
			"E":	ValueOf(math.E),
			"Erf":	ValueOf(math.Erf),
			"Erfc":	ValueOf(math.Erfc),
			"Erfcinv":	ValueOf(math.Erfcinv),
			"Erfinv":	ValueOf(math.Erfinv),
			"Exp":	ValueOf(math.Exp),
			"Exp2":	ValueOf(math.Exp2),
			"Expm1":	ValueOf(math.Expm1),
			"FMA":	ValueOf(math.FMA),
			"Float32bits":	ValueOf(math.Float32bits),
			"Float32frombits":	ValueOf(math.Float32frombits),
			"Float64bits":	ValueOf(math.Float64bits),
			"Float64frombits":	ValueOf(math.Float64frombits),
			"Floor":	ValueOf(math.Floor),
			"Frexp":	ValueOf(math.Frexp),
			"Gamma":	ValueOf(math.Gamma),
			"Hypot":	ValueOf(math.Hypot),
			"Ilogb":	ValueOf(math.Ilogb),
			"Inf":	ValueOf(math.Inf),
			"IsInf":	ValueOf(math.IsInf),
			"IsNaN":	ValueOf(math.IsNaN),
			"J0":	ValueOf(math.J0),
			"J1":	ValueOf(math.J1),
			"Jn":	ValueOf(math.Jn),
			"Ldexp":	ValueOf(math.Ldexp),
			"Lgamma":	ValueOf(math.Lgamma),
			"Ln10":	ValueOf(math.Ln10),
			"Ln2":	ValueOf(math.Ln2),
			"Log":	ValueOf(math.Log),
			"Log10":	ValueOf(math.Log10),
			"Log10E":	ValueOf(math.Log10E),
			"Log1p":	ValueOf(math.Log1p),
			"Log2":	ValueOf(math.Log2),
			"Log2E":	ValueOf(math.Log2E),
			"Logb":	ValueOf(math.Logb),
			"Max":	ValueOf(math.Max),
			"MaxFloat32":	ValueOf(math.MaxFloat32),
			"MaxFloat64":	ValueOf(math.MaxFloat64),
			"MaxInt":	ValueOf(int64(math.MaxInt)),
			"MaxInt16":	ValueOf(math.MaxInt16),
			"MaxInt32":	ValueOf(math.MaxInt32),
			"MaxInt64":	ValueOf(int64(math.MaxInt64)),
			"MaxInt8":	ValueOf(math.MaxInt8),
			"MaxUint":	ValueOf(uint64(math.MaxUint)),
			"MaxUint16":	ValueOf(math.MaxUint16),
			"MaxUint32":	ValueOf(uint32(math.MaxUint32)),
			"MaxUint64":	ValueOf(uint64(math.MaxUint64)),
			"MaxUint8":	ValueOf(math.MaxUint8),
			"Min":	ValueOf(math.Min),
			"MinInt":	ValueOf(int64(math.MinInt)),
			"MinInt16":	ValueOf(math.MinInt16),
			"MinInt32":	ValueOf(math.MinInt32),
			"MinInt64":	ValueOf(int64(math.MinInt64)),
			"MinInt8":	ValueOf(math.MinInt8),
			"Mod":	ValueOf(math.Mod),
			"Modf":	ValueOf(math.Modf),
			"NaN":	ValueOf(math.NaN),
			"Nextafter":	ValueOf(math.Nextafter),
			"Nextafter32":	ValueOf(math.Nextafter32),
			"Phi":	ValueOf(math.Phi),
			"Pi":	ValueOf(math.Pi),
			"Pow":	ValueOf(math.Pow),
			"Pow10":	ValueOf(math.Pow10),
			"Remainder":	ValueOf(math.Remainder),
			"Round":	ValueOf(math.Round),
			"RoundToEven":	ValueOf(math.RoundToEven),
			"Signbit":	ValueOf(math.Signbit),
			"Sin":	ValueOf(math.Sin),
			"Sincos":	ValueOf(math.Sincos),
			"Sinh":	ValueOf(math.Sinh),
			"SmallestNonzeroFloat32":	ValueOf(math.SmallestNonzeroFloat32),
			"SmallestNonzeroFloat64":	ValueOf(math.SmallestNonzeroFloat64),
			"Sqrt":	ValueOf(math.Sqrt),
			"Sqrt2":	ValueOf(math.Sqrt2),
			"SqrtE":	ValueOf(math.SqrtE),
			"SqrtPhi":	ValueOf(math.SqrtPhi),
			"SqrtPi":	ValueOf(math.SqrtPi),
			"Tan":	ValueOf(math.Tan),
			"Tanh":	ValueOf(math.Tanh),
			"Trunc":	ValueOf(math.Trunc),
			"Y0":	ValueOf(math.Y0),
			"Y1":	ValueOf(math.Y1),
			"Yn":	ValueOf(math.Yn),
		}, Untypeds: map[string]string{
			"E":	"float:271828182845904523536028747135266249775724709369995957496696763/100000000000000000000000000000000000000000000000000000000000000",
			"Ln10":	"float:23025850929940456840179914546843642076011014886287729760333279/10000000000000000000000000000000000000000000000000000000000000",
			"Ln2":	"float:693147180559945309417232121458176568075500134360255254120680009/1000000000000000000000000000000000000000000000000000000000000000",
			"Log10E":	"float:10000000000000000000000000000000000000000000000000000000000000/23025850929940456840179914546843642076011014886287729760333279",
			"Log2E":	"float:1000000000000000000000000000000000000000000000000000000000000000/693147180559945309417232121458176568075500134360255254120680009",
			"MaxFloat32":	"float:340282346638528859811704183484516925440",
			"MaxFloat64":	"float:179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368",
			"MaxInt":	"int:9223372036854775807",
			"MaxInt16":	"int:32767",
			"MaxInt32":	"int:2147483647",
			"MaxInt64":	"int:9223372036854775807",
			"MaxInt8":	"int:127",
			"MaxUint":	"int:18446744073709551615",
			"MaxUint16":	"int:65535",
			"MaxUint32":	"int:4294967295",
			"MaxUint64":	"int:18446744073709551615",
			"MaxUint8":	"int:255",
			"MinInt":	"int:-9223372036854775808",
			"MinInt16":	"int:-32768",
			"MinInt32":	"int:-2147483648",
			"MinInt64":	"int:-9223372036854775808",
			"MinInt8":	"int:-128",
			"Phi":	"float:80901699437494742410229341718281905886015458990288143106772431/50000000000000000000000000000000000000000000000000000000000000",
			"Pi":	"float:314159265358979323846264338327950288419716939937510582097494459/100000000000000000000000000000000000000000000000000000000000000",
			"SmallestNonzeroFloat32":	"float:1/713623846352979940529142984724747568191373312",
			"SmallestNonzeroFloat64":	"float:1/202402253307310618352495346718917307049556649764142118356901358027430339567995346891960383701437124495187077864316811911389808737385793476867013399940738509921517424276566361364466907742093216341239767678472745068562007483424692698618103355649159556340810056512358769552333414615230502532186327508646006263307707741093494784",
			"Sqrt2":	"float:70710678118654752440084436210484903928483593768847403658833987/50000000000000000000000000000000000000000000000000000000000000",
			"SqrtE":	"float:164872127070012814684865078781416357165377610071014801157507931/100000000000000000000000000000000000000000000000000000000000000",
			"SqrtPhi":	"float:63600982475703448212621123086874574585780402092004812430832019/50000000000000000000000000000000000000000000000000000000000000",
			"SqrtPi":	"float:177245385090551602729816748334114518279754945612238712821380779/100000000000000000000000000000000000000000000000000000000000000",
		}, 
	}
}
