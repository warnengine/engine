package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Transform assembles the scale, rotation and position of a 3D object.
type Transform struct {
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Scale    mgl32.Vec3
	Model    mgl32.Mat4
}

// CreateTransform returns a Transform created from given informations.
func CreateTransform(position mgl32.Vec3, rotation mgl32.Vec3, scale mgl32.Vec3) Transform {
	transform := Transform{position, rotation, scale, mgl32.Ident4()}
	transform.Model = transform.GetModelView()
	return transform
}

// GetModelView computes scale, rotation and position and returns the matrix representation of the model.
func (transform *Transform) GetModelView() mgl32.Mat4 {
	var posMat, scaleMat mgl32.Mat4
	var rotMatX, rotMatY, rotMatZ, rotMat mgl32.Mat3
	posMat = mgl32.Translate3D(transform.Position.X(), transform.Position.Y(), transform.Position.Z())
	scaleMat = mgl32.Scale3D(transform.Scale.X(), transform.Scale.Y(), transform.Scale.Z())
	rotMatX = mgl32.Rotate3DX(transform.Rotation.X())
	rotMatY = mgl32.Rotate3DY(transform.Rotation.Y())
	rotMatZ = mgl32.Rotate3DZ(transform.Rotation.Z())
	rotMat = rotMatX.Mul3(rotMatY).Mul3(rotMatZ)

	model := posMat.Mul4(rotMat.Mat4()).Mul4(scaleMat)

	return model
}

// SetRotation sets the rotation. No ? Sure ?! What ???!!!
func (transform *Transform) SetRotation(rotation mgl32.Vec3) {
	transform.Rotation = rotation
	transform.Model = transform.GetModelView()
}

// SetPosition sets the position. Hell, that's useful.
func (transform *Transform) SetPosition(position mgl32.Vec3) {
	transform.Position = position
	transform.Model = transform.GetModelView()
}

// SetScale sets the scale and gives you a cookie with a warm coffee.
func (transform *Transform) SetScale(scale mgl32.Vec3) {
	transform.Scale = scale
	transform.Model = transform.GetModelView()
}
