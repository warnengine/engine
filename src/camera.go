package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// Camera represents the eye of the user
type Camera struct {
	projection mgl32.Mat4
	view       mgl32.Mat4
	Position   mgl32.Vec3
	Width      int32
	Height     int32
	Distance   float32
}

// CreateCamera returns a camera according to given informations.
func CreateCamera(position mgl32.Vec3, window *glfw.Window, width int32, height int32) Camera {
	// window.SetScrollCallback(scrollCameraCb)
	projection := mgl32.Perspective(mgl32.DegToRad(60.0), float32(width)/float32(height), 0.1, 1000.0)
	view := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}.Add(position), mgl32.Vec3{0, 0, 0}.Add(position), mgl32.Vec3{0, 1, 0})
	return Camera{projection, view, position, width, height, 0.0}
}

// Update moves the camera according to the user input.
func (camera *Camera) Update( /*window *glfw.Window*/ ) {
	// Just update position
	delta := mgl32.Vec3{0.0, 0.0, 0.0}
	if input.IsKeyDown(UP) {
		delta = mgl32.Vec3{-0.1, 0.0, -0.1}
	} else if input.IsKeyDown(DOWN) {
		delta = mgl32.Vec3{0.1, 0.0, 0.1}
	}
	if input.IsKeyDown(LEFT) {
		delta = delta.Add(mgl32.Vec3{-0.1, 0.0, 0.1})
	} else if input.IsKeyDown(RIGHT) {
		delta = delta.Add(mgl32.Vec3{0.1, 0.0, -0.1})
	}
	// newCamera := UpdateCamPosition(camera, camera.Position.Add(delta))
	position := camera.Position.Add(delta)
	// camera.view = mgl32.LookAtV(mgl32.Vec3{3 + camera.Distance, 3 + camera.Distance, 3 + camera.Distance}.Add(position), mgl32.Vec3{0, 0, 0}.Add(position), mgl32.Vec3{0, 1, 0})
	camera.Position = position
	// return newCamera
	// Now update 'zoom'
	distance := camera.Distance
	if distance < 31 && distance > -2 {
		if (distance+0.2*input.scroll.Y()) < 30 && (distance+0.2*input.scroll.Y()) > -1 {
			camera.Distance += 0.25 * input.scroll.Y()
		}
	}
	camera.view = mgl32.LookAtV(mgl32.Vec3{3 + distance, 3 + distance, 3 + distance}.Add(position), mgl32.Vec3{0, 0, 0}.Add(position), mgl32.Vec3{0, 1, 0})
}

// UpdateCamPosition moves the camera and computes its view according to target position.
func UpdateCamPosition(camera Camera, position mgl32.Vec3) Camera {
	view := mgl32.LookAtV(mgl32.Vec3{3 + float32(camera.Distance), 3 + float32(camera.Distance), 3 + float32(camera.Distance)}.Add(position), mgl32.Vec3{0, 0, 0}.Add(position), mgl32.Vec3{0, 1, 0})
	return Camera{camera.projection, view, position, camera.Width, camera.Height, camera.Distance}
}

// Callback when the user use the scroll button.
/*func scrollCameraCb(window *glfw.Window, xoffset float64, yoffset float64) {
	distance := (*Camera)(window.GetUserPointer()).Distance
	if distance < 31 && distance > -2 {
		if (distance+0.2*yoffset) < 30 && (distance+0.2*yoffset) > -1 {
			(*Camera)(window.GetUserPointer()).Distance += 0.25 * yoffset
		}
	}
	distance = (*Camera)(window.GetUserPointer()).Distance
	position := (*Camera)(window.GetUserPointer()).Position
	(*Camera)(window.GetUserPointer()).view = mgl32.LookAtV(mgl32.Vec3{3 + float32(distance), 3 + float32(distance), 3 + float32(distance)}.Add(position), mgl32.Vec3{0, 0, 0}.Add(position), mgl32.Vec3{0, 1, 0})
}
*/
