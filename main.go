package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"flag"
	"bufio"
	"os"

	"github.com/fogleman/gg"
)

type point struct{ x, y float64 }
type line struct{p1, p2 point}

var r float64 = 500
var ps []point
var out string
var f string

func init() {
	flag.StringVar(&out, "o", "out", "output imange name")
	flag.StringVar(&f, "f", "", "lines definition file")
	ps = make([]point, 13)
	ps[0] = point{0, 0}
	for i := 1; i <= 12; i++ {
		ps[i] = point{r+50+r * math.Sin(math.Pi/6*float64(i)), r+50-r * math.Cos(math.Pi/6*float64(i))}
	}
}

func main() {
	// get lines from cli
	flag.Parse()
	if f!="" {
		file, err := os.Open(f)
		if  err!=nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			imgstr := strings.Fields(scanner.Text())
			draw(imgstr[0], imgstr[1:])
		}
	} else {
		draw(out, flag.Args())
	}
}

func draw(fn string, linestr []string) {
	lines:=make([]line, 10)
	for _, s := range linestr {
		ss := strings.Split(s, ",")
		p1,_ := strconv.Atoi(ss[0])
		p2,_ := strconv.Atoi(ss[1])
		lines=append(lines, line{ps[p1], ps[p2]})
	}
	dc := gg.NewContext(1100, 1100)
	// Draw white backgroud
	dc.DrawRectangle(0,0,1100,1100)
	dc.SetRGBA(1,1,1, 1)
	dc.Fill()
	// Draw black circle and lines
	dc.SetRGBA(0, 0, 0, 1)
	dc.DrawCircle(550, 550, 500)
	for _, l := range lines {
		dc.DrawLine(l.p1.x, l.p1.y, l.p2.x, l.p2.y)
		//dc.SetLineWidth(5)
		dc.Stroke()
	}
	dc.SavePNG(fn+".png")
}
