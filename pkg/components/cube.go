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

func CubeVerticies() [180]float32 {

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
