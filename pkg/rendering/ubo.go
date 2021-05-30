package rendering

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Ubo struct {
	uboID  uint32
	stride int
}

func NewViewProj(size int) Ubo {
	u := Ubo{
		stride: len(mgl32.Mat4{}) * 4,
	}

	gl.GenBuffers(1, &u.uboID)
	gl.BindBuffer(gl.UNIFORM_BUFFER, u.uboID)
	gl.BufferData(gl.UNIFORM_BUFFER, size*u.stride, nil, gl.STATIC_DRAW)

	gl.BindBufferBase(gl.UNIFORM_BUFFER, 0, u.uboID)
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
	return u
}

func (u Ubo) SetMat4(value mgl32.Mat4, idx int) {
	gl.BindBuffer(gl.UNIFORM_BUFFER, u.uboID)
	gl.BufferSubData(gl.UNIFORM_BUFFER, idx*u.stride, u.stride, gl.Ptr(&value[0]))
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}
