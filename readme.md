# bbmandelbrotGo

[![Build Status](https://travis-ci.org/SimonWaldherr/bbmandelbrotGo.svg?branch=master)](https://travis-ci.org/SimonWaldherr/bbmandelbrotGo) 
[![Coverage Status](https://coveralls.io/repos/SimonWaldherr/bbmandelbrotGo/badge.png)](https://coveralls.io/r/SimonWaldherr/bbmandelbrotGo) 
[![Go Report Card](https://goreportcard.com/badge/github.com/simonwaldherr/bbmandelbrotGo)](https://goreportcard.com/report/github.com/simonwaldherr/bbmandelbrotGo) 
[![codebeat badge](https://codebeat.co/badges/f99a42ee-8dae-4cb9-8c88-c27d20b79edd)](https://codebeat.co/projects/github-com-simonwaldherr-bbmandelbrotgo-master) 
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/SimonWaldherr/bbmandelbrotGo) 
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/SimonWaldherr/bbmandelbrotGo/master/LICENSE) 

## what

in memoriam [Benoît B. Mandelbrot](http://en.wikipedia.org/wiki/Benoit_Mandelbrot)  
the [Mandelbrot set](http://en.wikipedia.org/wiki/Mandelbrot_set) is a very popular [fractal](http://en.wikipedia.org/wiki/Fractal).
There is a great [Song (and Video) on YouTube about the Mandelbrot(/Julia) fractal](https://www.youtube.com/watch?v=ES-yKOYaXq0). 
Many thanks to [Jonathan Coulton](https://www.jonathancoulton.com) for this great song and for sharing it under [CC](https://creativecommons.org)-BY-NC

![Mandelbrot Fractal](https://raw.githubusercontent.com/SimonWaldherr/bbmandelbrotGo/master/mandelbrot.png)

## how

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
