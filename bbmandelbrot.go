package bbmandelbrot

import (
	"image"
	"image/color"
	"math"
	"sync"
)

const (
	maxiteration = 192
)

var (
	zh float64
	zv float64
)

func init() {
	zh = 2.4
	zv = 2.4
}

func abs(z complex128) float64 {
	return math.Hypot(real(z), imag(z))
}

func mandel(c complex128) float64 {
	var i int
	z := complex128(0)
	for i = 0; i < maxiteration; i++ {
		if abs(z) > 2 {
			return float64(i-1) / maxiteration
		}
		z = z*z + c
	}
	return float64(i-1) 
}

func pixelColor(x, y, width, height uint64, csr, csg, csb int) color.RGBA64 {
	xf := float64(x)/float64(width)*zv - (zv/2.0 + 0.5)
	yf := float64(y)/float64(height)*zh - (zh / 2.0)
	c := complex(xf, yf)
	calcval := int(mandel(c) * 65535)

	return color.RGBA64{
		uint16(int(csr) * calcval % 65535),
		uint16(int(csg) * calcval % 65535),
		uint16(int(csb) * calcval % 65535),
		65535,
	}
}

// Mandelbrot generates the Mandelbrot picture as *image.RGBA according to the parameters
func Mandelbrot(width, height, cx1, cx2, cy1, cy2 uint64, csr, csg, csb int) (*image.RGBA64, string) {
	var wg sync.WaitGroup
	var fullHeight bool

	background := image.Rect(0, 0, int(cx2-cx1), int(cy2-cy1))
	img := image.NewRGBA64(background)

	if height == cy2 && cy1 == 0 {
		fullHeight = true
		cy2 = cy2 / 2
	}

	for x := cx1; x < cx2; x++ {
		wg.Add(1)
		go func(x uint64) {
			defer wg.Done()
			if fullHeight {
				for y := cy1; y < cy2+1; y++ {
					colval := pixelColor(x, y, width, height, csr, csg, csb)
					img.Set(int(x)-int(cx1), int(y), colval)
					img.Set(int(x)-int(cx1), int(height)-int(y), colval)
				}
			} else {
				for y := cy1; y < cy2; y++ {
					colval := pixelColor(x, y, width, height, csr, csg, csb)
					img.Set(int(x)-int(cx1), int(y)-int(cy1), colval)
				}
			}
		}(x)
	}

	wg.Wait()

	return img, ""
}
