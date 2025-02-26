package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	hasArea := 40.0

	if got != hasArea {
		t.Errorf("got %.2f hasArea %.2f", got, hasArea)
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != hasArea {
// 			t.Errorf("got %g hasArea %g", got, hasArea)
// 		}
// 	}
// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 			checkArea(t, rectangle, 72.0)

// 	})

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %g want %g", got, tt.hasArea)
			}
		})
	}
}
