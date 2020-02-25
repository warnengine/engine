package main

import (
	"log"
	"time"
)

// Mesh represent a vertex buffer and has a specific length
type Mesh struct {
	Vao    uint32
	Length int
}

// CreateMesh loads a .obj file
func CreateMesh(obj string) Mesh {
	log.Printf("-> Loading %s", obj)
	start := time.Now()
	// Load model
	vertices, uvs, normals := LoadModel(obj)

	vao := gfxGenVertexArray()

	gfxGenBuffer(vao, len(vertices), 3, vertices, 0)

	gfxGenBuffer(vao, len(uvs), 2, uvs, 1)

	gfxGenBuffer(vao, len(normals), 3, normals, 2)

	elapsed := time.Now().Sub(start)
	log.Printf("-> End of loading %f", elapsed.Seconds())

	return Mesh{vao, len(vertices)}
}

// Draw draws the triangles composing the mesh
func (mesh *Mesh) Draw() {
	gfxDrawArrays(mesh.Vao, mesh.Length)
}
