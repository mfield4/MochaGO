package systems

import (
	"reflect"
	"testing"
)

func TestChunkManager_Poll(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// c := &ChunkManager{}
		})
	}
}

func TestChunkManager_Step(t *testing.T) {
	type args struct {
		dt float32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChunkManager{}
			if got := c.Step(tt.args.dt); got != tt.want {
				t.Errorf("Step() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunkManager_Stop(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// c := &ChunkManager{}
		})
	}
}

func TestNewChunkManager(t *testing.T) {
	tests := []struct {
		name string
		want *ChunkManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChunkManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChunkManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
