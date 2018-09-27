package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 圆形
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

// 三角开形
type Triangle struct {
	Base   float64
	Height float64
}

func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}
