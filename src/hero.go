package main

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Colon is model composed by a texture, material, mesh and a transform. Will receive a script component in the future.
type Hero struct {
	texture   uint32
	material  Material
	mesh      Mesh
	transform Transform
}

// Draw draws the hero according his components
func (colon *Hero) Draw(diffusePass bool) {
	bias := mgl32.Mat4{
		0.5, 0.0, 0.0, 0.0,
		0.0, 0.5, 0.0, 0.0,
		0.0, 0.0, 0.5, 0.0,
		0.5, 0.5, 0.5, 1.0}
	// Somewhere after BeginDiffuse
	if diffusePass {
		// Use the corresponding material
		UseMaterial(colon.material)
		// Use the corresponding  diffuse texture
		gl.ActiveTexture(gl.TEXTURE0)
		UseTexture(colon.texture)
		// HERE BIND SHADOW MAP
		// BIND OUR LIGHT
		UseInputMatrix(colon.material, colon.transform.Model, "model")
		UseInputMatrix(colon.material, bias, "bias")
		UseInputInt(colon.material, 0, "castShadow")
	}
}
