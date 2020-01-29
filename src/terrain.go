package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Terrain struct {
	vao       uint32
	program   Program
	texture   Texture
	transform Transform
}

// CreateTerrain creates terrain from a texture with a given side size.
// Terrain will be at center position.
func CreateTerrain(textureFile string, size float32) Terrain {
	// Position, rotation, scale
	transform := CreateTransform(mgl32.Vec3{0.0, -1.0, 0.0}, mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{1.0, 1.0, 1.0})
	// Load the texture
	texture, _, _ := CreateTexture(textureFile)
	// Load shader
	mat := CreateMaterial("Shaders/terrain.vs.glsl", "Shaders/terrain.fs.glsl")
	// Load a simple plane
	vertices := []float32{
		-size, -size, // upper left
		+size, -size, // upper right
		-size, +size, // down left

		+size, -size, // upper left
		+size, +size, // down right
		-size, +size} // down left

	uvs := []float32{
		0, 1,
		1, 1,
		0, 0,

		1, 1,
		0, 1,
		0, 0}

	// Vertex Buffer Object
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Vertex Array Buffer -> vertices
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	// Vertex Array Buffer -> uvs
	var uvsBuffer uint32
	gl.GenBuffers(1, &uvsBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, uvsBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(uvs)*4, gl.Ptr(uvs), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, uvsBuffer)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	return Terrain{vao, mat, texture, transform}
}

// Prepare prepares the terrain to be drawn.
// Bind the shader and the texture. Must bind camera after that
func (terrain *Terrain) Prepare() {
	// Bind the shader
	terrain.program.Use()
	// Bind the transform matrices
	terrain.program.UseInputMatrix(terrain.transform.Model, "model")
	terrain.transform.GetModelView()
	// Enable the texture and link it
	gl.ActiveTexture(gl.TEXTURE0)
	UseTexture(terrain.texture)
}

// Draw calls the rendering api to draw the terrain.
// Bind the mesh and draw it.
func (terrain *Terrain) Draw() {
	// Bind our vertices
	gl.BindVertexArray(terrain.vao)
	// Finally draw
	gl.DrawArrays(gl.TRIANGLES, 0, 9)
}
