package components

import "github.com/go-gl/glfw/v3.3/glfw"

type CursorPositionCommand struct {
	Xpos, Ypos float64
}

type MouseButtonCommand struct {
}

type MouseZoomCommand struct {
}

type KeyCommand struct {
	glfw.Key
}

type ChunkUpdate struct {
}
