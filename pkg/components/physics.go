package components

import "github.com/go-gl/mathgl/mgl32"

type Position struct {
	mgl32.Vec3
}

func (p *Position) CTypeID() COMPONENT_MASK {
	return POSITIONS
}

type Momentum struct {
	Mass float32
	Vel  mgl32.Vec3
}

func (m *Momentum) CTypeID() COMPONENT_MASK {
	return MOMENTUM
}
