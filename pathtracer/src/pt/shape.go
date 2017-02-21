package pt

import ()

type Shape interface {
	BoundingBox() Box
	Intersect(Ray) Hit
	NormalAt(Vector) Vector
	MaterialAt(Vector) Material
}
