package components

import (
	"github.com/go-gl/mathgl/mgl32"
	"go_entity_component_services/pkg/rendering"
)

type COMPONENT_MASK byte

const (
	CAMERA COMPONENT_MASK = 1 << iota
	POSITIONS
	MOMENTUM
)

type Entity uint32

func NewIDGenerator() func() Entity {
	var eid Entity
	return func() Entity {
		eid += 1
		return eid
	}
}

func NewCamera() *Camera {
	return &Camera{
		Position:         mgl32.Vec3{0.0, 0.0, 0.0},
		Front:            mgl32.Vec3{0.0, 0.0, -1.0},
		Up:               mgl32.Vec3{0.0, 1.0, 0.0},
		Right:            mgl32.Vec3{1.0, 0.0, 0.0},
		WorldUp:          mgl32.Vec3{0.0, 1.0, 0.0},
		Yaw:              0,
		Pitch:            0,
		MovementSpeed:    0.05,
		MouseSensitivity: 0.05,
		Zoom:             45,
		VPUBO: func() rendering.Ubo {
			vpUBO := rendering.NewViewProj(2)
			vpUBO.SetMat4(mgl32.Perspective(mgl32.DegToRad(45), float32(1920)/float32(1080), 0.1, 100), 0)
			return vpUBO
		}(),
	}
}
