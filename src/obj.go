package main

import (
	"bytes"

	"github.com/sheenobu/go-obj/obj"
)

// LoadModel loads a .obj from given path and return uvs/normals/vertices
func LoadModel(path string) ([]float32, []float32, []float32) {
	// Read content of the file
	cube, err := obj.NewReader(bytes.NewReader(ReadFile(path))).Read()
	if err != nil {
		panic(err)
	}

	var vertices, uvs, normals []float32

	// Iterate through face and vertex and feed vertices/uvs/normals
	for _, f := range cube.Faces {
		for _, p := range f.Points {
			vx := p.Vertex

			nx := float32(0.0)
			ny := float32(0.0)
			nz := float32(0.0)
			if p.Normal != nil {
				nx = float32(p.Normal.X)
				ny = float32(p.Normal.Y)
				nz = float32(p.Normal.Z)
			}

			u := float32(0.0)
			v := float32(0.0)
			if p.Texture != nil {
				u = float32(p.Texture.U)
				v = float32(p.Texture.V)
			}

			// Feed vertices/uvs/normals
			vertices = append(vertices, float32(vx.X), float32(vx.Y), float32(vx.Z))
			uvs = append(uvs, u, v)
			normals = append(normals, nx, ny, nz)
		}
	}

	return vertices, uvs, normals
}
