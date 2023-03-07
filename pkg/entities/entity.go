package entities

import "fcg/trab/pkg/camera"

type Entity interface {
	Update(float32)
	Draw(camera.Camera)
}
