package components

import "github.com/go-gl/mathgl/mgl32"

type Position struct {
	mgl32.Vec3
}

type Momentum struct {
	Mass float32
	Vel  mgl32.Vec3
}
