package main

import (
	"github.com/go-gl/gl/v4.2-core/gl"
)

type Pipeline struct {
	quadMat            uint32
	quadVao            uint32
	frameBufferDiffuse uint32
	frameBufferShadows uint32
	DiffuseTexture     uint32
	ShadowTexture      uint32
}

func CreatePipeline() Pipeline {

	frameBufferDiffuse, diffuseTexture := prepareDiffuse()
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

	// Vertex Buffer Object
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(quadVertices)*4, gl.Ptr(quadVertices), gl.STATIC_DRAW)

	// Vertex Array Buffer -> vertices
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	// Vertex Array Buffer -> uvs
	var uvsBuffer uint32
	gl.GenBuffers(1, &uvsBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, uvsBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(quadUvs)*4, gl.Ptr(quadUvs), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, uvsBuffer)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	quadMat := CreateMaterial("Shaders/sprite.vs.glsl", "Shaders/sprite.fs.glsl")

	return Pipeline{quadMat, vao, frameBufferDiffuse, frameBufferShadows, diffuseTexture, depthTexture}

}

func prepareDiffuse() (uint32, uint32) {
	var framebuffer uint32
	gl.GenFramebuffers(1, &framebuffer)
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)

	var renderedTexture uint32
	gl.GenTextures(1, &renderedTexture)
	gl.BindTexture(gl.TEXTURE_2D, renderedTexture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, 1280, 720, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	var depthrenderbuffer uint32
	gl.GenRenderbuffers(1, &depthrenderbuffer)
	gl.BindRenderbuffer(gl.RENDERBUFFER, depthrenderbuffer)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH_COMPONENT, 1280, 720)
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
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, depthTexture, 0)
	gl.DrawBuffer(gl.COLOR_ATTACHMENT0)

	return framebuffer, depthTexture
}

// BeginDiffuse bind the corresponding target texture and prepare it for rendering
func (pipeline *Pipeline) BeginDiffuse() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, pipeline.frameBufferDiffuse)
	gl.Viewport(0, 0, 1280, 720)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (pipeline *Pipeline) EndDiffuse() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Viewport(0, 0, 1280, 720)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	UseMaterial(pipeline.quadMat)
	gl.ActiveTexture(gl.TEXTURE0)
	UseTexture(pipeline.DiffuseTexture)
	gl.BindVertexArray(pipeline.quadVao)
	gl.DrawArrays(gl.TRIANGLES, 0, 9)
}

func (pipeline *Pipeline) BeginShadow() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, pipeline.frameBufferShadows)
	gl.Viewport(0, 0, 2048, 2048)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (pipeline *Pipeline) EndShadow() {
}
