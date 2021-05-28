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

func Test_initNilVBuf(t *testing.T) {
	type args struct {
		bufLen int
	}
	tests := []struct {
		name    string
		args    args
		wantVao uint32
		wantVbo uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVao, gotVbo := initNilVBuf(tt.args.bufLen)
			if gotVao != tt.wantVao {
				t.Errorf("initNilVBuf() gotVao = %v, want %v", gotVao, tt.wantVao)
			}
			if gotVbo != tt.wantVbo {
				t.Errorf("initNilVBuf() gotVbo = %v, want %v", gotVbo, tt.wantVbo)
			}
		})
	}
}

func Test_initVBuf(t *testing.T) {
	type args struct {
		points []float32
	}
	tests := []struct {
		name    string
		args    args
		wantVao uint32
		wantVbo uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVao, gotVbo := initVBuf(tt.args.points)
			if gotVao != tt.wantVao {
				t.Errorf("initVBuf() gotVao = %v, want %v", gotVao, tt.wantVao)
			}
			if gotVbo != tt.wantVbo {
				t.Errorf("initVBuf() gotVbo = %v, want %v", gotVbo, tt.wantVbo)
			}
		})
	}
}
