package components

type FACES uint8

const (
	INACTIVE FACES = 1 << iota
	ACTIVE
	BACK
	FRONT
	LEFT
	RIGHT
	BOTTOM
	TOP
)

type CUBE_T uint8

const (
	NONE CUBE_T = iota
	GRASS
	DIRT
	STONE
	COAL
	IRON
	COPPER
)

type Cube struct {
	TypeOf      CUBE_T
	ActiveFaces FACES
}

func CubeVertices() [108]float32 {
	// 3 floats per vertex
	// 6 vertices
	// 6 faces
	return [108]float32{
		// Back face
		-0.5, -0.5, -0.5, // Bottom-let
		0.5, 0.5, -0.5, // top-right
		0.5, -0.5, -0.5, // bottom-right
		0.5, 0.5, -0.5, // top-right
		-0.5, -0.5, -0.5, // bottom-let
		-0.5, 0.5, -0.5, // top-let
		// ront ace
		-0.5, -0.5, 0.5, // bottom-let
		0.5, -0.5, 0.5, // bottom-right
		0.5, 0.5, 0.5, // top-right
		0.5, 0.5, 0.5, // top-right
		-0.5, 0.5, 0.5, // top-let
		-0.5, -0.5, 0.5, // bottom-let
		// Let ace
		-0.5, 0.5, 0.5, // top-right
		-0.5, 0.5, -0.5, // top-let
		-0.5, -0.5, -0.5, // bottom-let
		-0.5, -0.5, -0.5, // bottom-let
		-0.5, -0.5, 0.5, // bottom-right
		-0.5, 0.5, 0.5, // top-right
		// Right ace
		0.5, 0.5, 0.5, // top-let
		0.5, -0.5, -0.5, // bottom-right
		0.5, 0.5, -0.5, // top-right
		0.5, -0.5, -0.5, // bottom-right
		0.5, 0.5, 0.5, // top-let
		0.5, -0.5, 0.5, // bottom-let
		// Bottom ace
		-0.5, -0.5, -0.5, // top-right
		0.5, -0.5, -0.5, // top-let
		0.5, -0.5, 0.5, // bottom-let
		0.5, -0.5, 0.5, // bottom-let
		-0.5, -0.5, 0.5, // bottom-right
		-0.5, -0.5, -0.5, // top-right
		// Top ace
		-0.5, 0.5, -0.5, // top-let
		0.5, 0.5, 0.5, // bottom-right
		0.5, 0.5, -0.5, // top-right
		0.5, 0.5, 0.5, // bottom-right
		-0.5, 0.5, -0.5, // top-let
		-0.5, 0.5, 0.5, // bottom-let
	}
}

func CubeVerticesTex() [180]float32 {

	// 5 floats per vertex
	// 6 vertices per face
	// 6 faces
	// 5 * 6 * 6 => 180

	return [180]float32{
		// Back face
		-0.5, -0.5, -0.5, 0.0, 0.0, // Bottom-let
		0.5, 0.5, -0.5, 1.0, 1.0, // top-right
		0.5, -0.5, -0.5, 1.0, 0.0, // bottom-right
		0.5, 0.5, -0.5, 1.0, 1.0, // top-right
		-0.5, -0.5, -0.5, 0.0, 0.0, // bottom-let
		-0.5, 0.5, -0.5, 0.0, 1.0, // top-let
		// ront ace
		-0.5, -0.5, 0.5, 0.0, 0.0, // bottom-let
		0.5, -0.5, 0.5, 1.0, 0.0, // bottom-right
		0.5, 0.5, 0.5, 1.0, 1.0, // top-right
		0.5, 0.5, 0.5, 1.0, 1.0, // top-right
		-0.5, 0.5, 0.5, 0.0, 1.0, // top-let
		-0.5, -0.5, 0.5, 0.0, 0.0, // bottom-let
		// Let ace
		-0.5, 0.5, 0.5, 1.0, 0.0, // top-right
		-0.5, 0.5, -0.5, 1.0, 1.0, // top-let
		-0.5, -0.5, -0.5, 0.0, 1.0, // bottom-let
		-0.5, -0.5, -0.5, 0.0, 1.0, // bottom-let
		-0.5, -0.5, 0.5, 0.0, 0.0, // bottom-right
		-0.5, 0.5, 0.5, 1.0, 0.0, // top-right
		// Right ace
		0.5, 0.5, 0.5, 1.0, 0.0, // top-let
		0.5, -0.5, -0.5, 0.0, 1.0, // bottom-right
		0.5, 0.5, -0.5, 1.0, 1.0, // top-right
		0.5, -0.5, -0.5, 0.0, 1.0, // bottom-right
		0.5, 0.5, 0.5, 1.0, 0.0, // top-let
		0.5, -0.5, 0.5, 0.0, 0.0, // bottom-let
		// Bottom ace
		-0.5, -0.5, -0.5, 0.0, 1.0, // top-right
		0.5, -0.5, -0.5, 1.0, 1.0, // top-let
		0.5, -0.5, 0.5, 1.0, 0.0, // bottom-let
		0.5, -0.5, 0.5, 1.0, 0.0, // bottom-let
		-0.5, -0.5, 0.5, 0.0, 0.0, // bottom-right
		-0.5, -0.5, -0.5, 0.0, 1.0, // top-right
		// Top ace
		-0.5, 0.5, -0.5, 0.0, 1.0, // top-let
		0.5, 0.5, 0.5, 1.0, 0.0, // bottom-right
		0.5, 0.5, -0.5, 1.0, 1.0, // top-right
		0.5, 0.5, 0.5, 1.0, 0.0, // bottom-right
		-0.5, 0.5, -0.5, 0.0, 1.0, // top-let
		-0.5, 0.5, 0.5, 0.0, 0.0, // bottom-let
	}
}
