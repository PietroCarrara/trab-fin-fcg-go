package main

import (
	"fcg/trab/pkg/camera"
	"fcg/trab/pkg/entities"
	"fcg/trab/pkg/graphics"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/neclepsio/gl/v3.3-core/gl"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(480, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		panic(err)
	}
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	gl.FrontFace(gl.CCW)

	err = graphics.Init()
	if err != nil {
		panic(err)
	}

	cam := camera.LookAtCamera{
		Distance: 3.5,
	}

	c1 := entities.MakeCube()

	for !window.ShouldClose() {
		// Update
		c1.Update(1 / 60.0)

		// Draw
		gl.ClearColor(100/255.0, 149/255.0, 237/255.0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		c1.Draw(&cam)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
