package systems

import (
	"reflect"
	"testing"
)

func TestNewSkyboxRenderer(t *testing.T) {
	tests := []struct {
		name string
		want *SkyboxRenderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSkyboxRenderer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSkyboxRenderer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkyboxRenderer_Render(t *testing.T) {
	type args struct {
		dt float32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SkyboxRenderer{}
			s.Render(0)
		})
	}
}
