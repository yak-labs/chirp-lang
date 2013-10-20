package img

// Example Usage:
//   set m,format,width,height [image-load /tmp/old.jpg]
//   set z [image-resample $m  320 240 ]
//   image-save /tmp/new.jpg $z jpg

// For more about the golang image library,
// see http://golang.org/doc/articles/image_package.html

import (
	"github.com/nfnt/resize"
	. "github.com/yak-labs/chirp-lang"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"os"
	R "reflect"
)

func cmdImageLoad(fr *Frame, argv []T) T {
	inFilename := Arg1(argv)

	r, err := os.Open(inFilename.String())
	if err != nil {
		panic(err)
	}

	img, format, err := image.Decode(r)
	if err != nil {
		panic(err)
	}

	b := img.Bounds()
	width := int64(b.Max.X - b.Min.X)
	height := int64(b.Max.Y - b.Min.Y)

	return MkList([]T{
		MkValue(R.ValueOf(img)),
		MkString(format),
		MkInt(width),
		MkInt(height),
	})
}

func cmdImageResample(fr *Frame, argv []T) T {
	imgT, widthT, heightT := Arg3(argv)

	img := imgT.Raw().(image.Image)
	width := uint(widthT.Int())
	height := uint(heightT.Int())

	z := resize.Resize(width, height, img, resize.Lanczos3)
	return MkValue(R.ValueOf(z))
}

func cmdImageSave(fr *Frame, argv []T) T {
	filenameT, imgT, formatT := Arg3(argv)

	w, err := os.Create(filenameT.String())
	if err != nil {
		panic(err)
	}
	img := imgT.Raw().(image.Image)
	switch formatT.String() {
	case "jpeg", "jpg":
		err = jpeg.Encode(w, img, nil)
	case "png":
		err = png.Encode(w, img)
	default:
		panic("Bad image format: " + formatT.String())
	}
	if err != nil {
		panic(err)
	}
	return Empty
}

func init() {
	if Unsafes == nil {
		Unsafes = make(map[string]Command, 333)
	}

	Unsafes["image-load"] = cmdImageLoad
	Unsafes["image-save"] = cmdImageSave
	Unsafes["image-resample"] = cmdImageResample
}
