// this file was generated by gomacro command: import _b "image"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	image "image"
	color "image/color"
)

// reflection: allow interpreted code to import "image"
func init() {
	Packages["image"] = Package{
		Name: "image",
		Binds: map[string]Value{
			"Black":	ValueOf(&image.Black).Elem(),
			"Decode":	ValueOf(image.Decode),
			"DecodeConfig":	ValueOf(image.DecodeConfig),
			"ErrFormat":	ValueOf(&image.ErrFormat).Elem(),
			"NewAlpha":	ValueOf(image.NewAlpha),
			"NewAlpha16":	ValueOf(image.NewAlpha16),
			"NewCMYK":	ValueOf(image.NewCMYK),
			"NewGray":	ValueOf(image.NewGray),
			"NewGray16":	ValueOf(image.NewGray16),
			"NewNRGBA":	ValueOf(image.NewNRGBA),
			"NewNRGBA64":	ValueOf(image.NewNRGBA64),
			"NewNYCbCrA":	ValueOf(image.NewNYCbCrA),
			"NewPaletted":	ValueOf(image.NewPaletted),
			"NewRGBA":	ValueOf(image.NewRGBA),
			"NewRGBA64":	ValueOf(image.NewRGBA64),
			"NewUniform":	ValueOf(image.NewUniform),
			"NewYCbCr":	ValueOf(image.NewYCbCr),
			"Opaque":	ValueOf(&image.Opaque).Elem(),
			"Pt":	ValueOf(image.Pt),
			"Rect":	ValueOf(image.Rect),
			"RegisterFormat":	ValueOf(image.RegisterFormat),
			"Transparent":	ValueOf(&image.Transparent).Elem(),
			"White":	ValueOf(&image.White).Elem(),
			"YCbCrSubsampleRatio410":	ValueOf(image.YCbCrSubsampleRatio410),
			"YCbCrSubsampleRatio411":	ValueOf(image.YCbCrSubsampleRatio411),
			"YCbCrSubsampleRatio420":	ValueOf(image.YCbCrSubsampleRatio420),
			"YCbCrSubsampleRatio422":	ValueOf(image.YCbCrSubsampleRatio422),
			"YCbCrSubsampleRatio440":	ValueOf(image.YCbCrSubsampleRatio440),
			"YCbCrSubsampleRatio444":	ValueOf(image.YCbCrSubsampleRatio444),
			"ZP":	ValueOf(&image.ZP).Elem(),
			"ZR":	ValueOf(&image.ZR).Elem(),
		}, Types: map[string]Type{
			"Alpha":	TypeOf((*image.Alpha)(nil)).Elem(),
			"Alpha16":	TypeOf((*image.Alpha16)(nil)).Elem(),
			"CMYK":	TypeOf((*image.CMYK)(nil)).Elem(),
			"Config":	TypeOf((*image.Config)(nil)).Elem(),
			"Gray":	TypeOf((*image.Gray)(nil)).Elem(),
			"Gray16":	TypeOf((*image.Gray16)(nil)).Elem(),
			"Image":	TypeOf((*image.Image)(nil)).Elem(),
			"NRGBA":	TypeOf((*image.NRGBA)(nil)).Elem(),
			"NRGBA64":	TypeOf((*image.NRGBA64)(nil)).Elem(),
			"NYCbCrA":	TypeOf((*image.NYCbCrA)(nil)).Elem(),
			"Paletted":	TypeOf((*image.Paletted)(nil)).Elem(),
			"PalettedImage":	TypeOf((*image.PalettedImage)(nil)).Elem(),
			"Point":	TypeOf((*image.Point)(nil)).Elem(),
			"RGBA":	TypeOf((*image.RGBA)(nil)).Elem(),
			"RGBA64":	TypeOf((*image.RGBA64)(nil)).Elem(),
			"RGBA64Image":	TypeOf((*image.RGBA64Image)(nil)).Elem(),
			"Rectangle":	TypeOf((*image.Rectangle)(nil)).Elem(),
			"Uniform":	TypeOf((*image.Uniform)(nil)).Elem(),
			"YCbCr":	TypeOf((*image.YCbCr)(nil)).Elem(),
			"YCbCrSubsampleRatio":	TypeOf((*image.YCbCrSubsampleRatio)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"Image":	TypeOf((*P_image_Image)(nil)).Elem(),
			"PalettedImage":	TypeOf((*P_image_PalettedImage)(nil)).Elem(),
			"RGBA64Image":	TypeOf((*P_image_RGBA64Image)(nil)).Elem(),
		}, Wrappers: map[string][]string{
			"NYCbCrA":	[]string{"Bounds","COffset","YCbCrAt","YOffset",},
		}, 
	}
}

// --------------- proxy for image.Image ---------------
type P_image_Image struct {
	Object	interface{}
	At_	func(_proxy_obj_ interface{}, x int, y int) color.Color
	Bounds_	func(interface{}) image.Rectangle
	ColorModel_	func(interface{}) color.Model
}
func (P *P_image_Image) At(x int, y int) color.Color {
	return P.At_(P.Object, x, y)
}
func (P *P_image_Image) Bounds() image.Rectangle {
	return P.Bounds_(P.Object)
}
func (P *P_image_Image) ColorModel() color.Model {
	return P.ColorModel_(P.Object)
}

// --------------- proxy for image.PalettedImage ---------------
type P_image_PalettedImage struct {
	Object	interface{}
	At_	func(_proxy_obj_ interface{}, x int, y int) color.Color
	Bounds_	func(interface{}) image.Rectangle
	ColorIndexAt_	func(_proxy_obj_ interface{}, x int, y int) uint8
	ColorModel_	func(interface{}) color.Model
}
func (P *P_image_PalettedImage) At(x int, y int) color.Color {
	return P.At_(P.Object, x, y)
}
func (P *P_image_PalettedImage) Bounds() image.Rectangle {
	return P.Bounds_(P.Object)
}
func (P *P_image_PalettedImage) ColorIndexAt(x int, y int) uint8 {
	return P.ColorIndexAt_(P.Object, x, y)
}
func (P *P_image_PalettedImage) ColorModel() color.Model {
	return P.ColorModel_(P.Object)
}

// --------------- proxy for image.RGBA64Image ---------------
type P_image_RGBA64Image struct {
	Object	interface{}
	At_	func(_proxy_obj_ interface{}, x int, y int) color.Color
	Bounds_	func(interface{}) image.Rectangle
	ColorModel_	func(interface{}) color.Model
	RGBA64At_	func(_proxy_obj_ interface{}, x int, y int) color.RGBA64
}
func (P *P_image_RGBA64Image) At(x int, y int) color.Color {
	return P.At_(P.Object, x, y)
}
func (P *P_image_RGBA64Image) Bounds() image.Rectangle {
	return P.Bounds_(P.Object)
}
func (P *P_image_RGBA64Image) ColorModel() color.Model {
	return P.ColorModel_(P.Object)
}
func (P *P_image_RGBA64Image) RGBA64At(x int, y int) color.RGBA64 {
	return P.RGBA64At_(P.Object, x, y)
}
