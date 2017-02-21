package pt

import (
	"math"
)

type Plane struct {
	Point    Vector
	Normal   Vector
	Material Material
}

func NewPlane(center Vector, normal Vector, material Material) *Plane {
	return &Plane{center,
		normal,
		material,
	}
}

func (p *Plane) BoundingBox() Box {
	return Box{Vector{-INF, -INF, -INF}, Vector{INF, INF, INF}}
}

func (p *Plane) Intersect(r Ray) Hit {
	d := r.Direction.Dot(p.Normal)
	if math.Abs(d) < EPS { // d > 0? TODO
		return NoHit
	}
	a := p.Point.Sub(r.Origin)
	t := a.Dot(p.Normal) / d
	if t < EPS {
		return NoHit
	}
	return Hit{p, t, &r, nil}

}
func (p *Plane) NormalAt(vec Vector) Vector {
	return p.Normal
}
func (p *Plane) MaterialAt(vec Vector) Material {
	return p.Material
}
