package pt

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func V(X, Y, Z float64) Vector {
	return Vector{X, Y, Z}
}
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vector) Cross(b Vector) Vector {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return Vector{x, y, z}
}
func (a Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func (a Vector) Normalize() Vector {
	d := a.Length()
	return Vector{a.X / d, a.Y / d, a.Z / d}
}

func (a Vector) Negate() Vector {
	return Vector{-a.X, -a.Y, -a.Z}
}

func (a Vector) Abs() Vector {
	return Vector{math.Abs(a.X), math.Abs(a.Y), math.Abs(a.Z)}
}

func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vector) Mul(b Vector) Vector {
	return Vector{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (a Vector) Div(b Vector) Vector {
	return Vector{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}
func (a Vector) MulScalar(b float64) Vector {
	return Vector{a.X * b, a.Y * b, a.Z * b}
}

func (a Vector) DivScalar(b float64) Vector {
	return Vector{a.X / b, a.Y / b, a.Z / b}
}
func (a Vector) AddScalar(b float64) Vector {
	return Vector{a.X + b, a.Y + b, a.Z + b}
}

func (a Vector) SubScalar(b float64) Vector {
	return Vector{a.X - b, a.Y - b, a.Z - b}
}
func (a Vector) Min(b Vector) Vector {
	return Vector{math.Min(a.X, b.X), math.Min(a.Y, b.Y), math.Min(a.Z, b.Z)}
}

func (a Vector) Max(b Vector) Vector {
	return Vector{math.Max(a.X, b.X), math.Max(a.Y, b.Y), math.Max(a.Z, b.Z)}
}
func (n Vector) Reflect(i Vector) Vector {
	return i.Sub(n.MulScalar(2 * n.Dot(i)))
}
func (n Vector) Refract(i Vector, n1, n2 float64) Vector {
	nr := n1 / n2
	cosI := -n.Dot(i)
	sinT2 := nr * nr * (1 - cosI*cosI)
	if sinT2 > 1 {
		return Vector{}
	}
	cosT := math.Sqrt(1 - sinT2)
	return i.MulScalar(nr).Add(n.MulScalar(nr*cosI - cosT))
}

func (n Vector) Reflectance(i Vector, n1, n2 float64) float64 {
	nr := n1 / n2
	cosI := -n.Dot(i)
	sinT2 := nr * nr * (1 - cosI*cosI)
	if sinT2 > 1 {
		return 1
	}
	cosT := math.Sqrt(1 - sinT2)
	rOrth := (n1*cosI - n2*cosT) / (n1*cosI + n2*cosT)
	rPar := (n2*cosI - n1*cosT) / (n2*cosI + n1*cosT)
	return (rOrth*rOrth + rPar*rPar) / 2
}
