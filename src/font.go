package main

import (
	"encoding/json"
	"strconv"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Screen contains characteristics of the current window.
type Screen struct {
	Width  int
	Height int
}

// Font contains informations for drawing all letters.
type Font struct {
	letters  map[string]Char
	texture  Texture
	material Material
}

// Char links to a vertex buffer and several font configurations.
type Char struct {
	vao     uint32
	advance float32
	top     float32
	uvs     []float32
}

// CreateFont loads from settings files all textures and informations to draw letters/
func CreateFont(textureFile string, configFile string, screen Screen) Font {
	// Load texture
	texture, textSizeX, textSizeY := CreateTexture(textureFile)
	// Load shader
	fontMat := CreateMaterial("Shaders/font.vs.glsl", "Shaders/font.fs.glsl")
	// Where we will store vao's letters
	letters := make(map[string]Char)
	// Read json
	jsonData := ReadFile(configFile)
	// Prepare json
	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})["chars"].(map[string]interface{})["char"].([]interface{})
	// Loop over char data
	for _, element := range data {
		current := element.(map[string]interface{})
		// Get width of the char
		width, err := strconv.Atoi(current["@width"].(string))
		if err != nil {
			panic(err)
		}
		height, err := strconv.Atoi(current["@height"].(string))
		if err != nil {
			panic(err)
		}
		advance, err := strconv.Atoi(current["@xadvance"].(string))
		if err != nil {
			panic(err)
		}
		top, err := strconv.Atoi(current["@yoffset"].(string))
		_ = top
		if err != nil {
			panic(err)
		}
		uvX, err := strconv.Atoi(current["@x"].(string))
		if err != nil {
			panic(err)
		}
		uvY, err := strconv.Atoi(current["@y"].(string))
		if err != nil {
			panic(err)
		}
		yOffset, err := strconv.Atoi(current["@yoffset"].(string))
		_ = yOffset
		if err != nil {
			panic(err)
		}
		xOffset, err := strconv.Atoi(current["@xoffset"].(string))
		_ = xOffset
		if err != nil {
			panic(err)
		}
		// According to width/height of the screen
		// Compute OpenGL coordinates
		realWidth := (1 / float32(screen.Width)) * float32(width)
		realHeight := (1 / float32(screen.Height)) * float32(height)
		vertices := []float32{
			0.0, 0.0,
			0.0, -realHeight,
			realWidth, 0.0,

			0.0, -realHeight,
			realWidth, -realHeight,
			realWidth, 0.0}

		realUvX := (1 / float32(textSizeX)) * float32(uvX)
		realUvY := (1 / float32(textSizeY)) * float32(textSizeY-uvY)

		realWidth = (1 / float32(textSizeX)) * float32(width)
		realHeight = (1 / float32(textSizeY)) * float32(height)

		realAdvance := (1 / float32(screen.Width)) * float32(advance)
		realTop := (1 / float32(screen.Width)) * float32(64-height)

		uvs := []float32{
			realUvX, realUvY,
			realUvX, realUvY - realHeight,
			realUvX + realWidth, realUvY,

			realUvX, realUvY - realHeight,
			realUvX + realWidth, realUvY - realHeight,
			realUvX + realWidth, realUvY}
		_ = uvs

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
		gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

		// Vertex Array Buffer -> uvs
		var uvsBuffer uint32
		gl.GenBuffers(1, &uvsBuffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, uvsBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, len(uvs)*4, gl.Ptr(uvs), gl.STATIC_DRAW)
		gl.EnableVertexAttribArray(1)
		gl.BindBuffer(gl.ARRAY_BUFFER, uvsBuffer)
		gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)
		asciiCode, err := strconv.Atoi(current["@id"].(string))
		if err != nil {
			panic(err)
		}
		letters[string(asciiCode)] = Char{vao, realAdvance, realTop, uvs}
	}
	return Font{letters, texture, fontMat}
}

// Draw draws a string from a given Font.
func (font *Font) Draw (content string, position mgl32.Vec2) {
	for i, char := range content {
		// Update position for next letter
		if i != 0 {
			position = position.Add(mgl32.Vec2{font.letters[string(content[i-1])].advance, 0.0})
		}
		UseMaterial(font.material)
		gl.ActiveTexture(gl.TEXTURE0)
		UseTexture(font.texture)
		UseInputVec2(font.material, position.Sub(mgl32.Vec2{0.0, font.letters[string(char)].top}), "i_position")
		gl.BindVertexArray(font.letters[string(char)].vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 9)
	}
}
