package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go_entity_component_services/pkg/components"
	"go_entity_component_services/pkg/systems"
	"log"
	"strings"
)

// The game struct holds all of the game datastructures.
// Upon initialization, the game struct is responsible for initializing glfw and opengl

const (
	vertexShaderSource = `

`
	fragmentShaderSource = `

`
)

type Game struct {
	window *glfw.Window

	positions *[]components.Position
	momentums *[]components.Momentum
	camera    *[]components.Camera

	// gameWorld *protos.GameWorld

	// Systems??
	inputs    []systems.InputSystem
	asyncs    []systems.AsyncSystem
	renderers []systems.RenderSystem
}

func NewGame(width, height int) *Game {
	// init opengl and stuffs
	g := &Game{
		window: initGlfw(width, height),
		// Data
		positions: &[]components.Position{},
		momentums: &[]components.Momentum{},
		camera:    &[]components.Camera{},
		// Systems
		inputs:    []systems.InputSystem{},
		asyncs:    []systems.AsyncSystem{},
		renderers: []systems.RenderSystem{},
	}
	initOpenGL()
	// inputs
	glfwInput := systems.NewGlfwInput(g.window)

	chunkManager := systems.NewChunkManager()

	cameraSys := systems.NewCameraSystem(g.camera)
	// rendering stuff
	chunkRenderer := systems.NewChunkRenderer()
	skyboxRenderer := systems.NewSkyboxRenderer()

	glfwInput.AddMouseMotionListeners(cameraSys.MouseMotionCommand)
	glfwInput.AddKeyListeners(cameraSys.KeyCommand)

	g.addInputSystems(glfwInput, chunkManager)
	g.addAsyncSystems(chunkManager)
	g.addRenderSystems(chunkRenderer, skyboxRenderer)
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

	for _, sys := range g.renderers {
		// TODO DO TIME STUFF MANG
		sys.Render(0)
	}
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

// initGlfw initializes glfw and returns a Window to use.
func initGlfw(width, height int) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// TODO CONFIGURE WINDOW SETTINGS
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
