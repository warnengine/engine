// All calls to OpenGL happens here !!!

package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

func gfxGenVertexArray() uint32 {
	var vao uint32
	vao = 0
	gl.GenVertexArrays(1, &vao)
	return vao
}

func gfxGenBuffer(vertexArray uint32, length int, size int, data []float32, attribIndex uint32) {
	// Bind future buffer to this vertex array
	gl.BindVertexArray(vertexArray)

	// Generate the actual buffer
	var buffer uint32
	buffer = 0
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		length*4,
		gl.Ptr(data),
		gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(attribIndex)
	gl.VertexAttribPointer(uint32(attribIndex), int32(size), gl.FLOAT, false, 0, nil)
}

func gfxDrawArrays(vertexArray uint32, length int) {
	gl.BindVertexArray(vertexArray)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(length))
}
