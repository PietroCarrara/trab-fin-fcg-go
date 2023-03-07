package matrix

import "github.com/go-gl/mathgl/mgl32"

func cross(u, v mgl32.Vec4) mgl32.Vec4 {
	u1 := u.X()
	u2 := u.Y()
	u3 := u.Z()
	v1 := v.X()
	v2 := v.Y()
	v3 := v.Z()

	return mgl32.Vec4{
		u2*v3 - u3*v2, // Primeiro coeficiente
		u3*v1 - u1*v3, // Segundo coeficiente
		u1*v2 - u2*v1, // Terceiro coeficiente
		0,             // w = 0 para vetores.
	}
}

func dot(u, v mgl32.Vec4) float32 {
	u1 := u.X()
	u2 := u.Y()
	u3 := u.Z()
	u4 := u.W()
	v1 := v.X()
	v2 := v.Y()
	v3 := v.Z()
	v4 := v.W()

	if u4 != 0 || v4 != 0 {
		panic("dot product is undefined on points!")
	}

	return u1*v1 + u2*v2 + u3*v3
}
