package pt

import (
	"math"
	"math/rand"
)

type Sampler struct {
	MaxBounces      int
	FirstHitSamples int
	SampleAllLights bool
}

type BounceType uint8

const (
	BounceTypeAny BounceType = iota
	BounceTypeDiffuse
	BounceTypeSpecular
)

func (s *Sampler) Sample(scene *Scene, ray Ray, rnd *rand.Rand) Color {
	return s.sample(scene, ray, true, s.FirstHitSamples, 0, rnd)
}

func (s *Sampler) sample(scene *Scene, ray Ray, emission bool, samples, depth int, rnd *rand.Rand) Color {
	if depth > s.MaxBounces {
		return Black
	}
	hit := scene.Intersect(ray)
	if !hit.Ok() {
		return scene.Color
	}
	hitInfo := hit.HitInfo()
	material := hitInfo.Material
	result := Black
	if material.Emittance > 0 {
		if !emission {
			return Black
		}
		result = result.Add(material.Color.MulScalar(material.Emittance * float64(samples)))
	}
	rootSpp := int(math.Sqrt(float64(samples)))

	// sample both diffuse and specular on first bounce to avoid noise
	ma, mb := BounceTypeAny, BounceTypeAny
	if depth == 0 {
		ma, mb = BounceTypeDiffuse, BounceTypeSpecular
	}
	for u := 0; u < rootSpp; u++ {
		for v := 0; v < rootSpp; v++ {
			for mode := ma; mode <= mb; mode++ {
				fu := (float64(u) + rnd.Float64()) / float64(rootSpp)
				fv := (float64(v) + rnd.Float64()) / float64(rootSpp)
				newRay, reflected, p := ray.Bounce(hitInfo, fu, fv, mode, rnd)
				if mode == BounceTypeAny {
					p = 1
				}

				if p > EPS && reflected {
					// specular
					indirect := s.sample(scene, newRay, reflected, 1, depth+1, rnd)
					tinted := indirect.Mix(material.Color.Mul(indirect), material.Tint)
					result = result.Add(tinted.MulScalar(p))
				}
				if p > EPS && !reflected {
					// diffuse
					indirect := s.sample(scene, newRay, reflected, 1, depth+1, rnd)
					direct := s.sampleLights(scene, hitInfo.Ray, rnd)
					light := direct.Add(indirect)
					result = result.Add(material.Color.Mul(light).MulScalar(p))
				}
			}
		}
	}
	return result.DivScalar(float64(rootSpp * rootSpp))
}
func (s *Sampler) sampleLights(scene *Scene, n Ray, rnd *rand.Rand) Color {
	nLights := len(scene.Lights)
	if nLights == 0 {
		return Black
	}
	if s.SampleAllLights {
		c := Black
		for _, light := range scene.Lights {
			c = c.Add(s.sampleLight(scene, n, rnd, light))
		}
		return c
	} else {
		// pick a random light
		light := scene.Lights[rand.Intn(nLights)]
		return s.sampleLight(scene, n, rnd, light).MulScalar(float64(nLights))
	}
}

func (s *Sampler) sampleLight(scene *Scene, n Ray, rnd *rand.Rand, light Shape) Color {
	// get bounding sphere center and radius
	var center Vector
	var radius float64
	switch t := light.(type) {
	case *Sphere:
		radius = t.Radius
		center = t.Center
	default:
		// get bounding sphere from bounding box
		box := t.BoundingBox()
		radius = box.OuterRadius()
		center = box.Center()
	}

	// get random point in disk
	point := center
	for {
		x := rnd.Float64()*2 - 1
		y := rnd.Float64()*2 - 1
		if x*x+y*y <= 1 {
			l := center.Sub(n.Origin).Normalize()
			u := l.Cross(RandomUnitVector(rnd)).Normalize()
			v := l.Cross(u)
			point = Vector{}
			point = point.Add(u.MulScalar(x * radius))
			point = point.Add(v.MulScalar(y * radius))
			point = point.Add(center)
			break
		}
	}

	// construct ray toward light point
	ray := Ray{n.Origin, point.Sub(n.Origin).Normalize()}

	// get cosine term
	diffuse := ray.Direction.Dot(n.Direction)
	if diffuse <= 0 {
		return Black
	}

	// check for light visibility
	hit := scene.Intersect(ray)
	if !hit.Ok() || hit.Shape != light {
		return Black
	}

	// compute solid angle (hemisphere coverage)
	hyp := center.Sub(n.Origin).Length()
	opp := radius
	theta := math.Asin(opp / hyp)
	adj := opp / math.Tan(theta)
	d := math.Cos(theta) * adj
	r := math.Sin(theta) * adj
	coverage := (r * r) / (d * d)

	// TODO: fix issue where hyp < opp (point inside sphere)
	if hyp < opp {
		coverage = 1
	}
	coverage = math.Min(coverage, 1)

	// get material properties from light
	material := light.MaterialAt(point)

	// combine factors
	m := material.Emittance * diffuse * coverage
	return material.Color.MulScalar(m)
}
