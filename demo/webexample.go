package main

import (
	bbmandelbrot ".."
	"encoding/base64"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func dataURI(fileName, contentType string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("data:%s;base64,%s", contentType, base64.StdEncoding.EncodeToString(data))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		fname  string
		width  uint64
		height uint64
		cx1    uint64
		cx2    uint64
		cy1    uint64
		cy2    uint64
		csr    int
		csg    int
		csb    int
	)

	urls := r.URL.Path
	if r.URL.RawQuery != "" {
		urls = strings.Join([]string{urls, "?", r.URL.RawQuery}, "")
		q, _ := url.ParseQuery(r.URL.RawQuery)
		width, _ = strconv.ParseUint(q["w"][0], 10, 64)
		height, _ = strconv.ParseUint(q["h"][0], 10, 64)
		cx1, _ = strconv.ParseUint(q["cx1"][0], 10, 64)
		cy1, _ = strconv.ParseUint(q["cy1"][0], 10, 64)
	} else {
		width = 640
		height = 640
		cx1 = 0
		cy1 = 0
	}

	csr = 2
	csg = 3
	csb = 1
	cx2 = cx1 + width/4
	cy2 = cy1 + height/4
	fmt.Printf("width: %d height: %d cx1: %d cy1: %d cx2: %d cy2: %d\n", width, height, cx1, cy1, cx2, cy2)
	fmt.Fprintf(w, "<html><head><style>img:hover{opacity:0.8}</style></head><body>")

	var x, y uint64

	for y = 0; y < 4; y++ {
		for x = 0; x < 4; x++ {
			fname = fmt.Sprintf("cache/%vx%v_%v_%v_%v_%v_mandelbrot.png", width, height, cx1+160*x, cx1+160*(x+1), cy1+160*y, cy1+160*(y+1))
			if _, err := os.Stat(fname); err != nil {
				fmt.Println("generating ", fname)
				img, _ := bbmandelbrot.Mandelbrot(width, height, cx1+160*x, cx1+160*(x+1), cy1+160*y, cy1+160*(y+1), csr, csg, csb)

				file, err := os.Create(fname)
				defer file.Close()

				if err != nil || file == nil {
					file, err = os.Open(fname)
					defer file.Close()
					if err != nil {
						fmt.Println(err)
					}
				}

				err = png.Encode(file, img)
				if err != nil {
					fmt.Println(err)
				}
			}

			uri := dataURI(fname, "image/png")
			fmt.Fprintf(w, "<a href=\"?w=%d&h=%d&cx1=%d&cy1=%d\"><img src=\"%v\" /></a>", width*4, height*4, cx1+x*width, cy1+y*height, uri)
		}
		fmt.Fprintf(w, "<br />")
	}

	fmt.Fprintf(w, "</body></html>")
	log.Println(urls)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
