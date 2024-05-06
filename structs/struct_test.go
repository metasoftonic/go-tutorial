package structs

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("Expected %g, got %g", want, got)
		}
	}
	t.Run("Testing area of rectange", func(t *testing.T) {
		rect := Rectangle{10, 10}
		checkArea(t, rect, 40)
	})
}

func TestAreaGroup(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, hasArea: 40},
		{name: "Circle", shape: Circle{4}, hasArea: 50.26548245743669},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, sh := range areaTest {
		got := sh.shape.Area()
		if got != sh.hasArea {
			t.Errorf("Shape %v, expected %g, got %g", sh.name, sh.hasArea, got)
		}
	}
}
