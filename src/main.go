package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// Init our file system
	InitFileSystem()

	// Current rendered frame
	// Increment each new frame
	var frame int

	// Init glfw
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Specifiy sizes of the screen
	glfw.WindowHint(glfw.Resizable, glfw.False)

	// Create our window
	screen := Screen{1280, 720}
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

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.CullFace(gl.BACK)
	gl.Disable(gl.CULL_FACE)

	glfw.SwapInterval(1)

	log.Println("Hello from OpenGL")
	log.Println(gl.GoStr(gl.GetString(gl.VENDOR)))
	log.Println(gl.GoStr(gl.GetString(gl.RENDERER)))

	pipeline := CreatePipeline()

	// HERE STARTS THE GOOD STUFF
	font := CreateFont("Fonts/DejaVuSans.png", "Fonts/DejaVuSans.fnt", screen)
	_ = font

	camera := CreateCamera(mgl32.Vec3{0, 3, 0}, window, int32(screen.Width), int32(screen.Height))
	window.SetUserPointer(unsafe.Pointer(&camera))
	/*========================
	Shaders
	========================*/
	// Load and compile our shaders
	mat := CreateMaterial("Shaders/basic.vs.glsl", "Shaders/basic.fs.glsl")
	/*========================
	Cube
	========================*/
	// Init our 3D model "Cube"
	cube := CreateMesh("Meshes/monkey.obj")
	_ = cube
	// Load our texture
	cubeTexture, _, _ := CreateTexture("Textures/abstract.jpg")
	_ = cubeTexture
	// Prepare our transform that will describe position/rotation/scale of our object
	cubeTransform := CreateTransform(mgl32.Vec3{0.0, 1.5, 0.0}, mgl32.Vec3{1.0, 1.0, 1.0}, mgl32.Vec3{1.0, 1.0, 1.0})

	// Init our 3D model "Cube"
	theMap := CreateMesh("Meshes/map.obj")
	_ = cube
	// Load our texture
	theMapTexture, _, _ := CreateTexture("Textures/map.png")
	// Prepare our transform that will describe position/rotatigeometric intersection testingon/scale of our object
	theMapTransform := CreateTransform(mgl32.Vec3{0.0, 0, 0.0}, mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{1.0, 1.0, 1.0})
	// Link it to our shadergeometric intersection testing
	UseMaterial(mat)
	UseInputMatrix(mat, cubeTransform.Model, "model")

	now := time.Now().UnixNano()

	shadowMat := CreateMaterial("Shaders/shadows.vs.glsl", "Shaders/shadows.fs.glsl")
	_ = shadowMat

	light := CreateLight(mgl32.Vec3{0.5, 2, 2})

	bias := mgl32.Mat4{
		0.5, 0.0, 0.0, 0.0,
		0.0, 0.5, 0.0, 0.0,
		0.0, 0.0, 0.5, 0.0,
		0.5, 0.5, 0.5, 1.0}

	gl.ClearColor(0.51, 0.51, 0.8, 1.0)

	input := Input{window}

	for !window.ShouldClose() && !input.IsKeyDown(ESC) {
		// Picking stuff
		if window.GetMouseButton(glfw.MouseButtonLeft) == glfw.Press {
			cubeTransform.SetPosition(input.GetRayPosition(camera, 0.0))
		}
		// Rendering stuff
		timePerFrame := float64((time.Now().UnixNano() - now)) / 1e+9 // seconde
		_ = timePerFrame
		now = time.Now().UnixNano()

		pipeline.BeginDiffuse()
		// Draw stuff
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		camera = UpdateCamera(camera, window)
		UseMaterial(mat)
		UseCamera(camera, mat)
		// ========================
		// DIFFUSE
		// ========================
		// Bind shader
		UseMaterial(mat)
		// CUBE
		// Bind texture
		gl.ActiveTexture(gl.TEXTURE0)
		UseTexture(cubeTexture)
		// Bind shadow mapping
		gl.ActiveTexture(gl.TEXTURE1)
		UseTexture(pipeline.ShadowTexture)
		// Bind model view
		UseInputMatrix(mat, cubeTransform.Model, "model")
		UseInputMatrix(mat, bias, "bias")
		UseInputInt(mat, 0, "castShadow")
		// Bind our light
		light.Use(mat, false)
		// Draw our cube mesh
		cube.Draw()
		// MAP
		// Bind texture
		gl.ActiveTexture(gl.TEXTURE0)
		UseTexture(theMapTexture)
		// Bind shadow mapping
		gl.ActiveTexture(gl.TEXTURE1)
		UseTexture(pipeline.ShadowTexture)
		// Bind model view
		UseInputMatrix(mat, theMapTransform.Model, "model")
		UseInputMatrix(mat, bias, "bias")
		UseInputInt(mat, 1, "castShadow")
		// Bind our light
		light.Use(mat, false)
		// Draw our cube mesh
		theMap.Draw()
		font.Draw(fmt.Sprintf("%d", int32(math.Ceil(1/timePerFrame)))+" fps", mgl32.Vec2{0.0, 0.0})
		pipeline.EndDiffuse()
		// ========================
		// SHADOWS
		// ========================
		pipeline.BeginShadow()
		UseMaterial(shadowMat)
		light.Use(shadowMat, true)
		UseInputMatrix(shadowMat, cubeTransform.Model, "model")
		cube.Draw()
		/*UseInputMatrix(shadowMat, theMapTransform.Model, "model")
		DrawMesh(theMap)*/
		pipeline.EndShadow()

		cubeTransform.SetRotation(cubeTransform.Rotation.Add(mgl32.Vec3{0.01, 0.01, 0.01}))

		window.SwapBuffers()
		glfw.PollEvents()
		// Statistics stuff
		frame++

	}

}
