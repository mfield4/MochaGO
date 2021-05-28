package components

import "github.com/go-gl/gl/v4.1-core/gl"

const (
	CHUNK_SIZE = 16
)

type Chunk struct {
	Cubes        [CHUNK_SIZE][CHUNK_SIZE][CHUNK_SIZE]Cube
	ShowingFaces int
	VAO, VBO     uint32
}

func NewChunk() *Chunk {
	vao, vbo := initNilVBuf(1000)

	return &Chunk{
		Cubes:        [16][16][16]Cube{},
		ShowingFaces: 0,
		VAO:          vao,
		VBO:          vbo,
	}
}

func NewChunkWithCubes(cubes [CHUNK_SIZE][CHUNK_SIZE][CHUNK_SIZE]Cube) *Chunk {

	vao, vbo := initNilVBuf(len(cubes))

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

build vao method
*/

// Chunk constructors of various types
// makeVao initializes and returns a vertex array from the points provided.

func initVBuf(points []float32) (vao, vbo uint32) {
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	// pos vertex
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
	// texture coords
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)

	return vao, vbo
}

func initNilVBuf(bufLen int) (vao, vbo uint32) {
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*bufLen, nil, gl.STATIC_DRAW)

	// pos vertex
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
	// texture coords
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)

	return vao, vbo
}
