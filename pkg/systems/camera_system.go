package systems

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"go_entity_component_services/pkg/components"
	"log"
	"math"
)

// Handles basic camera movement and updating the camera with information

type CameraSystem struct {
	cameras *[]components.Camera
	deltaT  float64
	lastT   float64
}

func NewCameraSystem(cams *[]components.Camera) *CameraSystem {
	*cams = append(*cams, components.NewCamera())
	return &CameraSystem{
		cameras: cams,
	}
}

func (c *CameraSystem) MouseMotionCommand(mouse components.CursorPositionCommand) {
	cam := &(*c.cameras)[0]

	// update camera vectors
	const sensitivity = 0.05

	xOffset := mouse.Xpos * sensitivity
	yOffset := mouse.Ypos * sensitivity
	cam.Yaw += xOffset
	cam.Pitch += yOffset

	if cam.Pitch > 89.0 {
		cam.Pitch = 89.0
	}
	if cam.Pitch < -89.0 {
		cam.Pitch = -89.0
	}

	log.Printf("Print the camera my man : %v\n", cam)
	c.updateCameraVectors()
}

func (c *CameraSystem) KeyCommand(key components.KeyCommand) {
	cam := &(*c.cameras)[0]
	// Update DeltaT
	currT := glfw.GetTime()
	velocity := float32(cam.MovementSpeed * (currT - c.lastT))
	c.deltaT = currT
	// Edit mouse position based on key command
	switch key.Key {
	case glfw.KeyW:
		cam.Position = cam.Position.Add(cam.Front.Mul(velocity))
	case glfw.KeyA:
		cam.Position = cam.Position.Sub(cam.Front.Mul(velocity))
	case glfw.KeyS:
		cam.Position = cam.Position.Sub(cam.Right.Mul(velocity))
	case glfw.KeyD:
		cam.Position = cam.Position.Add(cam.Right.Mul(velocity))
	}

	// c.updateCameraVectors()
}

func (c CameraSystem) updateCameraVectors() {
	cam := &(*c.cameras)[0]
	x := float32(math.Cos(mgl64.DegToRad(cam.Yaw)) * math.Cos(mgl64.DegToRad(cam.Pitch)))
	y := float32(math.Sin(mgl64.DegToRad(cam.Pitch)))
	z := float32(math.Sin(mgl64.DegToRad(cam.Yaw)) * math.Cos(mgl64.DegToRad(cam.Pitch)))
	front := mgl32.Vec3{x, y, z}
	front = front.Normalize()
	// Also re-calculate the Right and Up vector
	// Normalize the vectors, because their length gets closer to 0 the more you look up or down which results in slower movement.
	cam.Right = front.Cross(cam.WorldUp).Normalize()
	cam.Up = cam.Right.Cross(cam.Front).Normalize()
}
