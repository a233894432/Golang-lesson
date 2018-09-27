package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		got := Area(rectangle)
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := 314.16

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})

// }
// func TestArea(t *testing.T) {

// 	areaTests := []struct {
// 		shape Shape
// 		want  float64
// 	}{
// 		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
// 		{shape: Circle{Radius: 10}, want: 314.1592653589793},
// 		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
// 	}

// 	for _, tt := range areaTests {
// 		got := tt.shape.Area()
// 		if got != tt.want {
// 			t.Errorf("got %.2f want %.2f", got, tt.want)
// 		}
// 	}

// }

func TestArea(t *testing.T) {
​
    areaTests := []struct {
        name    string
        shape   Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
        {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
    }
​
    for _, tt := range areaTests {
        // using tt.name from the case to use it as the `t.Run` test name
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
            }
        })
​
    }
​
}