package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	cscheme string
	fname   string
	todo    uint64
	done    uint64
	width   float64
	height  float64
	csr     int
	csg     int
	csb     int
	zh      float64
	zv      float64
)

const (
	maxiteration = 192
)

func mandel(c complex128) float64 {
	z := complex128(0)
	for i := 0; i < maxiteration; i++ {
		if cmplx.Abs(z) > 2 {
			return float64(i-1) / maxiteration
		}
		z = z*z + c
	}
	return 0
}

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.StringVar(&fname, "f", "mandelbrot.png", "destination filename")
	flag.Float64Var(&width, "w", 2560, "fractal width")
	flag.Float64Var(&height, "h", 2560, "fractal height")
	flag.IntVar(&csr, "r", 2, "color scheme (red)")
	flag.IntVar(&csg, "g", 3, "color scheme (green)")
	flag.IntVar(&csb, "b", 1, "color scheme (blue)")
	flag.Parse()
	zh = 2.4
	zv = 2.4

	background := image.Rect(0, 0, int(width), int(height))
	img := image.NewRGBA(background)

	todo = uint64(width)
	done = 0

	for x := 0; x < int(width); x++ {
		go func(width float64, x int) {
			for y := 0; y < int(height); y++ {
				xf := float64(x)/width*zv - (zv/2.0 + 0.5)
				yf := float64(y)/height*zh - (zh / 2.0)
				c := complex(xf, yf)
				calcval := int(mandel(c) * 255)
				colval := color.RGBA{
					uint8(int(csr) * calcval % 255),
					uint8(int(csg) * calcval % 255),
					uint8(int(csb) * calcval % 255),
					255,
				}
				img.Set(x, y, colval)
			}
			atomic.AddUint64(&done, 1)
		}(width, x)
	}

	for todo > done {
		fmt.Printf("\033[2Jcalculated %v%v of Mandelbrot set\n", int(100/float64(todo)*float64(done)), "%")
		time.Sleep(time.Millisecond * 10)
	}

	file, err := os.Create(fname)
	defer file.Close()

	if err != nil || file == nil {
		file, err = os.Open(fname)
		defer file.Close()
		if err != nil {
			panic(fmt.Sprintf("Error opening file: %s\n", err))
		}
	}

	err = png.Encode(file, img)
	if err != nil {
		panic(fmt.Sprintf("Error encoding image: %s\n", err))
	}
	fmt.Printf("\033[2Jimage saved to %v after %v\n", fname, time.Since(start))
}
