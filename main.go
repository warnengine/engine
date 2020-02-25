package main

import (
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread()
}

var dejaVuSans Font
var input Input

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
	log.Println(gl.GoStr(gl.GetString(gl.VERSION)))
	log.Println(gl.GoStr(gl.GetString(gl.RENDERER)))

	// pipeline := CreatePipeline(screen)

	// HERE STARTS THE GOOD STUFF
	dejaVuSans = CreateFont("Fonts/DejaVuSans.png", "Fonts/DejaVuSans.fnt", 64, screen)

	// Basic form for testing purpose
	form := CreateForm(display.window, dejaVuSans, screen)
	form.AddButton(CreateDefaultButton("CLICK ME", mgl32.Vec2{1, -1}))

	// camera := CreateCamera(mgl32.Vec3{0, 3, 0}, display.window, int32(screen.Width), int32(screen.Height))

	scene1 := CreateScene(TerrainDefinition{"Textures/terrainHeight.jpg", "Textures/terrainDiffuse.jpg", 10}, screen, &display)
	scene1.Register(ModelDefinition{"Meshes/monkey.obj", "Textures/abstract.jpg", "Shaders/basic", false, true})
	scene1.Register(ModelDefinition{"Meshes/sphere.obj", "Textures/map.png", "Shaders/basic", false, true})
	// scene1.Register(ModelDefinition{"Meshes/terrain.obj", "Textures/terrainDiffuse.jpg", "Shaders/basic", true, false})
	scene1.Load()
	scene1.Activate(&display)

	// To compute frame per second
	now := time.Now().UnixNano()

	// light := CreateLight(mgl32.Vec3{0.5, 2.0, 2.0})

	bias := mgl32.Mat4{
		0.5, 0.0, 0.0, 0.0,
		0.0, 0.5, 0.0, 0.0,
		0.0, 0.0, 0.5, 0.0,
		0.5, 0.5, 0.5, 1.0}
	_ = bias

	gl.ClearColor(0.51, 0.51, 0.8, 1.0)

	input = CreateInput(display)

	// terrain := CreateTerrain("Textures/terrainHeight.jpg", "Textures/terrainDiffuse.jpg", 10)
	// _ = terrain

	for !display.window.ShouldClose() && !input.IsKeyDown(ESC) {
		// Picking stuff
		if display.window.GetMouseButton(glfw.MouseButtonLeft) == glfw.Press {
			scene1.models[0].transform.SetPosition(input.GetRayPosition(scene1.camera, 3.0))
		}
		// Rendering stuff
		timePerFrame := float64((time.Now().UnixNano() - now)) / 1e+9 // seconde
		_ = timePerFrame
		now = time.Now().UnixNano()

		scene1.Draw()

		/*pipeline.BeginDiffuse()
		// Draw stuff
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		camera = UpdateCamera(camera, display.window)
		// ========================
		// DIFFUSE
		// ========================
		// Prepare for model rendering
		monkey.Prepare()
		// Bind shadow mapping
		gl.ActiveTexture(gl.TEXTURE1)
		UseTexture(pipeline.ShadowTexture)
		// Bind our light
		monkey.program.UseLight(light, false)
		// Bind our camera
		monkey.program.UseCamera(camera)
		// Draw our cube mesh
		monkey.Draw()
		// SPHERE
		// Prepare for model rendering
		sphere.Prepare()
		// Bind shadow mapping
		gl.ActiveTexture(gl.TEXTURE1)
		UseTexture(pipeline.ShadowTexture)
		// Bind our light
		sphere.program.UseLight(light, false)
		// Bind our camera
		sphere.program.UseCamera(camera)
		// Draw our cube mesh
		sphere.Draw()
		// TERRAIN
		terrain.Prepare()
		terrain.program.UseCamera(camera)
		gl.ActiveTexture(gl.TEXTURE1)
		UseTexture(pipeline.ShadowTexture)
		terrain.program.UseInputInt(1, "castShadow")
		terrain.program.UseInputMatrix(bias, "bias")
		terrain.program.UseLight(light, false)
		terrain.Draw()
		// A bit of text rendering
		dejaVuSans.Draw(fmt.Sprintf("%d", int32(math.Ceil(1/timePerFrame)))+"fps", Color{0.2, 0.8, 0.2}, mgl32.Vec2{0.0, 0.0})
		form.Draw()
		// End for color rendering
		pipeline.EndDiffuse()
		// ========================
		// SHADOWS
		// ========================
		/*pipeline.BeginShadow()
		pipeline.shadowMat.UseLight(light, true)
		pipeline.shadowMat.UseInputMatrix(monkey.transform.Model, "model")
		monkey.Draw()
		pipeline.shadowMat.UseInputMatrix(sphere.transform.Model, "model")
		sphere.Draw()
		pipeline.EndShadow()*/

		// monkey.transform.SetRotation(monkey.transform.Rotation.Add(mgl32.Vec3{0.01, 0.01, 0.01}))

		form.Draw()

		display.window.SwapBuffers()
		glfw.PollEvents()
		// Statistics stuff
		frame++

	}
	display.StopDisplay()
}
