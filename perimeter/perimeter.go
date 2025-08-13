package Perimeter

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	AreaStruct() float64
}

func Perimeter(width, height float64) float64 {
	result := width*2 + height*2
	return result
}

func Area(width, height float64) float64 {
	result := width * height
	return result
}
func (rectangle Rectangle) AreaStruct() float64 {
	return rectangle.Width * rectangle.Height
}

func (c Circle) AreaStruct() float64 {
	return 3.1416 * c.Radius * c.Radius
}
