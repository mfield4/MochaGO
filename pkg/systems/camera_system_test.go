package systems

import (
	"go_entity_component_services/pkg/components"
	"reflect"
	"testing"
)

func TestCameraSystem_KeyCommand(t *testing.T) {
	type fields struct {
		cameras []components.Camera
		deltaT  float64
		lastT   float64
	}
	type args struct {
		key components.KeyCommand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// c := &CameraSystem{
			// 	cameras: tt.fields.cameras,
			// 	deltaT:  tt.fields.deltaT,
			// 	lastT:   tt.fields.lastT,
			// }
		})
	}
}

func TestCameraSystem_MouseMotionCommand(t *testing.T) {
	type fields struct {
		cameras []components.Camera
		deltaT  float64
		lastT   float64
	}
	type args struct {
		mouse components.CursorPositionCommand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// c := &CameraSystem{
			// 	cameras: tt.fields.cameras,
			// 	deltaT:  tt.fields.deltaT,
			// 	lastT:   tt.fields.lastT,
			// }
		})
	}
}

func TestNewCameraSystem(t *testing.T) {
	type args struct {
		cams []components.Camera
	}
	tests := []struct {
		name string
		args args
		want *CameraSystem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCameraSystem(tt.args.cams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCameraSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}
