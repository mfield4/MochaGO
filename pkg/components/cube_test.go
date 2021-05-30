package components

import (
	"reflect"
	"testing"
)

func TestCubeVerticies(t *testing.T) {
	tests := []struct {
		name string
		want [180]float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CubeVertices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CubeVertices() = %v, want %v", got, tt.want)
			}
		})
	}
}
