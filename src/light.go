package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Light is a weird orthogonal camera used to draw shadow map
type Light struct {
	projection mgl32.Mat4
	view       mgl32.Mat4
	position   mgl32.Vec3
}

// CreateLight creates a light according to given position
func CreateLight(position mgl32.Vec3) Light {
	depthProjectionMatrix := mgl32.Ortho(-10, 10, -10, 10, -10, 20)
	depthViewMatrix := mgl32.LookAt(position.X(), position.Y(), position.Z(), 0, 0, 0, 0, 1, 0)

	return Light{depthProjectionMatrix, depthViewMatrix, position}
}

// UseLight binds the current light to the shader
func (program *Program) UseLight(light Light, shadowPass bool) {
	if shadowPass {
		program.UseInputMatrix(light.projection, "projection")
		program.UseInputMatrix(light.view, "view")
	} else {
		program.UseInputMatrix(light.projection, "lightProjection")
		program.UseInputMatrix(light.view, "lightView")
	}
}
