package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 8.0}, area: 80.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, area: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 11, Height: 7}, area: 38},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.area {
			t.Errorf("%#v got %g, want %g", tt.shape, got, tt.area)
		}
	}

}
