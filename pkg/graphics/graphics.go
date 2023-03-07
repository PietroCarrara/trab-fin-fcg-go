package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/neclepsio/gl/v3.3-core/gl"
)

var screenRatio float32 = 1

var shaderID uint32
var modelUniform int32
var viewUniform int32
var projectionUniform int32

var fov float32 = 3.141592 / 3
var near float32 = 0.01
var far float32 = 10

var perspectiveProjection mgl32.Mat4 = mgl32.Perspective(fov, 1, near, far)

func Init() error {
	var err error
	shaderID, err = loadShaders()
	if err != nil {
		return err
	}

	modelUniform = gl.GetUniformLocation(shaderID, gl.Str("model\x00"))
	viewUniform = gl.GetUniformLocation(shaderID, gl.Str("view\x00"))
	projectionUniform = gl.GetUniformLocation(shaderID, gl.Str("projection\x00"))

	return nil
}

func DrawElements(model, view mgl32.Mat4, vertexArrayID uint32, drawMode uint32, count int32, _type uint32) {
	gl.UseProgram(shaderID)

	gl.BindVertexArray(vertexArrayID)

	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
	gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])
	gl.UniformMatrix4fv(projectionUniform, 1, false, &perspectiveProjection[0])

	gl.DrawElements(
		drawMode,
		count,
		_type,
		nil,
	)

	gl.BindVertexArray(0)
}
