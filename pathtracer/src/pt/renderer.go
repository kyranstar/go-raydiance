package pt

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

type Renderer struct {
	scene           *Scene
	camera          *Camera
	sampler         *Sampler
	buffer          *Buffer
	SamplesPerPixel int
	NumCPU          int
}


func NewRenderer(scene *Scene, camera *Camera, sampler *Sampler, SPP, w, h int) *Renderer {
	return &Renderer{
		scene,
		camera,
		sampler,
		NewBuffer(w, h),
		SPP,
		runtime.NumCPU(),
	}
}

func (r *Renderer) run() {
	
	buf := r.buffer
	runtime.GOMAXPROCS(r.NumCPU)
	sppRoot := int(math.Sqrt(float64(r.SamplesPerPixel)))
	ch := make(chan int, buf.H)
	fmt.Printf("%d x %d pixels, %d spp, %d cores\n", buf.W, buf.H, r.SamplesPerPixel, r.NumCPU)
	start := time.Now()

	for i := 0; i < r.NumCPU; i++ {
		go func(i int) {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for y := i; y < buf.H; y += r.NumCPU {
				for x := 0; x < buf.W; x++ {
					// stratified subsampling
					for u := 0; u < sppRoot; u++ {
						for v := 0; v < sppRoot; v++ {
							fu := (float64(u) + 0.5) / float64(sppRoot)
							fv := (float64(v) + 0.5) / float64(sppRoot)
							ray := r.camera.CastRay(x, y, buf.W, buf.H, fu, fv, rnd)
							sample := r.sampler.Sample(r.scene, ray, rnd)
							buf.AddSample(x, y, sample)
						}
					}
				}
				// finished a row
				ch <- 1
			}
		}(i)
	}
	showProgress(start, r.scene.RayCount(), 0, buf.H)
	for i := 0; i < buf.H; i++ {
		// consume row
		<-ch
		showProgress(start, r.scene.RayCount(), i+1, buf.H)
	}

}
func (r *Renderer) Render(path string, channel Channel) {
	r.run()
	im := r.buffer.Image(channel)
	if err := SavePNG(path, im); err != nil {
		panic(err)
	}
}

func showProgress(start time.Time, rays uint64, i, h int) {
	pct := int(100 * float64(i) / float64(h))
	elapsed := time.Since(start)
	rps := float64(rays) / elapsed.Seconds()
	fmt.Printf("\r%4d / %d (%3d%%) [", i, h, pct)
	for p := 0; p < 100; p += 3 {
		if pct > p {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	timeUntilDone := int(float64(elapsed.Nanoseconds()) /  float64(i) * (float64(h - i)));
	fmt.Printf("] %s elapsed, %s rays/s, %s until done", DurationString(elapsed), NumberString(rps), DurationString(time.Duration(timeUntilDone)))
}
