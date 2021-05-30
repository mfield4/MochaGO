package rendering

import "testing"

func TestEnableAttributes(t *testing.T) {
	type args struct {
		vao        uint32
		vbo        uint32
		attributes []Attribute
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

func TestInitNilVBO(t *testing.T) {
	type args struct {
		bufLen int
		attrs  []Attribute
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
			gotVao, gotVbo := InitNilVBO(tt.args.bufLen, tt.args.attrs)
			if gotVao != tt.wantVao {
				t.Errorf("InitNilVBO() gotVao = %v, want %v", gotVao, tt.wantVao)
			}
			if gotVbo != tt.wantVbo {
				t.Errorf("InitNilVBO() gotVbo = %v, want %v", gotVbo, tt.wantVbo)
			}
		})
	}
}

func TestInitNilVBOPos(t *testing.T) {
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
			gotVao, gotVbo := InitNilVBOPos(tt.args.bufLen)
			if gotVao != tt.wantVao {
				t.Errorf("InitNilVBOPos() gotVao = %v, want %v", gotVao, tt.wantVao)
			}
			if gotVbo != tt.wantVbo {
				t.Errorf("InitNilVBOPos() gotVbo = %v, want %v", gotVbo, tt.wantVbo)
			}
		})
	}
}

func TestInitNilVBOTex(t *testing.T) {
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
			gotVao, gotVbo := InitNilVBOTex(tt.args.bufLen)
			if gotVao != tt.wantVao {
				t.Errorf("InitNilVBOTex() gotVao = %v, want %v", gotVao, tt.wantVao)
			}
			if gotVbo != tt.wantVbo {
				t.Errorf("InitNilVBOTex() gotVbo = %v, want %v", gotVbo, tt.wantVbo)
			}
		})
	}
}

func TestInitVBO(t *testing.T) {
	type args struct {
		points []float32
		attrs  []Attribute
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
			gotVao, gotVbo := InitVBO(tt.args.points, tt.args.attrs)
			if gotVao != tt.wantVao {
				t.Errorf("InitVBO() gotVao = %v, want %v", gotVao, tt.wantVao)
			}
			if gotVbo != tt.wantVbo {
				t.Errorf("InitVBO() gotVbo = %v, want %v", gotVbo, tt.wantVbo)
			}
		})
	}
}

func TestInitVBOPos(t *testing.T) {
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
			gotVao, gotVbo := InitVBOPos(tt.args.points)
			if gotVao != tt.wantVao {
				t.Errorf("InitVBOPos() gotVao = %v, want %v", gotVao, tt.wantVao)
			}
			if gotVbo != tt.wantVbo {
				t.Errorf("InitVBOPos() gotVbo = %v, want %v", gotVbo, tt.wantVbo)
			}
		})
	}
}

func TestInitVBOTex(t *testing.T) {
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
			gotVao, gotVbo := InitVBOTex(tt.args.points)
			if gotVao != tt.wantVao {
				t.Errorf("InitVBOTex() gotVao = %v, want %v", gotVao, tt.wantVao)
			}
			if gotVbo != tt.wantVbo {
				t.Errorf("InitVBOTex() gotVbo = %v, want %v", gotVbo, tt.wantVbo)
			}
		})
	}
}
