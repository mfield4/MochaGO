package components

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
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
	Zoom             float64
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

func (c *Camera) UpdateCameraVectors() {
	x := float32(math.Cos(mgl64.DegToRad(c.Yaw)) * math.Cos(mgl64.DegToRad(c.Pitch)))
	y := float32(math.Sin(mgl64.DegToRad(c.Pitch)))
	z := float32(math.Sin(mgl64.DegToRad(c.Yaw)) * math.Cos(mgl64.DegToRad(c.Pitch)))
	front := mgl32.Vec3{x, y, z}
	front = front.Normalize()
	// Also re-calculate the Right and Up vector
	// Normalize the vectors, because their length gets closer to 0 the more you look up or down which results in slower movement.
	c.Right = front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}
