package main

import (
	"github.com/go-gl/gl/v4.2-core/gl"
)

// Mesh represent a vertex buffer and has a specific length
type Mesh struct {
	Vao    uint32
	Length int32
}

// CreateMesh loads a .obj file
func CreateMesh(obj string) Mesh {
	// Load model
	vertices, uvs, normals := LoadModel(obj)

	// Vertex Buffer Object
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	

	// Vertex Array Buffer -> vertices
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	

	// Vertex Array Buffer -> uvs
	var uvsBuffer uint32
	gl.GenBuffers(1, &uvsBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, uvsBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(uvs)*4, gl.Ptr(uvs), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, uvsBuffer)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)
	

	// Vertex Array Buffer -> normals
	var normalsBuffer uint32
	gl.GenBuffers(1, &normalsBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, normalsBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(normals)*4, gl.Ptr(normals), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(2)
	gl.BindBuffer(gl.ARRAY_BUFFER, normalsBuffer)
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)
	

	return Mesh{vao, int32(len(vertices))}
}

// Draw draws the triangles composing the mesh
func (mesh *Mesh) Draw() {
	// Bind our vertices
	gl.BindVertexArray(mesh.Vao)
	// Finally draw
	gl.DrawArrays(gl.TRIANGLES, 0, mesh.Length)
	
}
