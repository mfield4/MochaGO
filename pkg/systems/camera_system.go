package systems

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"go_entity_component_services/pkg/data"
)

// Handles basic camera movement and updating the camera with information

type CameraSystem struct {
	cameras []data.Camera
	deltaT  float64
}

func NewCameraSystem(cams []data.Camera) *CameraSystem {
	return &CameraSystem{
		cameras: cams,
	}
}

func (c *CameraSystem) MouseMotionCommand(mouse data.CursorPositionCommand) {
	cam := &c.cameras[0]

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

	cam.UpdateCameraVectors()
}

func (c *CameraSystem) KeyCommand(key data.KeyCommand) {
	cam := &c.cameras[0]
	velocity := float32(cam.MovementSpeed * c.deltaT)
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
}
