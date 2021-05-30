package components

import (
	"go_entity_component_services/pkg/rendering"
)

const (
	CHUNK_SIZE = 16
)

type Chunk struct {
	Cubes        [CHUNK_SIZE][CHUNK_SIZE][CHUNK_SIZE]Cube
	ShowingFaces int
	VAO, VBO     uint32
}

func NewChunk() *Chunk {
	vao, vbo := rendering.InitNilVBOTex(1000)

	return &Chunk{
		Cubes:        [16][16][16]Cube{},
		ShowingFaces: 0,
		VAO:          vao,
		VBO:          vbo,
	}
}

func NewChunkWithCubes(cubes [CHUNK_SIZE][CHUNK_SIZE][CHUNK_SIZE]Cube) *Chunk {
	vao, vbo := rendering.InitNilVBOTex(len(cubes))

	return &Chunk{
		Cubes:        cubes,
		ShowingFaces: 0,
		VAO:          vao,
		VBO:          vbo,
	}
}

func RebuildChunkWithUpdate(chunk *Chunk, update ChunkUpdate) *Chunk {
	// REBIND VBO BECUASE OF POSSIBLE SIZE DIFFERENCE??

	return &Chunk{
		Cubes:        [16][16][16]Cube{},
		ShowingFaces: 0,
		VAO:          0,
		VBO:          0,
	}
}

/* Chunk TODO list
constructor that figures out how to load one

fill each chunk with necessary graphics data
*/
