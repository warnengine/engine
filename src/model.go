package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Model is a group composed by a shader, a mesh and a texture.
// Model is drawable.
type Model struct {
	mesh       Mesh
	texture    Texture
	program    Program
	transform  Transform
	castShadow bool
	genShadow  bool
}

// ModelDefinition characterizes the components of a Model.
// When a model is registered, it's not immediately loaded.
type ModelDefinition struct {
	mesh       string
	texture    string
	program    string
	castShadow bool
	genShadow  bool
}

// CreateModel inits the three parts of a model.
// Mesh, Texture and Program.
func CreateModel(meshPath string, texturePath string, programPath string, castShadow bool, genShadow bool) Model {
	mesh := CreateMesh(meshPath)
	texture, _, _ := CreateTexture(texturePath)
	program := CreateProgram(programPath+".vs.glsl", programPath+".fs.glsl")
	transform := CreateTransform(mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{1.0, 1.0, 1.0})

	return Model{mesh, texture, program, transform, castShadow, genShadow}
}

// Prepare bind shaders, texture and matrices.
func (model *Model) Prepare() {
	// Bind the shader
	model.program.Use()
	// Bind the transform matrices
	model.program.UseInputMatrix(model.transform.Model, "model")
	// Use the diffuse texture
	gl.ActiveTexture(gl.TEXTURE0)
	UseTexture(model.texture)
	// There will be binding of camera and lights
}

// Draw calls the rendering api to draw the model
func (model *Model) Draw() {
	model.mesh.Draw()
}
