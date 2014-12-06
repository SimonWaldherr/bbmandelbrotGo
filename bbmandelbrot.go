package bbmandelbrot

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"sync/atomic"
	"time"
)

const (
	maxiteration = 192
)

var (
	todo uint64
	done uint64
	zh   float64
	zv   float64
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

func Mandelbrot(width, height, cx1, cx2, cy1, cy2 uint64, csr, csg, csb int) (*image.RGBA, string) {
	background := image.Rect(0, 0, int(cx2-cx1), int(cy2-cy1))
	img := image.NewRGBA(background)

	todo = uint64(cx2 - cx1)
	done = 0
	zh = 2.4
	zv = 2.4

	for x := cx1; x < cx2; x++ {
		go func(width int, x int) {
			for y := cy1; y < cy2; y++ {
				xf := float64(x)/float64(width)*zv - (zv/2.0 + 0.5)
				yf := float64(y)/float64(height)*zh - (zh / 2.0)
				c := complex(xf, yf)
				calcval := int(mandel(c) * 255)
				colval := color.RGBA{
					uint8(int(csr) * calcval % 255),
					uint8(int(csg) * calcval % 255),
					uint8(int(csb) * calcval % 255),
					255,
				}
				img.Set(int(x)-int(cx1), int(y)-int(cy1), colval)
			}
			atomic.AddUint64(&done, 1)
		}(int(width), int(x))
	}

	var retstr string
	for todo > done {
		retstr = fmt.Sprintf("\033[2Jcalculated %v%v of Mandelbrot set", int(100/float64(todo)*float64(done)), "%")
		time.Sleep(time.Millisecond * 10)
	}

	return img, retstr
}
