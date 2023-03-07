package entities

import (
	"fcg/trab/pkg/camera"
	"fcg/trab/pkg/graphics"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var vertexArrayID uint32 = 0

type Cube struct {
	Position mgl32.Vec3
	Scale    mgl32.Vec3
	Rotation mgl32.Vec3
}

func MakeCube() Cube {
	return Cube{
		Scale: mgl32.Vec3{1, 1, 1},
	}
}

func (c *Cube) Draw(cam camera.Camera) {
	if vertexArrayID == 0 {
		vertexArrayID = buildTriangles()
	}

	model := mgl32.Translate3D(c.Position.X(), c.Position.Y(), c.Position.Z()).
		Mul4(mgl32.Scale3D(c.Scale.X(), c.Scale.Y(), c.Scale.Z())).
		Mul4(mgl32.HomogRotate3DX(c.Rotation.X())).
		Mul4(mgl32.HomogRotate3DY(c.Rotation.Y())).
		Mul4(mgl32.HomogRotate3DZ(c.Rotation.Z()))

	graphics.DrawElements(model, cam.GetMatrix(), vertexArrayID, gl.TRIANGLES, 36, gl.UNSIGNED_INT)
}

func (c *Cube) Update(dt float32) {
	c.Rotation[0] += dt
	c.Rotation[1] += dt
	c.Rotation[2] += dt
}

func buildTriangles() uint32 {
	modelVertices := []float32{
		-0.5, 0.5, 0.5, 1.0,
		-0.5, -0.5, 0.5, 1.0,
		0.5, -0.5, 0.5, 1.0,
		0.5, 0.5, 0.5, 1.0,
		-0.5, 0.5, -0.5, 1.0,
		-0.5, -0.5, -0.5, 1.0,
		0.5, -0.5, -0.5, 1.0,
		0.5, 0.5, -0.5, 1.0,
	}
	var vertexBufferID uint32 = 0
	gl.GenBuffers(1, &vertexBufferID)

	var vertexArrayID uint32 = 0
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)

	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferID)
	gl.BufferData(gl.ARRAY_BUFFER, len(modelVertices)*4, nil, gl.STATIC_DRAW)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(modelVertices)*4, gl.Ptr(modelVertices))
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	colors := []float32{
		1.0, 0.5, 0.0, 1.0,
		1.0, 0.5, 0.0, 1.0,
		0.0, 0.5, 1.0, 1.0,
		0.0, 0.5, 1.0, 1.0,
		1.0, 0.5, 0.0, 1.0,
		1.0, 0.5, 0.0, 1.0,
		0.0, 0.5, 1.0, 1.0,
		0.0, 0.5, 1.0, 1.0,
	}
	var colorsID uint32 = 0
	gl.GenBuffers(1, &colorsID)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorsID)
	gl.BufferData(gl.ARRAY_BUFFER, len(colors)*4, nil, gl.STATIC_DRAW)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(colors)*4, gl.Ptr(colors))
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	indexes := []uint32{
		0, 1, 2,
		7, 6, 5,
		3, 2, 6,
		4, 0, 3,
		4, 5, 1,
		1, 5, 6,
		0, 2, 3,
		7, 5, 4,
		3, 6, 7,
		4, 3, 7,
		4, 1, 0,
		1, 6, 2,
	}
	var indexesID uint32 = 0
	gl.GenBuffers(1, &indexesID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexesID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indexes)*4, nil, gl.STATIC_DRAW)
	gl.BufferSubData(gl.ELEMENT_ARRAY_BUFFER, 0, len(indexes)*4, gl.Ptr(indexes))

	gl.BindVertexArray(0)
	return vertexArrayID
}
