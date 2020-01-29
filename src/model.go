package main

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Model struct {
	mesh      Mesh
	texture   Texture
	program   Program
	transform Transform
}

// CreateModel inits the three parts of a model.
// Mesh, Texture and Program.
func CreateModel(meshPath string, texturePath string, programPath string) Model {
	mesh := CreateMesh(meshPath)
	texture, _, _ := CreateTexture(texturePath)
	program := CreateProgram(programPath+".vs.glsl", programPath+".fs.glsl")
	transform := CreateTransform(mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{1.0, 1.0, 1.0})

	return Model{mesh, texture, program, transform}
}

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

func (model *Model) Draw() {
	model.mesh.Draw()
}
