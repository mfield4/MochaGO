package systems

// TODO??
type InputSystem interface {
	Poll()
}

type AsyncSystem interface {
	Step(dt float32) bool
	Stop()
}

func AsyncStart(sys AsyncSystem) {
	// TODO ADD TIMING THING
	for sys.Step(0) {
	}
}

type RenderSystem interface {
	Render(dt float32)
}
