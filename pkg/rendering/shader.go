package rendering

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Shader struct {
	Program    uint32
	Uniforms   map[string]int32
	Attributes map[string]uint32
}

func NewShader(vertFile, fragFile, geomFile string) (Shader, error) {
	var shader Shader
	vertSrc, err := readFile(vertFile)
	if err != nil {
		return shader, err
	}

	fragSrc, err := readFile(fragFile)
	if err != nil {
		return shader, err
	}

	var geomSrc []byte
	if geomFile != "" {
		geomSrc, err = readFile(geomFile)
		if err != nil {
			return shader, err
		}
	}

	p, err := createProgram(vertSrc, fragSrc, geomSrc)
	if err != nil {
		return shader, err
	}
	shader = setupShader(p)
	return shader, nil
}

func (s Shader) Use() {
	gl.UseProgram(s.Program)
}

func (s *Shader) SetInt32(name string, value int32) {
	gl.Uniform1i(s.Uniforms[name], value)
}

func (s Shader) SetFloat32(name string, value float32) {
	gl.Uniform1f(s.Uniforms[name], value)
}

func (s Shader) SetFloat64(name string, value float64) {
	gl.Uniform1d(s.Uniforms[name], value)
}

func (s Shader) SetVec2(name string, value mgl32.Vec2) {
	gl.Uniform2f(s.Uniforms[name], value.X(), value.Y())
}

func (s Shader) SetVec3(name string, value mgl32.Vec3) {
	gl.Uniform3f(s.Uniforms[name], value.X(), value.Y(), value.Z())
}

func (s Shader) SetVec4(name string, value mgl32.Vec4) {
	gl.Uniform4f(s.Uniforms[name], value.X(), value.Y(), value.Z(), value.W())
}

func (s Shader) SetMat2(name string, value mgl32.Mat2) {
	gl.UniformMatrix2fv(s.Uniforms[name], 1, false, &value[0])
}

func (s Shader) SetMat3(name string, value mgl32.Mat3) {
	gl.UniformMatrix3fv(s.Uniforms[name], 1, false, &value[0])
}

func (s Shader) SetMat4(name string, value mgl32.Mat4) {
	gl.UniformMatrix4fv(s.Uniforms[name], 1, false, &value[0])
}

func (s *Shader) Delete() {
	gl.DeleteProgram(s.Program)
}

/*

Shader setup code taken from https://github.com/raedatoui/glutils/

*/

func setupShader(program uint32) Shader {
	var (
		c int32
		i uint32
	)
	gl.UseProgram(program)
	uniforms := map[string]int32{}
	attributes := map[string]uint32{} // make(map[string]uint32)

	var size int32
	var xtype uint32

	gl.GetProgramiv(program, gl.ACTIVE_UNIFORMS, &c)
	for i = 0; i < uint32(c); i++ {
		var buf [256]byte
		gl.GetActiveUniform(program, i, 256, nil, &size, &xtype, &buf[0])
		loc := gl.GetUniformLocation(program, &buf[0])
		name := gl.GoStr(&buf[0])
		uniforms[name] = loc
	}

	gl.GetProgramiv(program, gl.ACTIVE_ATTRIBUTES, &c)
	for i = 0; i < uint32(c); i++ {
		var buf [256]byte
		gl.GetActiveAttrib(program, i, 256, nil, &size, &xtype, &buf[0])
		loc := gl.GetAttribLocation(program, &buf[0])
		name := gl.GoStr(&buf[0])
		attributes[name] = uint32(loc)
	}

	return Shader{
		Program:    program,
		Uniforms:   uniforms,
		Attributes: attributes,
	}
}

func createProgram(v, f, g []byte) (uint32, error) {
	var p, vertex, frag, geom uint32
	use_geom := false

	if val, err := compileShader(string(v)+"\x00", gl.VERTEX_SHADER); err != nil {
		return 0, err
	} else {
		vertex = val
		defer func(s uint32) { deleteShader(p, s) }(vertex)
	}

	if val, err := compileShader(string(f)+"\x00", gl.FRAGMENT_SHADER); err != nil {
		return 0, err
	} else {
		frag = val
		defer func(s uint32) { deleteShader(p, s) }(frag)
	}

	if len(g) > 0 {
		if val, err := compileShader(string(g)+"\x00", gl.GEOMETRY_SHADER); err != nil {
			return 0, err
		} else {
			geom = val
			defer func(s uint32) { deleteShader(p, s) }(geom)
		}
	}

	p, err := linkProgram(vertex, frag, geom, use_geom)
	if err != nil {
		return 0, err
	}

	return p, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func deleteShader(p, s uint32) {
	gl.DetachShader(p, s)
	gl.DeleteShader(s)
}

func linkProgram(v, f, g uint32, use_geom bool) (uint32, error) {
	program := gl.CreateProgram()
	gl.AttachShader(program, v)
	gl.AttachShader(program, f)
	if use_geom {
		gl.AttachShader(program, g)
	}

	gl.LinkProgram(program)
	// check for program linking errors
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	return program, nil
}

func readFile(file string) ([]byte, error) {
	reader, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return data, err
}

/*func BasicProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}*/
