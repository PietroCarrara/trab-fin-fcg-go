package camera

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera interface {
	GetMatrix() mgl32.Mat4
}
