package rendering

import (
	"reflect"
	"testing"
)

func TestNewShader(t *testing.T) {
	type args struct {
		vertFile string
		fragFile string
		geomFile string
	}
	tests := []struct {
		name    string
		args    args
		want    Shader
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewShader(tt.args.vertFile, tt.args.fragFile, tt.args.geomFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewShader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShader() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShader_Delete(t *testing.T) {
	type fields struct {
		Program    uint32
		Uniforms   map[string]int32
		Attributes map[string]uint32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Shader{
				Program:    tt.fields.Program,
				Uniforms:   tt.fields.Uniforms,
				Attributes: tt.fields.Attributes,
			}
			s.Delete()
		})
	}
}

func Test_compileShader(t *testing.T) {
	type args struct {
		source     string
		shaderType uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compileShader(tt.args.source, tt.args.shaderType)
			if (err != nil) != tt.wantErr {
				t.Errorf("compileShader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("compileShader() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createProgram(t *testing.T) {
	type args struct {
		v []byte
		f []byte
		g []byte
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createProgram(tt.args.v, tt.args.f, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("createProgram() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createProgram() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteShader(t *testing.T) {
	type args struct {
		p uint32
		s uint32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_linkProgram(t *testing.T) {
	type args struct {
		v        uint32
		f        uint32
		g        uint32
		use_geom bool
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := linkProgram(tt.args.v, tt.args.f, tt.args.g, tt.args.use_geom)
			if (err != nil) != tt.wantErr {
				t.Errorf("linkProgram() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("linkProgram() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setupShader(t *testing.T) {
	type args struct {
		program uint32
	}
	tests := []struct {
		name string
		args args
		want Shader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setupShader(tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupShader() = %v, want %v", got, tt.want)
			}
		})
	}
}
