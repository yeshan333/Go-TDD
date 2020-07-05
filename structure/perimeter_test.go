package structure

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{1.0}
		got := circle.Perimeter()
		want := 6.28

		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{1.0}
		got := circle.Area()
		want := 3.14

		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}
