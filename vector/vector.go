package vector

import (
	"fmt"
	"math"
)

type Vector struct {
	x, y, z, w float64
}

func (v *Vector) X() float64 { return v.x }
func (v *Vector) Y() float64 { return v.y }
func (v *Vector) Z() float64 { return v.z }
func (v *Vector) W() float64 { return v.w }

func (v *Vector) ToString() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", v.x, v.y, v.z, v.w)
}

func NewRawVector(x, y, z, w float64) Vector { return Vector{x: x, y: y, z: z, w: w} }
func NewVector(x, y, z float64) Vector       { return Vector{x: x, y: y, z: z, w: 0.0} }
func NewPoint(x, y, z float64) Vector        { return Vector{x: x, y: y, z: z, w: 1.0} }

func ZeroVector() Vector { return NewVector(0.0, 0.0, 0.0) }
func ZeroPoint() Vector  { return NewPoint(0.0, 0.0, 0.0) }

func Add(v1, v2 Vector) Vector {
	return Vector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
		z: v1.z + v2.z,
		w: v1.w + v2.w,
	}
}

func Subtract(v1, v2 Vector) Vector {
	return Vector{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
		z: v1.z - v2.z,
		w: v1.w - v2.w,
	}
}

func Negate(v Vector) Vector { return NewVector(-v.x, -v.y, -v.z) }

func Multiply(v Vector, value float64) Vector {
	return Vector{
		x: v.x * value,
		y: v.y * value,
		z: v.z * value,
		w: v.w * value,
	}
}

func Divide(v Vector, value float64) Vector {
	return Vector{
		x: v.x / value,
		y: v.y / value,
		z: v.z / value,
		w: v.w / value,
	}
}

func Magnitude(v Vector) float64 {
	val := (v.x * v.x) + (v.y * v.y) + (v.z * v.z)
	return math.Sqrt(val)
}

func Normalize(v Vector) Vector {
	return Divide(v, Magnitude(v))
}

func Dot(v1, v2 Vector) float64 {
	return (v1.x * v2.x) + (v1.y * v2.y) + (v1.z * v2.z) + (v1.w * v2.w)
}

func Cross(v1, v2 Vector) Vector {
	return NewVector(
		(v1.y*v2.z)-(v1.z*v2.y),
		(v1.z*v2.x)-(v1.x*v2.z),
		(v1.x*v2.y)-(v1.y*v2.x))
}
