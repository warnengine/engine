package main

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.2-core/gl"

	"github.com/go-gl/mathgl/mgl32"
)

// Scene is a group of models
type Scene struct {
	loaded   bool
	pipeline Pipeline
	light    Light
	camera   Camera
	defs     []ModelDefinition
	models   []Model
}

// CreateScene creates an empty scene.
func CreateScene(screen Screen, display *Display) Scene {
	return Scene{
		false,
		CreatePipeline(screen),
		CreateLight(mgl32.Vec3{0.5, 2.0, 2.0}),
		CreateCamera(mgl32.Vec3{0, 3, 0}, display.window, int32(screen.Width), int32(screen.Height)),
		[]ModelDefinition{},
		[]Model{}}
}

// Activate makes the scene the current scene
func (scene *Scene) Activate(display *Display) {
	// display.window.SetUserPointer(unsafe.Pointer(&scene.camera))
}

// Register adds a new model to the scene but don't load it.
func (scene *Scene) Register(def ModelDefinition) {
	scene.defs = append(scene.defs, def)
}

// Load loads all models of the scene.
func (scene *Scene) Load() {
	// Link the camera
	scene.models = make([]Model, len(scene.defs))
	for i, model := range scene.defs {
		log.Println(fmt.Sprintf("Loading Model:"))
		scene.models[i] = CreateModel(model.mesh, model.texture, model.program, model.castShadow, model.genShadow)
	}
	scene.loaded = true
}

// Draw calls the rendering api to draw all instance of model
func (scene *Scene) Draw() {
	// First update our little camera
	scene.camera.Update()
	// Update our input
	input.Update()
	bias := mgl32.Mat4{
		0.5, 0.0, 0.0, 0.0,
		0.0, 0.5, 0.0, 0.0,
		0.0, 0.0, 0.5, 0.0,
		0.5, 0.5, 0.5, 1.0}
	if !scene.loaded {
		log.Println("Scene wasn't actually loaded...")
		log.Println("Loading it during a drawing frame :-/")
		scene.Load()
	}
	scene.pipeline.BeginDiffuse()
	for _, model := range scene.models {
		// Prepare for model rendering
		model.Prepare()
		// Bind shadow mapping
		gl.ActiveTexture(gl.TEXTURE1)
		UseTexture(scene.pipeline.ShadowTexture)
		// Bind our light
		model.program.UseLight(scene.light, false)
		// Bind our camera
		model.program.UseCamera(scene.camera)
		// Yep play with shadow
		if model.castShadow {
			model.program.UseInputInt(1, "castShadow")
		} else {
			model.program.UseInputInt(0, "castShadow")
		}
		// Some bias
		model.program.UseInputMatrix(bias, "bias")
		// Finally draw the model
		model.Draw()
	}
	scene.pipeline.EndDiffuse()

	scene.pipeline.BeginShadow()
	for _, model := range scene.models {
		if model.genShadow {
			scene.pipeline.shadowMat.UseLight(scene.light, true)
			scene.pipeline.shadowMat.UseInputMatrix(model.transform.Model, "model")
			model.Draw()
		}
	}
	scene.pipeline.EndShadow()
}
