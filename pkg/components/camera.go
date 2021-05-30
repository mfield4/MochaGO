package components

import (
	"github.com/go-gl/mathgl/mgl32"
	"go_entity_component_services/pkg/rendering"
)

// Camera is the camera object maintaing the stae
type Camera struct {
	Position mgl32.Vec3
	Front    mgl32.Vec3
	Up       mgl32.Vec3
	Right    mgl32.Vec3
	WorldUp  mgl32.Vec3

	// Eular Angles
	Yaw   float64
	Pitch float64

	// Camera options
	MovementSpeed    float64
	MouseSensitivity float64
	Zoom             float32
	VPUBO            rendering.Ubo
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

// GetViewMatrix returns the view natrix
func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	eye := c.Position
	center := c.Position.Add(c.Front)
	up := c.Up

	return mgl32.LookAt(
		eye.X(), eye.Y(), eye.Z(),
		center.X(), center.Y(), center.Z(),
		up.X(), up.Y(), up.Z())
}
