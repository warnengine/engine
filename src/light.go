package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Light struct {
	projection mgl32.Mat4
	view       mgl32.Mat4
	position   mgl32.Vec3
}

func CreateLight(position mgl32.Vec3) Light {
	depthProjectionMatrix := mgl32.Ortho(-10, 10, -10, 10, -10, 20)
	depthViewMatrix := mgl32.LookAt(position.X(), position.Y(), position.Z(), 0, 0, 0, 0, 1, 0)

	return Light{depthProjectionMatrix, depthViewMatrix, position}
}

func (light *Light) Use(material Material, shadowPass bool) {
	if shadowPass {
		UseInputMatrix(material, light.projection, "projection")
		UseInputMatrix(material, light.view, "view")
	} else {
		UseInputMatrix(material, light.projection, "lightProjection")
		UseInputMatrix(material, light.view, "lightView")
	}
}
