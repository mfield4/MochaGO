package components

import (
	"github.com/go-gl/mathgl/mgl32"
	"reflect"
	"testing"
)

func TestCamera_GetViewMatrix(t *testing.T) {
	type fields struct {
		Position         mgl32.Vec3
		Front            mgl32.Vec3
		Up               mgl32.Vec3
		Right            mgl32.Vec3
		WorldUp          mgl32.Vec3
		Yaw              float64
		Pitch            float64
		MovementSpeed    float64
		MouseSensitivity float64
		Zoom             float64
	}
	tests := []struct {
		name   string
		fields fields
		want   mgl32.Mat4
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Camera{
				Position:         tt.fields.Position,
				Front:            tt.fields.Front,
				Up:               tt.fields.Up,
				Right:            tt.fields.Right,
				WorldUp:          tt.fields.WorldUp,
				Yaw:              tt.fields.Yaw,
				Pitch:            tt.fields.Pitch,
				MovementSpeed:    tt.fields.MovementSpeed,
				MouseSensitivity: tt.fields.MouseSensitivity,
				Zoom:             tt.fields.Zoom,
			}
			if got := c.GetViewMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetViewMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamera_UpdateCameraVectors(t *testing.T) {
	type fields struct {
		Position         mgl32.Vec3
		Front            mgl32.Vec3
		Up               mgl32.Vec3
		Right            mgl32.Vec3
		WorldUp          mgl32.Vec3
		Yaw              float64
		Pitch            float64
		MovementSpeed    float64
		MouseSensitivity float64
		Zoom             float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Camera{
				Position:         tt.fields.Position,
				Front:            tt.fields.Front,
				Up:               tt.fields.Up,
				Right:            tt.fields.Right,
				WorldUp:          tt.fields.WorldUp,
				Yaw:              tt.fields.Yaw,
				Pitch:            tt.fields.Pitch,
				MovementSpeed:    tt.fields.MovementSpeed,
				MouseSensitivity: tt.fields.MouseSensitivity,
				Zoom:             tt.fields.Zoom,
			}
			c.UpdateCameraVectors()
		})
	}
}
