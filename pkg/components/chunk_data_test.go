package components

import (
	"reflect"
	"testing"
)

func TestNewChunk(t *testing.T) {
	tests := []struct {
		name string
		want *Chunk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChunk(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewChunkWithCubes(t *testing.T) {
	type args struct {
		cubes [CHUNK_SIZE][CHUNK_SIZE][CHUNK_SIZE]Cube
	}
	tests := []struct {
		name string
		args args
		want *Chunk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChunkWithCubes(tt.args.cubes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChunkWithCubes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRebuildChunkWithUpdate(t *testing.T) {
	type args struct {
		chunk  *Chunk
		update ChunkUpdate
	}
	tests := []struct {
		name string
		args args
		want *Chunk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RebuildChunkWithUpdate(tt.args.chunk, tt.args.update); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RebuildChunkWithUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
