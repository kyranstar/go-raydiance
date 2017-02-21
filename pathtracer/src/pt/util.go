package pt

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

const INF = 1e9
const EPS = 1e-9

func DurationString(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d", h, m, s)
}

func NumberString(x float64) string {
	suffixes := []string{"", "k", "M", "G"}
	for _, suffix := range suffixes {
		if x < 1000 {
			return fmt.Sprintf("%.1f%s", x, suffix)
		}
		x /= 1000
	}
	return fmt.Sprintf("%.1f%s", x, "T")
}
func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
func SavePNG(path string, im image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, im)
}
func Median(items []float64) float64 {
	n := len(items)
	switch {
	case n == 0:
		return 0
	case n%2 == 1:
		return items[n/2]
	default:
		a := items[n/2-1]
		b := items[n/2]
		return (a + b) / 2
	}
}

func Cone(direction Vector, theta, u, v float64, rnd *rand.Rand) Vector {
	if theta < EPS {
		return direction
	}
	theta = theta * (1 - (2 * math.Acos(u) / math.Pi))
	m1 := math.Sin(theta)
	m2 := math.Cos(theta)
	a := v * 2 * math.Pi
	q := RandomUnitVector(rnd)
	s := direction.Cross(q)
	t := direction.Cross(s)
	d := Vector{}
	d = d.Add(s.MulScalar(m1 * math.Cos(a)))
	d = d.Add(t.MulScalar(m1 * math.Sin(a)))
	d = d.Add(direction.MulScalar(m2))
	d = d.Normalize()
	return d
}

func RandomUnitVector(rnd *rand.Rand) Vector {
	for {
		var x, y, z float64
		if rnd == nil {
			x = rand.Float64()*2 - 1
			y = rand.Float64()*2 - 1
			z = rand.Float64()*2 - 1
		} else {
			x = rnd.Float64()*2 - 1
			y = rnd.Float64()*2 - 1
			z = rnd.Float64()*2 - 1
		}
		if x*x+y*y+z*z > 1 {
			continue
		}
		return Vector{x, y, z}.Normalize()
	}
}
