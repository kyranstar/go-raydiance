package main

import (
	. "pt"
)

func RenderTransparent() {
	// create a scene
	scene := Scene{}

	// add the floor (a plane)
	plane := NewPlane(V(0, 0, 0), V(0, 0, 1), DiffuseMaterial(White))
	scene.Add(plane)

	// add the ball (a sphere)
	sphere := NewSphere(V(0, 0, 1), 1,  TransparentMaterial(Color{0,0,1}, 1.1, 0.05, 0.5))
	scene.Add(sphere)
	sphere2 := NewSphere(V(-1.5, -1.5, 1), 0.5,  DiffuseMaterial(Color{1, 0, 0}))
	scene.Add(sphere2)
	sphere3 := NewSphere(V(-0.5, -1.5, 1), 0.5,  DiffuseMaterial(Color{1, 1, 0}))
	scene.Add(sphere3)
	sphere4 := NewSphere(V(-1.5, -0.5, 1), 0.5,  DiffuseMaterial(Color{1, 0, 1}))
	scene.Add(sphere4)

	// add a spherical light source
	light := NewSphere(V(0, 0, 5), 1, LightMaterial(White, 8))
	scene.Add(light)
	scene.Compile()
	
	// position the camera
	camera := LookAt(V(3, 3, 3), V(0, 0, 0.5), V(0, 0, 1), 50)

	// render the scene
	sampler := Sampler{3, 16, true}
	renderer := NewRenderer(&scene, &camera, &sampler, 32, 500, 500)
	renderer.Render("transparent.png", ColorChannel)
}
