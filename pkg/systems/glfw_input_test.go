package systems

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"go_entity_component_services/pkg/components"
	"reflect"
	"testing"
)

func TestGlfwInput_AddKeyListeners(t *testing.T) {
	type fields struct {
		cursorPositionListeners []func(components.CursorPositionCommand)
		mouseButtonListeners    []func(components.MouseButtonCommand)
		keyListeners            []func(components.KeyCommand)
	}
	type args struct {
		funcs []func(components.KeyCommand)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// g := &GlfwInput{
			// 	cursorPositionListeners: tt.fields.cursorPositionListeners,
			// 	mouseButtonListeners:    tt.fields.mouseButtonListeners,
			// 	keyListeners:            tt.fields.keyListeners,
			// }
		})
	}
}

func TestGlfwInput_AddMouseButtonListeners(t *testing.T) {
	type fields struct {
		cursorPositionListeners []func(components.CursorPositionCommand)
		mouseButtonListeners    []func(components.MouseButtonCommand)
		keyListeners            []func(components.KeyCommand)
	}
	type args struct {
		funcs []func(components.MouseButtonCommand)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// g := &GlfwInput{
			// 	cursorPositionListeners: tt.fields.cursorPositionListeners,
			// 	mouseButtonListeners:    tt.fields.mouseButtonListeners,
			// 	keyListeners:            tt.fields.keyListeners,
			// }
		})
	}
}

func TestGlfwInput_AddMouseMotionListeners(t *testing.T) {
	type fields struct {
		cursorPositionListeners []func(components.CursorPositionCommand)
		mouseButtonListeners    []func(components.MouseButtonCommand)
		keyListeners            []func(components.KeyCommand)
	}
	type args struct {
		funcs []func(components.CursorPositionCommand)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// g := &GlfwInput{
			// 	cursorPositionListeners: tt.fields.cursorPositionListeners,
			// 	mouseButtonListeners:    tt.fields.mouseButtonListeners,
			// 	keyListeners:            tt.fields.keyListeners,
			// }
		})
	}
}

func TestGlfwInput_Poll(t *testing.T) {
	type fields struct {
		cursorPositionListeners []func(components.CursorPositionCommand)
		mouseButtonListeners    []func(components.MouseButtonCommand)
		keyListeners            []func(components.KeyCommand)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// g := &GlfwInput{
			// 	cursorPositionListeners: tt.fields.cursorPositionListeners,
			// 	mouseButtonListeners:    tt.fields.mouseButtonListeners,
			// 	keyListeners:            tt.fields.keyListeners,
			// }
		})
	}
}

func TestGlfwInput_cursorPositionCallback(t *testing.T) {
	type fields struct {
		cursorPositionListeners []func(components.CursorPositionCommand)
		mouseButtonListeners    []func(components.MouseButtonCommand)
		keyListeners            []func(components.KeyCommand)
	}
	type args struct {
		window *glfw.Window
		xpos   float64
		ypos   float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// g := &GlfwInput{
			// 	cursorPositionListeners: tt.fields.cursorPositionListeners,
			// 	mouseButtonListeners:    tt.fields.mouseButtonListeners,
			// 	keyListeners:            tt.fields.keyListeners,
			// }
		})
	}
}

func TestGlfwInput_keyCallback(t *testing.T) {
	type fields struct {
		cursorPositionListeners []func(components.CursorPositionCommand)
		mouseButtonListeners    []func(components.MouseButtonCommand)
		keyListeners            []func(components.KeyCommand)
	}
	type args struct {
		window   *glfw.Window
		key      glfw.Key
		scancode int
		action   glfw.Action
		mode     glfw.ModifierKey
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// g := &GlfwInput{
			// 	cursorPositionListeners: tt.fields.cursorPositionListeners,
			// 	mouseButtonListeners:    tt.fields.mouseButtonListeners,
			// 	keyListeners:            tt.fields.keyListeners,
			// }
		})
	}
}

func TestGlfwInput_mouseButtonCallback(t *testing.T) {
	type fields struct {
		cursorPositionListeners []func(components.CursorPositionCommand)
		mouseButtonListeners    []func(components.MouseButtonCommand)
		keyListeners            []func(components.KeyCommand)
	}
	type args struct {
		window *glfw.Window
		button glfw.MouseButton
		action glfw.Action
		mods   glfw.ModifierKey
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// g := &GlfwInput{
			// 	cursorPositionListeners: tt.fields.cursorPositionListeners,
			// 	mouseButtonListeners:    tt.fields.mouseButtonListeners,
			// 	keyListeners:            tt.fields.keyListeners,
			// }
		})
	}
}

func TestNewGlfwInput(t *testing.T) {
	type args struct {
		window *glfw.Window
	}
	tests := []struct {
		name string
		args args
		want *GlfwInput
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGlfwInput(tt.args.window); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGlfwInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
