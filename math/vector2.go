package math

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

func (v Vector2) Add(x, y float64) Vector2 {
	return Vector2{v.X + x, v.Y + y}
}

func (v Vector2) AddVector(u Vector2) Vector2 {
	return Vector2{v.X + u.X, v.Y + u.Y}
}

func (v Vector2) Subtract(x, y float64) Vector2 {
	return Vector2{v.X - x, v.Y - y}
}

func (v Vector2) SubtractVector(u Vector2) Vector2 {
	return Vector2{v.X - u.X, v.Y - u.Y}
}

func (v Vector2) Multiply(multiplier float64) Vector2 {
	return Vector2{v.X * multiplier, v.Y * multiplier}
}

func (v Vector2) Divide(divisor float64) Vector2 {
	return Vector2{v.X / divisor, v.Y / divisor}
}

func (v Vector2) Normalize() Vector2 {
	return v.Divide(v.Length())
}

func (v Vector2) Scale(length float64) Vector2 {
	return v.Normalize().Multiply(length)
}

func (v Vector2) Rotate(angle float64) Vector2 {
	x := (math.Cos(angle) * v.X) - (math.Sin(angle) * v.Y) + v.X
	y := (math.Sin(angle) * v.X) + (math.Cos(angle) * v.Y) + v.Y
	return Vector2{x, y}
}
