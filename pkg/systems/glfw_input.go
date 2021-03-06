package systems

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"go_entity_component_services/pkg/components"
)

type GlfwInput struct {
	cursorPositionListeners []func(components.CursorPositionCommand)
	mouseButtonListeners    []func(components.MouseButtonCommand)
	keyListeners            []func(components.KeyCommand)
	cursorState             bool

	lastX, lastY float64
}

// This type creates a set of callbacks to communicate between glfw and the rest of the program

func NewGlfwInput(window *glfw.Window) *GlfwInput {
	// set up methods of GlfwInput struct to handle glfw inputs
	g := &GlfwInput{
		cursorState: true,
	}

	window.SetCursorPosCallback(g.cursorPositionCallback)
	window.SetMouseButtonCallback(g.mouseButtonCallback)
	window.SetKeyCallback(g.keyCallback)

	return g
}

func (g *GlfwInput) Poll() {
	glfw.PollEvents()
}

func (g *GlfwInput) AddMouseMotionListeners(funcs ...func(components.CursorPositionCommand)) {
	for _, f := range funcs {
		g.cursorPositionListeners = append(g.cursorPositionListeners, f)
	}
}

func (g *GlfwInput) AddMouseButtonListeners(funcs ...func(components.MouseButtonCommand)) {
	for _, f := range funcs {
		g.mouseButtonListeners = append(g.mouseButtonListeners, f)
	}
}

func (g *GlfwInput) AddKeyListeners(funcs ...func(components.KeyCommand)) {
	for _, f := range funcs {
		g.keyListeners = append(g.keyListeners, f)
	}
}

func (g *GlfwInput) cursorPositionCallback(window *glfw.Window, xpos, ypos float64) {
	// construct mouse position proto
	for _, lis := range g.cursorPositionListeners {
		go lis(components.CursorPositionCommand{Xpos: xpos - g.lastX, Ypos: g.lastY - ypos})
	}
	g.lastX = xpos
	g.lastY = ypos
}

func (g *GlfwInput) mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	// construct mouse position proto

	for _, pos := range g.mouseButtonListeners {
		go pos(components.MouseButtonCommand{})
	}
}

func (g *GlfwInput) keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mode glfw.ModifierKey) {
	// construct mouse position window

	if key == glfw.KeyEscape {
		window.SetShouldClose(true)
	}

	if key == glfw.KeyTab && action == glfw.Press {
		if g.cursorState {
			window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
			g.cursorState = false
		} else {
			window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
			g.cursorState = true
		}
	}

	for _, key := range g.keyListeners {
		go key(components.KeyCommand{})
	}
}
