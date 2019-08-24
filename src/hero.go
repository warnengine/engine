package main

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Hero is model composed by a texture, material, mesh and a transform. Will receive a script component in the future.
type Hero struct {
	texture   uint32
	material  Material
	mesh      Mesh
	transform Transform
}

// Draw draws the hero according his components
func (hero *Hero) Draw(bindDiffuse bool) {
	bias := mgl32.Mat4{
		0.5, 0.0, 0.0, 0.0,
		0.0, 0.5, 0.0, 0.0,
		0.0, 0.0, 0.5, 0.0,
		0.5, 0.5, 0.5, 1.0}
	// Somewhere after BeginDiffuse
	if bindDiffuse {
		// Use the corresponding material
		UseMaterial(hero.material)
		// Use the corresponding  diffuse texture
		gl.ActiveTexture(gl.TEXTURE0)
		UseTexture(hero.texture)
		// HERE BIND SHADOW MAP
		// BIND OUR LIGHT
		UseInputMatrix(hero.material, hero.transform.Model, "model")
		UseInputMatrix(hero.material, bias, "bias")
		UseInputInt(hero.material, 0, "castShadow")
	}
}
