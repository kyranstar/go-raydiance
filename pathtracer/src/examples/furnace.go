package main

import (
	. "pt"
)

func RenderFurnace() {
	// create a scene
	scene := Scene{}
	scene.Color = White

	// add the ball (a sphere)
	sphere := NewSphere(V(0, 0, 1), 1, DiffuseMaterial(Color{.18,.18,.18}))
	scene.Add(sphere)

	scene.Compile()
	
	// position the camera
	camera := LookAt(V(3, 3, 3), V(0, 0, 0.5), V(0, 0, 1), 50)

	// render the scene
	sampler := Sampler{4, 4, true}
	renderer := NewRenderer(&scene, &camera, &sampler, 8, 1000, 1000)
	renderer.Render("furnace.png", ColorChannel)
}
