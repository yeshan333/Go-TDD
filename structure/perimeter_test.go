package structure

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{1.0}
		checkPerimeter(t, circle, 6.28)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{1.0}
		checkArea(t, circle, 3.14)
	})
}
