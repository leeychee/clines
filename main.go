package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

type point struct{ x, y float64 }
type line struct{ p1, p2 point }

// how many segment on the circle
var segment int
var lfile string

var r, r25, r50 float64
var fontsize float64
var fontpath string
var imgfn string

// contain points on the circle
var ps, ts []point

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "用来生成包含等分连线的原先图片。")
		flag.PrintDefaults()
	}
	flag.IntVar(&segment, "s", 10, "将圆按指定数字等分")
	flag.StringVar(&lfile, "f", "", "定义连线的文件路径，当此参数为空时，将从命令行参数读取")
	flag.Float64Var(&r, "r", 500, "圆半径，间接指定了图片大小")
	flag.Float64Var(&fontsize, "z", 20, "字体大小")
	flag.StringVar(&fontpath, "p", "./luxisr.ttf", "字体路径")
	flag.StringVar(&imgfn, "o", "output", "输出图片路径")
	flag.Parse()
	r25 = r + fontsize
	r50 = r + 50
	ps = make([]point, segment+1, segment+1)
	for i := 0; i < len(ps)-1; i++ {
		x := r50 + r*math.Sin(2*math.Pi/float64(segment)*float64(i))
		y := r50 - r*math.Cos(2*math.Pi/float64(segment)*float64(i))
		ps[i+1] = point{x, y}
	}
	ts = make([]point, segment+1, segment+1)
	for i := 0; i < len(ts)-1; i++ {
		x := r50 + r25*math.Sin(2*math.Pi/float64(segment)*float64(i))
		y := r50 - r25*math.Cos(2*math.Pi/float64(segment)*float64(i))
		ts[i+1] = point{x, y}
	}
}

func main() {
	if lfile == "" {
		// get lines from cli
		lines := parse(flag.Args())
		draw(imgfn, lines)
	} else {
		file, err := os.Open(lfile)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			imgstr := strings.Fields(scanner.Text())
			if len(imgstr) > 2 {
				fn := imgstr[0]
				lines := parse(imgstr[1:])
				draw(fn, lines)
			}
		}
	}
}

func parse(linestr []string) []line {
	lines := make([]line, 10)
	for _, s := range linestr {
		ss := strings.Split(s, ",")
		p1, _ := strconv.Atoi(ss[0])
		p2, _ := strconv.Atoi(ss[1])
		if p1 > segment || p2 > segment {
			log.Fatalf("line error: (%d, %d) overlap %d\n", p1, p2, segment)
		}
		lines = append(lines, line{ps[p1], ps[p2]})
	}
	return lines
}

func draw(fn string, lines []line) {
	dc := gg.NewContext(int(2*r50), int(2*r50))
	if err := dc.LoadFontFace(fontpath, fontsize); err != nil {
		log.Printf("Fail to load font: %s", err)
	}
	// Draw white backgroud
	dc.DrawRectangle(0, 0, 2*r50, 2*r50)
	dc.SetRGBA(1, 1, 1, 1)
	dc.Fill()
	// Draw black circle and lines
	dc.SetRGBA(0, 0, 0, 1)
	dc.DrawCircle(r50, r50, r)
	// Draw point text
	for i := 1; i < len(ts); i++ {
		dc.DrawStringAnchored(strconv.Itoa(i), ts[i].x, ts[i].y, 0.5, 0.5)
	}
	for _, l := range lines {
		dc.DrawLine(l.p1.x, l.p1.y, l.p2.x, l.p2.y)
		//dc.SetLineWidth(5)
		dc.Stroke()
	}
	if err := dc.SavePNG(fn + ".png"); err != nil {
		log.Printf("Fail to save PNG: %s", err)
	}
}
