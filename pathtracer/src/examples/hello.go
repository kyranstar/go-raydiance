package main

import (
	. "pt"
)

func RenderHello() {
	// create a scene
	scene := Scene{}

	// create a material
	material := DiffuseMaterial(White)

	// add the floor (a plane)
	plane := NewPlane(V(0, 0, 0), V(0, 0, 1), material)
	scene.Add(plane)

	// add the ball (a sphere)
	sphere := NewSphere(V(0, 0, 1), 1, SpecularMaterial(White, 1))
	scene.Add(sphere)

	// add a spherical light source
	light := NewSphere(V(0, 0, 5), 1, LightMaterial(White, 8))
	scene.Add(light)
	scene.Compile()
	
	// position the camera
	camera := LookAt(V(3, 3, 3), V(0, 0, 0.5), V(0, 0, 1), 50)

	// render the scene
	sampler := Sampler{4, 4}
	renderer := NewRenderer(&scene, &camera, &sampler, 8, 1000, 1000)
	renderer.Render("hello.png", ColorChannel)
}
