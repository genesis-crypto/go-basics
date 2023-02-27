package main

import "math"

type Point struct {
	x float64
	y float64
}

func sides(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Range(a, b Point) float64 {
	cx, cy := sides(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
