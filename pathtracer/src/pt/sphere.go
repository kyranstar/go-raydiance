package pt

import (
	"math"
)

type Sphere struct {
	Center   Vector
	Radius   float64
	Material Material
	Box      Box
}

func NewSphere(center Vector, radius float64, material Material) *Sphere {
	return &Sphere{center,
		radius,
		material,
		Box{center.SubScalar(radius), center.AddScalar(radius)},
	}
}

func (s *Sphere) BoundingBox() Box {
	return s.Box
}

func (s *Sphere) Intersect(r Ray) Hit {
	to := r.Origin.Sub(s.Center)
	b := to.Dot(r.Direction)
	c := to.Dot(to) - s.Radius*s.Radius
	d := b*b - c
	if d > 0 {
		d = math.Sqrt(d)
		t1 := -b - d
		if t1 > EPS {
			return Hit{s, t1, &r, nil}
		}
		t2 := -b + d
		if t2 > EPS {
			return Hit{s, t2, &r, nil}
		}
	}
	return NoHit
}
func (s *Sphere) NormalAt(vec Vector) Vector {
	return vec.Sub(s.Center).Normalize()
}
func (s *Sphere) MaterialAt(vec Vector) Material {
	return s.Material
}
