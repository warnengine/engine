package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// Color is a combinaison of two primary color
type Color struct {
	red   float32
	green float32
	blue  float32
}

// Style isn't yet implemented
type Style struct {
	color Color
}

// Form contains a group of UI element
type Form struct {
	window  *glfw.Window
	buttons []Button
	texts   []Text
	font    Font
	screen  Screen
	visible bool
}

// CreateForm creates an empty form from a given window
// By default, set to visible
func CreateForm(window *glfw.Window, font Font, screen Screen) Form {
	return Form{window, []Button{}, []Text{}, font, screen, true}
}

// Draw calls rendering api to draw all elements of this form
func (form *Form) Draw() {
	if form.visible {
		// Reset cursor
		form.window.SetCursor(glfw.CreateStandardCursor(glfw.ArrowCursor))
		// Draw all texts
		for _, text := range form.texts {
			_ = text
		}
		// Draw all buttons
		// Get mousePosition first
		mouseX, mouseY := form.window.GetCursorPos()
		// Get real mousePosition
		mouseX = (1.0 / (float64(form.screen.Width) / 2)) * mouseX
		mouseY = (1.0 / (float64(form.screen.Height) / 2)) * mouseY * -1
		for _, button := range form.buttons {
			button.Draw(form, mgl32.Vec2{float32(mouseX), float32(mouseY)})
		}
	}
}

// AddButton appends a new button to draw
func (form *Form) AddButton(button Button) {
	// Compute the size of the button
	width := form.font.GetTextSize(button.content)
	size := mgl32.Vec2{width, 1.0 / float32(form.screen.Height) * float32(form.font.fontSize)}
	button.size = size
	form.buttons = append(form.buttons, button)
}

// AddText appends a new text to draw
func (form *Form) AddText(text Text) {
	form.texts = append(form.texts, text)
}

// Hide sets visible to false
func (form *Form) Hide() {
	form.visible = false
}

// Show sets visible to true
func (form *Form) Show() {
	form.visible = true
}

// Button is a text giving action when mouse give focus
type Button struct {
	content     string
	position    mgl32.Vec2
	fgColor     Color
	bgColor     Color
	size        mgl32.Vec2
	justPressed bool
}

// CreateDefaultButton creates default button with default color.
// Size not computed yet. Must add it to a form to do it.
func CreateDefaultButton(content string, position mgl32.Vec2) Button {
	return Button{content, position, Color{0.8, 0.8, 0.8}, Color{0.0, 0.0, 0.0}, mgl32.Vec2{32, 128}, false}
}

// Draw calls the rendering api to draw the button
func (button *Button) Draw(parent *Form, mousePosition mgl32.Vec2) {
	// Check if mouse is over the button and change his color
	text := button.content
	if mousePosition.X() > button.position.X() && mousePosition.Y() < button.position.Y() && mousePosition.X() < button.position.X()+button.size.X() && mousePosition.Y() > button.position.Y()-button.size.Y() {
		text = "YEAH !!!"
		parent.window.SetCursor(glfw.CreateStandardCursor(glfw.HandCursor))
	}
	// println(fmt.Sprintf("X: %f, Y: %f", mousePosition.X(), mousePosition.Y()))
	parent.font.Draw(text, button.fgColor, button.position)
}

// Text is a combinaison of letters to draw
type Text struct {
}
