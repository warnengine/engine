package main

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Terrain is a plane modified by a texture
type Terrain struct {
	mesh      Mesh
	program   Program
	heightMap Texture
	diffuse   Texture
	transform Transform
}

// CreateTerrain creates terrain from a texture with a given side size.
// Terrain will be at center position.
func CreateTerrain(heightMapFile string, diffuseFile string, size float32) Terrain {
	// Position, rotation, scale
	transform := CreateTransform(mgl32.Vec3{0.0, -1.0, 0.0}, mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{1.0, 1.0, 1.0})
	// Load the texture
	heightMap, _, _ := CreateTexture(heightMapFile)
	diffuse, _, _ := CreateTexture(diffuseFile)
	// Load shader
	mat := CreateProgram("Shaders/terrain.vs.glsl", "Shaders/terrain.fs.glsl")

	// Simple plane
	plane := CreateMesh("Meshes/terrain.obj")

	return Terrain{plane, mat, heightMap, diffuse, transform}
}

// Prepare prepares the terrain to be drawn.
// Bind the shader and the texture. Must bind camera after that
func (terrain *Terrain) Prepare() {
	// Bind the shader
	terrain.program.Use()
	// Bind the transform matrices
	terrain.program.UseInputMatrix(terrain.transform.Model, "model")
	// Enable the texture and link it
	gl.ActiveTexture(gl.TEXTURE0)
	UseTexture(terrain.diffuse)
	// Enable the normal texture and link it
	gl.ActiveTexture(gl.TEXTURE2)
	UseTexture(terrain.heightMap)
}

// Draw calls the rendering api to draw the terrain.
// Bind the mesh and draw it.
func (terrain *Terrain) Draw() {
	// Bind our vertices
	terrain.mesh.Draw()
}
