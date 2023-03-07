package camera

import (
	"fcg/trab/pkg/matrix"

	"github.com/chewxy/math32"
	"github.com/go-gl/mathgl/mgl32"
)

const max = 3.141592 / 2

type LookAtCamera struct {
	Focus      mgl32.Vec3
	RotX, RotY float32
	Distance   float32
}

func (l *LookAtCamera) GetMatrix() mgl32.Mat4 {

	// if l.RotY > max {
	// 	l.RotY = max
	// } else if l.RotY < -max {
	// 	l.RotY = -max
	// }

	r := l.Distance
	y := r * math32.Sin(l.RotY)
	z := r * math32.Cos(l.RotY) * math32.Cos(l.RotX)
	x := r * math32.Cos(l.RotY) * math32.Sin(l.RotX)

	camPos := mgl32.Vec4{x, y, z, 1}
	// camFocus := l.Focus.Vec4(1)
	camFocus := mgl32.Vec4{0, 0, 0, 1}
	camView := camFocus.Sub(camPos)
	camUp := mgl32.Vec4{0, 1, 0, 0}

	return matrix.CameraView(camPos, camView, camUp)
}
