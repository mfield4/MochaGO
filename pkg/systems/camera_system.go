package systems

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"go_entity_component_services/pkg/components"
)

// Handles basic camera movement and updating the camera with information

type CameraSystem struct {
	cameras *[]components.Camera
	deltaT  float64
	lastT   float64
}

func NewCameraSystem(cams *[]components.Camera) *CameraSystem {
	*cams = append(*cams, components.Camera{
		Position:         mgl32.Vec3{},
		Front:            mgl32.Vec3{0.0, 0.0, -1.0},
		Up:               mgl32.Vec3{0.0, 1.0, 0.0},
		Right:            mgl32.Vec3{1.0, 0.0, 0.0},
		WorldUp:          mgl32.Vec3{0.0, 1.0, 0.0},
		Yaw:              0,
		Pitch:            0,
		MovementSpeed:    0.05,
		MouseSensitivity: 0.05,
		Zoom:             45,
	})
	return &CameraSystem{
		cameras: cams,
	}
}

func (c *CameraSystem) MouseMotionCommand(mouse components.CursorPositionCommand) {
	cam := (*c.cameras)[0]

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

func (c *CameraSystem) KeyCommand(key components.KeyCommand) {
	cam := (*c.cameras)[0]
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
}
