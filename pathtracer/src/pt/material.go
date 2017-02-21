package pt

import ()

type Material struct {
	Color        Color
	Emittance    float64
	Index        float64 // refractive index
	Gloss        float64 // reflection cone angle in radians
	Tint         float64 // specular and refractive tinting
	Reflectivity float64 // metallic reflection
	Transparent  bool
}

func DiffuseMaterial(color Color) Material {
	return Material{color, 0, 1, 0, 0, -1, false}
}

func SpecularMaterial(color Color, index float64) Material {
	return Material{color, 0, index, 0, 0, -1, false}
}

func GlossyMaterial(color Color, index, gloss float64) Material {
	return Material{color, 0, index, gloss, 0, -1, false}
}

func ClearMaterial(index, gloss float64) Material {
	return Material{Black, 0, index, gloss, 0, -1, true}
}

func TransparentMaterial(color Color, index, gloss, tint float64) Material {
	return Material{color, 0, index, gloss, tint, -1, true}
}

func MetallicMaterial(color Color, gloss, tint float64) Material {
	return Material{color, 0, 1, gloss, tint, 1, false}
}

func LightMaterial(color Color, emittance float64) Material {
	return Material{color, emittance, 1, 0, 0, -1, false}
}
