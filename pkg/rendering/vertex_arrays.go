package rendering

import "github.com/go-gl/gl/v4.1-core/gl"

type Attribute struct {
	idx          uint32
	size, stride int32
	offset       uintptr
}

// InitVBOPos : This function is a default wrapper around InitVBO that initializes only the position vertex attribute
func InitVBOPos(points []float32) (vao, vbo uint32) {
	return InitVBO(points, []Attribute{{idx: 0, size: 3, stride: 3 * 4, offset: 0}})
}

// InitVBOTex : This function is a default wrapper around InitVBO that initializes the position and the texture vertex attribute
func InitVBOTex(points []float32) (vao, vbo uint32) {
	return InitVBO(points, []Attribute{
		{idx: 0, size: 3, stride: 3 * 4, offset: 0},
		{idx: 1, size: 2, stride: 3 * 4, offset: 3 * 4}})
}

// InitNilVBOPos : This function is a default wrapper around InitVBO that initializes only the position vertex attribute
func InitNilVBOPos(bufLen int) (vao, vbo uint32) {
	return InitNilVBO(bufLen, []Attribute{{idx: 0, size: 3, stride: 3 * 4, offset: 0}})
}

// InitNilVBOTex : This function is a default wrapper around InitNilVBO that initializes the position and the texture vertex attribute
func InitNilVBOTex(bufLen int) (vao, vbo uint32) {
	return InitNilVBO(bufLen, []Attribute{
		{idx: 0, size: 3, stride: 3 * 4, offset: 0},
		{idx: 1, size: 2, stride: 3 * 4, offset: 3 * 4}})
}

// InitVBO
// Initializes a gl vertex array and buffer object with given values. Size will be implicitly deduced
// points : The data to bind to the generated buffer
// attrs  : variadic arguement containing the vertex attributes to bind to the generated buffers
func InitVBO(points []float32, attrs []Attribute) (vao, vbo uint32) {
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	// pos vertex
	enableAttributes(vao, vbo, attrs)

	return vao, vbo
}

// InitNilVBO
// Initializes a gl vertex array and buffer object with a given initial size and no initial values.
// This function allows the user to later update the buffer with gl.BindSubBufferData.
// bufLen : length of the generated buffer object, in bytes
// attrs  : variadic arguement containing the vertex attributes to bind to the generated buffers
func InitNilVBO(bufLen int, attrs []Attribute) (vao, vbo uint32) {
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*bufLen, nil, gl.STATIC_DRAW)

	enableAttributes(vao, vbo, attrs)

	return vao, vbo
}

// enableAttributes iterates through the given attributes and enables them for the given vao and vbo
// Assumes that both the vao and the vbo are bound
func enableAttributes(vao, vbo uint32, attributes []Attribute) {
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	for idx, attr := range attributes {
		gl.EnableVertexAttribArray(uint32(idx))
		gl.VertexAttribPointerWithOffset(uint32(idx), attr.size, gl.FLOAT, false, attr.stride, attr.offset)
	}

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
