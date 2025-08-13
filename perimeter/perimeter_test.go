package Perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	t.Run("test number", func(t *testing.T) {
		getWidth := 12
		getHeight := 24
		get := Perimeter(float64(getWidth), float64(getHeight))
		want := 72.0
		if float64(get) != float64(want) {
			t.Errorf("got '%f' want '%f'", get, want)
		}
	})

	t.Run("test Area", func(t *testing.T) {
		width := 12
		height := 2
		get := Area(float64(width), float64(height))
		want := 24.0
		if get != float64(want) {
			t.Errorf("got '%f' want '%f'", get, want)
		}
	})

	t.Run("test Area struct", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		get := rectangle.AreaStruct()
		want := 72.0
		if get != float64(want) {
			t.Errorf("got '%f' want '%f'", get, want)
		}
	})

	t.Run("circle Area", func(t *testing.T) {
		circle := Circle{10}
		get := circle.AreaStruct()
		want := 314.16
		if get != float64(want) {
			t.Errorf("got '%f' want '%f'", get, want)
		}
	})

	t.Run("shape Area", func(t *testing.T) {
		testList := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.16},
		}

		for _, testArea := range testList {
			get := testArea.shape.AreaStruct()
			want := testArea.want
			if get != float64(want) {
				t.Errorf("got '%f' want '%f'", get, want)
			}
		}
	})

	t.Run("shape Area", func(t *testing.T) {
		areaTests := []struct {
			name    string
			shape   Shape
			hasArea float64
		}{
			{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
			{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.16},
		}

		for _, tt := range areaTests {
			// using tt.name from the case to use it as the `t.Run` test name
			t.Run(tt.name, func(t *testing.T) {
				got := tt.shape.AreaStruct()
				if got != tt.hasArea {
					t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
				}
			})

		}
	})
}
