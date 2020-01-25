package main

import (
	"log"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Display struct {
	window *glfw.Window
}

func createDisplay(screen Screen) Display {
	// Init glfw
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	// Specifiy sizes of the screen
	glfw.WindowHint(glfw.Resizable, glfw.False)

	// Create window
	window, err := glfw.CreateWindow(screen.Width, screen.Height, "Warnengine", nil, nil)
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}

	// V-sync please
	glfw.SwapInterval(1)

	// Some OpenGL tweaks
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.CullFace(gl.BACK)
	gl.Disable(gl.CULL_FACE)

	return Display{
		window: window,
	}
}

// StopDisplay kills glfw and destroys the window
func (display *Display) StopDisplay() {
	display.window.Destroy()
	glfw.Terminate()
}
