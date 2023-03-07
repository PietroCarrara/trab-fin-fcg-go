package matrix

import (
	"github.com/chewxy/math32"
	"github.com/go-gl/mathgl/mgl32"
)

func CameraView(camPos, view, up mgl32.Vec3) mgl32.Mat4 {
	w := view.Vec4(0).Mul(-1).Normalize()
	u := cross(up.Vec4(0), w).Normalize()

	v := cross(w, u)

	origin := mgl32.Vec3{}.Vec4(1)

	ux := u.X()
	uy := u.Y()
	uz := u.Z()
	vx := v.X()
	vy := v.Y()
	vz := v.Z()
	wx := w.X()
	wy := w.Y()
	wz := w.Z()

	p := camPos.Vec4(1).Sub(origin)

	return mgl32.Mat4{
		ux, uy, uz, -dot(u, p),
		vx, vy, vz, -dot(v, p),
		wx, wy, wz, -dot(w, p),
		0, 0, 0, 1,
	}.Transpose()
}

func Translate(v mgl32.Vec3) mgl32.Mat4 {
	return mgl32.Mat4{
		1, 0, 0, v.X(),
		0, 1, 0, v.Y(),
		0, 0, 1, v.Z(),
		0, 0, 0, 1,
	}.Transpose()
}

func Scale(s mgl32.Vec3) mgl32.Mat4 {
	x := s.X()
	y := s.Y()
	z := s.Z()

	return mgl32.Mat4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}.Transpose()
}

func RotateX(r float32) mgl32.Mat4 {
	c := math32.Cos(r)
	s := math32.Sin(r)

	return mgl32.Mat4{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}.Transpose()
}

func RotateY(r float32) mgl32.Mat4 {
	c := math32.Cos(r)
	s := math32.Sin(r)

	return mgl32.Mat4{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}.Transpose()
}

func RotateZ(r float32) mgl32.Mat4 {
	c := math32.Cos(r)
	s := math32.Sin(r)

	return mgl32.Mat4{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}.Transpose()
}

func Orthographic(l, r, b, t, n, f float32) mgl32.Mat4 {
	return mgl32.Mat4{
		2 / (r - l), 0, 0, -(r + l) / (r - l),
		0, 2 / (t - b), 0, -(t + b) / (t - b),
		0, 0, 2 / (f - n), -(f + n) / (f - n),
		0, 0, 0, 1,
	}.Transpose()
}

func Perspective(fov, aspect, n, f float32) mgl32.Mat4 {
	{
		t := math32.Abs(n) * math32.Tan(fov/2)
		b := -t
		r := t * aspect
		l := -r

		P := mgl32.Mat4{
			n, 0, 0, 0,
			0, n, 0, 0,
			0, 0, n + f, -f * n,
			0, 0, 1, 0,
		}.Transpose()

		M := Orthographic(l, r, b, t, n, f)

		return M.Mul(-1).Mul4(P)
	}
}
