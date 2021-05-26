package systems

import "go_entity_component_services/protos"

// Handles basic camera movement and updating the camera with information

type CameraSystem struct {
}

func NewCameraSystem() *CameraSystem {
	return &CameraSystem{}
}

func (c *CameraSystem) Start() {
	panic("implement me")
}

func (c *CameraSystem) Step(dt float32) {
	panic("implement me")
}

func (c *CameraSystem) Stop() {
	panic("implement me")
}

func (c *CameraSystem) MouseCommand(mouse protos.MouseCommand) {

}

func (c *CameraSystem) KeyCommand(key protos.KeyCommand) {}
