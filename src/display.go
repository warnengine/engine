package main

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Display manages the window of the game
type Display struct {
	window *glfw.Window
}

func createDisplay(screen Screen) Display {
	// Init glfw
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	// Some window hint
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)

	// Create window
	window, err := glfw.CreateWindow(screen.Width, screen.Height, "Warnengine", nil, nil)
	if err != nil {
		panic(err)
	}

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
	gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS)
	// gl.Enable(gl.MULTISAMPLE)

	gl.DebugMessageCallback(gl.DebugProc(ogldebugcb), gl.Ptr(nil))

	gl.DebugMessageInsert(
		gl.DEBUG_SOURCE_APPLICATION,
		gl.DEBUG_TYPE_ERROR,
		1, // Id
		gl.DEBUG_SEVERITY_NOTIFICATION,
		-1, // Length (negative => null-terminated)
		gl.Str("hello world\x00"))

	return Display{
		window: window,
	}
}

// StopDisplay kills glfw and destroys the window
func (display *Display) StopDisplay() {
	display.window.Destroy()
	glfw.Terminate()
}
