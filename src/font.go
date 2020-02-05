package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/go-gl/gl/v3.3-core/gl"
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
	program  Program
	fontSize int
}

// Char links to a vertex buffer and several font configurations.
type Char struct {
	vao     uint32
	advance float32
	uvs     []float32
}

// CreateFont loads from settings files all textures and informations to draw letters/
func CreateFont(textureFile string, configFile string, size int, screen Screen) Font {
	log.Printf("Loading Font:")
	// Load texture
	texture, textSizeX, textSizeY := CreateTexture(textureFile)
	// Load shader
	fontMat := CreateProgram("Shaders/font.vs.glsl", "Shaders/font.fs.glsl")
	// Where we will store vao's letters
	letters := make(map[string]Char)
	// Read json
	jsonData := ReadFile(configFile)
	// Prepare json
	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})["chars"].(map[string]interface{})["char"].([]interface{})
	// Get font size
	fontSize, err := strconv.Atoi(v.(map[string]interface{})["info"].(map[string]interface{})["@size"].(string))
	// Get desired scale
	scale := 1.0 / float32(fontSize) * float32(size)
	if err != nil {
		panic(err)
	}
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
		realWidth := (1 / float32(screen.Width)) * float32(width) * scale
		realHeight := (1 / float32(screen.Height)) * float32(height) * scale
		realYOffset := ((1 / float32(screen.Height)) * float32(yOffset)) * scale
		realXOffset := ((1 / float32(screen.Width)) * float32(xOffset)) / 2 * scale
		vertices := []float32{
			realXOffset, -realYOffset,
			realXOffset, -realHeight - realYOffset,
			realWidth + realXOffset, -realYOffset,

			realXOffset, -realHeight - realYOffset,
			realWidth + realXOffset, -realHeight - realYOffset,
			realWidth + realXOffset, -realYOffset}

		realUvX := (1 / float32(textSizeX)) * float32(uvX)
		realUvY := (1 / float32(textSizeY)) * float32(textSizeY-uvY)

		realWidth = (1 / float32(textSizeX)) * float32(width)
		realHeight = (1 / float32(textSizeY)) * float32(height)

		realAdvance := (1 / float32(screen.Width)) * float32(advance) * scale
		// realTop := (1 / float32(screen.Width)) * float32(fontSize-height)

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
		letters[string(asciiCode)] = Char{vao, realAdvance, uvs}
	}
	return Font{letters, texture, fontMat, fontSize}
}

// GetTextSize computes the width of a given text
func (font *Font) GetTextSize(content string) float32 {
	var size float32
	for i := range content {
		size += font.letters[string(content[i])].advance
	}
	return size
}

// Draw draws a string from a given Font.
func (font *Font) Draw(content string, color Color, position mgl32.Vec2) {
	for i, char := range content {
		// Update position for next letter
		if i != 0 {
			position = position.Add(mgl32.Vec2{font.letters[string(content[i-1])].advance, 0.0})
		}
		font.program.Use()
		font.program.UseInputVec3(mgl32.Vec3{color.red, color.green, color.blue}, "i_color")
		gl.ActiveTexture(gl.TEXTURE0)
		UseTexture(font.texture)

		font.program.UseInputVec2(position, "i_position")
		gl.BindVertexArray(font.letters[string(char)].vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 9)

	}
}
