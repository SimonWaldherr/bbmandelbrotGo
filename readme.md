##bbmandelbrot.go

[![Build Status](https://travis-ci.org/SimonWaldherr/bbmandelbrot.go.svg?branch=master)](https://travis-ci.org/SimonWaldherr/bbmandelbrot.go) 
[![Coverage Status](https://coveralls.io/repos/SimonWaldherr/bbmandelbrot.go/badge.png)](https://coveralls.io/r/SimonWaldherr/bbmandelbrot.go)

in memoriam [Benoît B. Mandelbrot](http://en.wikipedia.org/wiki/Benoit_Mandelbrot)  
the [Mandelbrot set](http://en.wikipedia.org/wiki/Mandelbrot_set) is a very popular [fractal](http://en.wikipedia.org/wiki/Fractal).
There is a great [Song (and Video) on YouTube about the Mandelbrot(/Julia) fractal](https://www.youtube.com/watch?v=ES-yKOYaXq0). 
Many thanks to [Jonathan Coulton](https://www.jonathancoulton.com) for this great song and for sharing it under [CC](https://creativecommons.org)-BY-NC

```
go test .
```

```
go run demo/cliexample.go --help
Usage of ./bbmandelbrot:
  -b=1: color scheme (blue)
  -f="mandelbrot.png": destination filename
  -g=3: color scheme (green)
  -h=2560: fractal height
  -r=2: color scheme (red)
  -w=2560: fractal width
exit status 2
```
