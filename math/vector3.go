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

func (v Vector3) Add(x, y, z float64) Vector3 {
	return Vector3{v.X + x, v.Y + y, v.Z + z}
}

func (v Vector3) AddVector(u Vector3) Vector3 {
	return Vector3{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vector3) Subtract(x, y, z float64) Vector3 {
	return Vector3{v.X - x, v.Y - y, v.Z - z}
}

func (v Vector3) SubtractVector(u Vector3) Vector3 {
	return Vector3{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vector3) Multiply(multiplier float64) Vector3 {
	return Vector3{v.X * multiplier, v.Y * multiplier, v.Z * multiplier}
}

func (v Vector3) Divide(divisor float64) Vector3 {
	return Vector3{v.X / divisor, v.Y / divisor, v.Z / divisor}
}

func (v Vector3) Normalize() Vector3 {
	return v.Divide(v.Length())
}

func (v Vector3) Scale(length float64) Vector3 {
	return v.Normalize().Multiply(length)
}
