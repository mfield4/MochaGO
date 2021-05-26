package systems

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"go_entity_component_services/protos"
)

type GlfwInput struct {
	mouseListeners []func(command protos.MouseCommand)
	keyListeners   []func(command protos.KeyCommand)
}

// This type creates a set of callbacks to communicate between glfw and the rest of the program

func NewGlfwInput(window *glfw.Window) *GlfwInput {
	// set up methods of GlfwInput struct to handle glfw inputs
	g := &GlfwInput{}

	window.SetCursorPosCallback(g.cursorPositionCallback)
	window.SetMouseButtonCallback(g.mouseButtonCallback)
	window.SetKeyCallback(g.keyCallback)

	return g
}

func (g *GlfwInput) Poll() {
	glfw.PollEvents()
}

func (g *GlfwInput) AddMouseListeners(funcs ...func(command protos.MouseCommand)) {
	for _, f := range funcs {
		g.mouseListeners = append(g.mouseListeners, f)
	}
}

func (g *GlfwInput) AddKeyListeners(funcs ...func(command protos.KeyCommand)) {
	for _, f := range funcs {
		g.keyListeners = append(g.keyListeners, f)
	}
}

func (g *GlfwInput) cursorPositionCallback(window *glfw.Window, xpos, ypos float64) {
	// construct mouse position proto
}

func (g *GlfwInput) mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	// construct mouse position proto
}

func (g *GlfwInput) keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mode glfw.ModifierKey) {
	// construct mouse position window
}
