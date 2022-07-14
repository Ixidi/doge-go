package math

import "math"

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vector3) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}
