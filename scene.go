package main

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"

	"github.com/go-gl/mathgl/mgl32"
)

// Scene is a group of nodes.
type Scene struct {
	loaded   bool
	pipeline Pipeline
	light    Light
	camera   Camera
	nodes    []interface{}
	defs     []ModelDefinition
	models   []Model
	terrain  Terrain
}

// CreateScene creates an empty scene from a given terrain definition.
func CreateScene(terrain TerrainDefinition, screen Screen, display *Display) Scene {
	return Scene{
		false,
		CreatePipeline(screen),
		CreateLight(mgl32.Vec3{0.5, 2.0, 2.0}),
		CreateCamera(mgl32.Vec3{0, 3, 0}, display.window, int32(screen.Width), int32(screen.Height)),
		[]interface{},
		[]ModelDefinition{},
		[]Model{},
		CreateTerrain(terrain.heightMapFile, terrain.diffuseFile, terrain.size)}
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
	// Play with
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
	// Draw the terrain linked to this scene
	scene.terrain.Prepare()
	scene.terrain.program.UseCamera(scene.camera)
	gl.ActiveTexture(gl.TEXTURE1)
	UseTexture(scene.pipeline.ShadowTexture)
	scene.terrain.program.UseInputInt(1, "castShadow")
	scene.terrain.program.UseInputMatrix(bias, "bias")
	scene.terrain.program.UseLight(scene.light, false)
	scene.terrain.Draw()

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
	dejaVuSans.Draw("Warnengine 0.1", Color{1.0, 1.0, 1.0}, mgl32.Vec2{0.0, 0.0})
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
