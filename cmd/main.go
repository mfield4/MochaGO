package main

import (
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	runtime.LockOSThread()
}

func main() {
	game := NewGame(1920, 1080)
	game.Start()
	for game.IsRunning() {
		game.Update()
		game.Draw()
	}
}
