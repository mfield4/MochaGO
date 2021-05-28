package components

import (
	"reflect"
	"testing"
)

func TestNewCubeMap(t *testing.T) {
	type args struct {
		wrap_s int32
		wrap_t int32
		wrap_r int32
		min_f  int32
		mag_f  int32
		faces  [6]string
	}
	tests := []struct {
		name          string
		args          args
		wantCubeMapId uint32
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCubeMapId, err := NewCubeMap(tt.args.wrap_s, tt.args.wrap_t, tt.args.wrap_r, tt.args.min_f, tt.args.mag_f, tt.args.faces)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCubeMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCubeMapId != tt.wantCubeMapId {
				t.Errorf("NewCubeMap() gotCubeMapId = %v, want %v", gotCubeMapId, tt.wantCubeMapId)
			}
		})
	}
}

func TestNewCubeMapD(t *testing.T) {
	type args struct {
		faces [6]string
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
			got, err := NewCubeMapD(tt.args.faces)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCubeMapD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewCubeMapD() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTexture(t *testing.T) {
	type args struct {
		wrap_s int32
		wrap_t int32
		min_f  int32
		mag_f  int32
		file   string
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
			got, err := NewTexture(tt.args.wrap_s, tt.args.wrap_t, tt.args.min_f, tt.args.mag_f, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTexture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewTexture() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTextureD(t *testing.T) {
	type args struct {
		file string
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
			got, err := NewTextureD(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTextureD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewTextureD() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imageToPixelData(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *image.RGBA
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := imageToPixelData(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageToPixelData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageToPixelData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
