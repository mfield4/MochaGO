// Package game is meant to encapsulate the data and behavior used to to manage the game, it's window, and the renderer
package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go_entity_component_services/pkg/components"
	"go_entity_component_services/pkg/rendering"
	"go_entity_component_services/pkg/systems"
)

// The game struct holds all of the game datastructures.
// Upon initialization, the game struct is responsible for initializing glfw and opengl

type Game struct {
	window *glfw.Window

	entities *[]components.Entity

	positions *[]components.Position
	momentums *[]components.Momentum
	camera    *[]*components.Camera

	// gameWorld *protos.GameWorld

	// Systems??
	inputs    []systems.InputSystem
	asyncs    []systems.AsyncSystem
	renderers []systems.RenderSystem
}

func NewGame(width, height int) *Game {
	// init opengl and stuffs
	g := &Game{
		window: rendering.InitGlfw(width, height),
		// Data
		positions: &[]components.Position{},
		momentums: &[]components.Momentum{},
		camera:    &[]*components.Camera{},
		// Systems
		inputs:    []systems.InputSystem{},
		asyncs:    []systems.AsyncSystem{},
		renderers: []systems.RenderSystem{},
	}
	rendering.InitOpenGL()
	// inputs
	glfwInput := systems.NewGlfwInput(g.window)

	chunkManager := systems.NewChunkManager()

	cameraSys := systems.NewCameraSystem(g.camera)
	// rendering stuff
	skyboxRenderer := systems.NewSkyboxRenderer()

	glfwInput.AddMouseMotionListeners(cameraSys.MouseMotionCommand)
	glfwInput.AddKeyListeners(cameraSys.KeyCommand)

	g.addInputSystems(glfwInput, chunkManager)
	g.addAsyncSystems(chunkManager)
	g.addRenderSystems(skyboxRenderer)
	return g
}

func (g *Game) Destroy() {
	g.window.Destroy()
	glfw.Terminate()
}

func (g *Game) IsRunning() bool {
	return !g.window.ShouldClose()
}

// Start async systems
func (g *Game) Start() {
	for _, sys := range g.asyncs {
		go systems.AsyncStart(sys)
	}
}

// Update whatever needs updating on main thread
func (g *Game) Update() {
	for _, sys := range g.inputs {
		sys.Poll()
	}
}

// Draw all objects on scene.
func (g *Game) Draw() {
	// TODO set view and projection matrix at the correct times via the ubo
	// create ubo component and connect it to various entities

	gl.ClearColor(0.1, 0.1, 0.1, 0.1)
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)

	// SET THE VIEW MATRIX TO CURRENTLY ACTIVE CAMERA
	cam := (*g.camera)[0]

	view := cam.GetViewMatrix()

	cam.VPUBO.SetMat4(view, 1)

	for _, sys := range g.renderers {
		// TODO DO TIME STUFF MANG
		sys.Render(0)
	}

	g.window.SwapBuffers()
}

func (g *Game) addInputSystems(inputs ...systems.InputSystem) {
	for _, sys := range inputs {
		g.inputs = append(g.inputs, sys)
	}
}

func (g *Game) addAsyncSystems(asyncs ...systems.AsyncSystem) {
	for _, sys := range asyncs {
		g.asyncs = append(g.asyncs, sys)
	}
}

func (g *Game) addRenderSystems(renderers ...systems.RenderSystem) {
	for _, sys := range renderers {
		g.renderers = append(g.renderers, sys)
	}
}
