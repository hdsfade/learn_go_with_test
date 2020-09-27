package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	area := r.Width * r.Height
	return area
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	area := math.Pi * c.Radius * c.Radius
	return area
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = 2 * (rectangle.Width + rectangle.Height)
	return perimeter
}

