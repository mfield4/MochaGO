package components

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"go_entity_component_services/pkg/rendering"
	"math"
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

func (c *Camera) UpdateVectors() {
	c.Front[0] = float32(math.Cos(mgl64.DegToRad(c.Yaw)) * math.Cos(mgl64.DegToRad(c.Pitch)))
	c.Front[1] = float32(math.Sin(mgl64.DegToRad(c.Pitch)))
	c.Front[2] = float32(math.Sin(mgl64.DegToRad(c.Yaw)) * math.Cos(mgl64.DegToRad(c.Pitch)))
	c.Front = c.Front.Normalize()
	// Also re-calculate the Right and Up vector
	// // Normalize the vectors, because their length gets closer to 0 the more you look up or down which results in slower movement.
	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}
