package pt

import ()

type Hit struct {
	Shape   Shape
	T       float64
	Ray     *Ray
	hitInfo *HitInfo
}

type HitInfo struct {
	Shape    Shape
	Position Vector
	Normal   Vector
	Ray      Ray
	Material Material
	Inside   bool
}

var NoHit = Hit{nil, INF, nil, nil}

func (hit *Hit) Ok() bool {
	return hit.T < INF
}

func (hit *Hit) HitInfo() *HitInfo {
	if hit.hitInfo != nil {
		return hit.hitInfo
	}
	shape := hit.Shape
	position := hit.Ray.Step(hit.T)
	normal := shape.NormalAt(position)
	material := hit.Shape.MaterialAt(position)
	inside := false
	if normal.Dot(hit.Ray.Direction) > 0 {
		normal = normal.Negate()
		inside = true
		//		switch shape.(type) {
		//		case *Volume, *SDFShape, *SphericalHarmonic:
		//			inside = false
		//		}
	}
	ray := Ray{position, normal}
	hit.hitInfo = &HitInfo{shape, position, normal, ray, material, inside}
	return hit.hitInfo
}
