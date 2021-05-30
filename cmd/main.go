package main

import (
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	runtime.LockOSThread()
}

func main() {
	myGame := NewGame(1920, 1080)
	myGame.Start()
	for myGame.IsRunning() {
		myGame.Update()
		myGame.Draw()
	}
}
