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

	screen := Screen{1280, 720}
	display := createDisplay(screen)

	log.Println("Hello from OpenGL")
	log.Println(gl.GoStr(gl.GetString(gl.VENDOR)))
	log.Println(gl.GoStr(gl.GetString(gl.RENDERER)))

	pipeline := CreatePipeline(screen)

	// HERE STARTS THE GOOD STUFF
	font := CreateFont("Fonts/DejaVuSans.png", "Fonts/DejaVuSans.fnt", screen)
	_ = font

	camera := CreateCamera(mgl32.Vec3{0, 3, 0}, display.window, int32(screen.Width), int32(screen.Height))
	display.window.SetUserPointer(unsafe.Pointer(&camera))
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
	cubeTransform := CreateTransform(mgl32.Vec3{0.0, 1.0, 0.0}, mgl32.Vec3{1.0, 1.0, 1.0}, mgl32.Vec3{1.0, 1.0, 1.0})

	// Init our 3D model "Cube"
	theMap := CreateMesh("Meshes/map.obj")
	_ = cube
	// Load our texture
	theMapTexture, _, _ := CreateTexture("Textures/map.png")
	// Prepare our transform that will describe position/rotatigeometric intersection testingon/scale of our object
	theMapTransform := CreateTransform(mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{1.0, 1.0, 1.0})
	// Link it to our shadergeometric intersection testing
	mat.Use()
	mat.UseInputMatrix(cubeTransform.Model, "model")

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

	input := Input{display.window}

	for !display.window.ShouldClose() && !input.IsKeyDown(ESC) {
		// Picking stuff
		if display.window.GetMouseButton(glfw.MouseButtonLeft) == glfw.Press {
			cubeTransform.SetPosition(input.GetRayPosition(camera, 1.0))
		}
		// Rendering stuff
		timePerFrame := float64((time.Now().UnixNano() - now)) / 1e+9 // seconde
		_ = timePerFrame
		now = time.Now().UnixNano()

		pipeline.BeginDiffuse()
		// Draw stuff
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		camera = UpdateCamera(camera, display.window)
		mat.Use()
		mat.UseCamera(camera)
		// ========================
		// DIFFUSE
		// ========================
		// Bind shader
		mat.Use()
		// CUBE
		// Bind texture
		gl.ActiveTexture(gl.TEXTURE0)
		UseTexture(cubeTexture)
		// Bind shadow mapping
		gl.ActiveTexture(gl.TEXTURE1)
		UseTexture(pipeline.ShadowTexture)
		// Bind model view
		mat.UseInputMatrix(cubeTransform.Model, "model")
		mat.UseInputMatrix(bias, "bias")
		mat.UseInputInt(0, "castShadow")
		// Bind our light
		mat.UseLight(light, false)
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
		mat.UseInputMatrix(theMapTransform.Model, "model")
		mat.UseInputMatrix(bias, "bias")
		mat.UseInputInt(1, "castShadow")
		// Bind our light
		mat.UseLight(light, false)
		// Draw our cube mesh
		theMap.Draw()
		font.Draw(fmt.Sprintf("%d", int32(math.Ceil(1/timePerFrame)))+" fps", mgl32.Vec2{0.0, 0.0})
		pipeline.EndDiffuse()
		// ========================
		// SHADOWS
		// ========================
		pipeline.BeginShadow()
		shadowMat.Use()
		shadowMat.UseLight(light, true)
		shadowMat.UseInputMatrix(cubeTransform.Model, "model")
		cube.Draw()
		/*UseInputMatrix(shadowMat, theMapTransform.Model, "model")
		DrawMesh(theMap)*/
		pipeline.EndShadow()

		cubeTransform.SetRotation(cubeTransform.Rotation.Add(mgl32.Vec3{0.01, 0.01, 0.01}))

		display.window.SwapBuffers()
		glfw.PollEvents()
		// Statistics stuff
		frame++

	}
	display.StopDisplay()
}
