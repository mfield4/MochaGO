package systems

import (
	"reflect"
	"testing"
)

func TestChunkRenderer_Render(t *testing.T) {
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
			// c := &ChunkRenderer{}
		})
	}
}

func TestNewChunkRenderer(t *testing.T) {
	tests := []struct {
		name string
		want *ChunkRenderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChunkRenderer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChunkRenderer() = %v, want %v", got, tt.want)
			}
		})
	}
}
