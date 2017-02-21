package main

import . "pt"

var materials = []Material{
	GlossyMaterial(HexColor(0x730046), 1.4, Radians(1)),
	GlossyMaterial(HexColor(0xBFBB11), 1.4, Radians(2)),
	GlossyMaterial(HexColor(0xFFC200), 1.4, Radians(3)),
	GlossyMaterial(HexColor(0xE88801), 1.4, Radians(4)),
	GlossyMaterial(HexColor(0xC93C00), 1.4, Radians(5)),
}

func sphere(scene *Scene, previous, center Vector, radius float64, depth int) {
	if depth <= 0 {
		return
	}
	material := materials[(depth+5)%len(materials)]
	scene.Add(NewSphere(center, radius, material))
	r2 := radius / 2.5
	offset := radius + r2
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				n := 0
				if dx != 0 {
					n++
				}
				if dy != 0 {
					n++
				}
				if dz != 0 {
					n++
				}
				if n != 1 {
					continue
				}
				d := Vector{float64(dx), float64(dy), float64(dz)}
				if d == previous.Negate() {
					continue
				}
				c2 := center.Add(d.MulScalar(offset))
				sphere(scene, d, c2, r2, depth-1)
			}
		}
	}
}

func RenderSpheres() {
	scene := Scene{}
	scene.Color = HexColor(0xFFFFFF)
	sphere(&scene, Vector{}, Vector{}, 1, 6)
	scene.Add(NewSphere(Vector{0, 0, 6}, 0.5, LightMaterial(White, 1)))
	scene.Compile()
	
	camera := LookAt(Vector{3.8, 1.75, 1}, Vector{1, 1, 0}, Vector{0, 0, 1}, 45)
	sampler := Sampler{4, 4}
	renderer := NewRenderer(&scene, &camera, &sampler, 16, 1000, 1000)
	renderer.Render("spheres.png", ColorChannel)
}
