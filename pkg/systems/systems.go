package systems

// TODO??
type InputSystem interface {
	Poll()
}

type AsyncSystem interface {
	Start()
	Step(dt float32)
	Stop()
}

type RenderSystem interface {
	Render(dt float32)
}
