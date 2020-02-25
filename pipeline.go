package main

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Pipeline regroups diffuse and shadow map and handle the step to render a frame
type Pipeline struct {
	quadMat            Program
	shadowMat          Program
	quadVao            uint32
	frameBufferDiffuse uint32
	frameBufferShadows uint32
	DiffuseTexture     uint32
	ShadowTexture      uint32
	screen             Screen
}

// CreatePipeline creates a pipeline (creates and binds opengl textures)
func CreatePipeline(screen Screen) Pipeline {
	log.Printf("Loading Pipeline:")
	frameBufferDiffuse, diffuseTexture := prepareDiffuse(screen)
	frameBufferShadows, depthTexture := prepareShadows()
	// RECT TO FIT THE SCREEN
	quadVertices := []float32{
		0.0, 0.0,
		0.0, -2.0,
		2.0, 0.0,

		0.0, -2.0,
		2.0, -2.0,
		2.0, 0.0}

	quadUvs := []float32{
		0.0, 0.0,
		0.0, -1.0,
		1.0, 0.0,

		0.0, -1.0,
		1.0, -1.0,
		1.0, 0.0}

	vao := gfxGenVertexArray()

	gfxGenBuffer(vao, len(quadVertices), 2, quadVertices, 0)
	gfxGenBuffer(vao, len(quadUvs), 2, quadUvs, 1)

	quadMat := CreateProgram("Shaders/sprite.vs.glsl", "Shaders/sprite.fs.glsl")
	shadowMat := CreateProgram("Shaders/shadows.vs.glsl", "Shaders/shadows.fs.glsl")

	return Pipeline{quadMat, shadowMat, vao, frameBufferDiffuse, frameBufferShadows, diffuseTexture, depthTexture, screen}

}

func prepareDiffuse(screen Screen) (uint32, uint32) {
	var framebuffer uint32
	gl.GenFramebuffers(1, &framebuffer)
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)

	var renderedTexture uint32
	gl.GenTextures(1, &renderedTexture)
	gl.BindTexture(gl.TEXTURE_2D, renderedTexture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, int32(screen.Width), int32(screen.Height), 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	var depthrenderbuffer uint32
	gl.GenRenderbuffers(1, &depthrenderbuffer)
	gl.BindRenderbuffer(gl.RENDERBUFFER, depthrenderbuffer)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH_COMPONENT, int32(screen.Width), int32(screen.Height))
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.RENDERBUFFER, depthrenderbuffer)

	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, renderedTexture, 0)
	gl.DrawBuffer(gl.COLOR_ATTACHMENT0)

	return framebuffer, renderedTexture
}

func prepareShadows() (uint32, uint32) {
	var framebuffer uint32
	gl.GenFramebuffers(1, &framebuffer)
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)

	var depthTexture uint32
	gl.GenTextures(1, &depthTexture)
	gl.BindTexture(gl.TEXTURE_2D, depthTexture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH_COMPONENT16, 2048, 2048, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	// gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, depthTexture, 0)
	gl.DrawBuffer(gl.NONE)

	return framebuffer, depthTexture
}

// BeginDiffuse binds the corresponding target texture and prepare it for rendering
func (pipeline *Pipeline) BeginDiffuse() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, pipeline.frameBufferDiffuse)
	gl.Viewport(0, 0, int32(pipeline.screen.Width), int32(pipeline.screen.Height))

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.CullFace(gl.BACK)
}

// EndDiffuse draws the diffuse map on screen
func (pipeline *Pipeline) EndDiffuse() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Viewport(0, 0, int32(pipeline.screen.Width), int32(pipeline.screen.Height))
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	pipeline.quadMat.Use()
	gl.ActiveTexture(gl.TEXTURE0)
	UseTexture(pipeline.DiffuseTexture)
	gl.BindVertexArray(pipeline.quadVao)
	gl.DrawArrays(gl.TRIANGLES, 0, 9)

}

// BeginShadow binds the shadow map for depth rendering
func (pipeline *Pipeline) BeginShadow() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, pipeline.frameBufferShadows)
	gl.Viewport(0, 0, 2048, 2048)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.CullFace(gl.FRONT)

	pipeline.shadowMat.Use()
}

// EndShadow does nothing but it's cool to have an end before a begin isn'it ?
func (pipeline *Pipeline) EndShadow() {
}
