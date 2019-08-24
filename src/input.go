package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// KeyCode is a keyboard key represented by an integer
type KeyCode uint8

// KeyCode is a keyboard key represented by an integer
const (
	A KeyCode = iota
	Z
	E
	R
	T
	H
	ESC
)

// Input allows high level input testing
type Input struct {
	window *glfw.Window
}

// GetRayPosition compute the position pointed by the mouse position on a infinite plane
func (input *Input) GetRayPosition(camera Camera, groundHeight float32) mgl32.Vec3 {
	mouseX, mouseY := input.window.GetCursorPos()
	x := (2.0*mouseX)/float64(camera.Width) - 1.0
	y := 1.0 - (2.0*mouseY)/float64(camera.Height)
	rayClip := mgl32.Vec4{float32(x), float32(y), -1.0, 1.0}

	rayEye := camera.projection.Inv().Mul4x1(rayClip)
	rayEye = mgl32.Vec4{rayEye.X(), rayEye.Y(), -1.0, 0.0}

	rayWor := camera.view.Inv().Mul4x1(rayEye)
	ray := mgl32.Vec3{rayWor.X(), rayWor.Y(), rayWor.Z()}.Normalize()
	ip := intersectPoint(ray, camera.Position.Add(mgl32.Vec3{3 + float32(camera.Distance), 3 + float32(camera.Distance), 3 + float32(camera.Distance)}), mgl32.Vec3{0, 1, 0}, mgl32.Vec3{0, 0, 0})

	return ip
}

// IsKeyDown check if user is currently pressing a key
func (input *Input) IsKeyDown(keyCode KeyCode) bool {
	switch keyCode {
	case A:
		return input.window.GetKey(glfw.KeyA) == glfw.Press
	case Z:
		return input.window.GetKey(glfw.KeyZ) == glfw.Press
	case E:
		return input.window.GetKey(glfw.KeyE) == glfw.Press
	case R:
		return input.window.GetKey(glfw.KeyR) == glfw.Press
	case T:
		return input.window.GetKey(glfw.KeyT) == glfw.Press
	case H:
		return input.window.GetKey(glfw.KeyH) == glfw.Press
	case ESC:
		return input.window.GetKey(glfw.KeyEscape) == glfw.Press
	default:
		return false
	}
}

func intersectPoint(rayVector mgl32.Vec3, rayPoint mgl32.Vec3, planeNormal mgl32.Vec3, planePoint mgl32.Vec3) mgl32.Vec3 {
	diff := rayPoint.Sub(planePoint)
	prod1 := diff.Dot(planeNormal)
	prod2 := rayVector.Dot(planeNormal)
	prod3 := prod1 / prod2
	return rayPoint.Sub(rayVector.Mul(prod3))
}
