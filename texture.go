package main

import (
	"bytes"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Texture represents the add of the stored image.
type Texture = uint32

// CreateTexture loads and stores an image and returns its address and dimension.
func CreateTexture(file string) (Texture, int, int) {
	start := time.Now()
	// Open file
	log.Printf("-> Loading %s", file)
	// Decode image
	img, _, err := image.Decode(bytes.NewReader(ReadFile(file)))
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	// Load texture with OpenGL
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	elapsed := time.Now().Sub(start)
	log.Printf("-> End of loading %f", elapsed.Seconds())

	return texture, rgba.Rect.Size().X, rgba.Rect.Size().Y
}

// UseTexture tells the rendering API to use a specific texture.
func UseTexture(texture Texture) {
	gl.BindTexture(gl.TEXTURE_2D, uint32(texture))
}
