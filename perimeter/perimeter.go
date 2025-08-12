package Perimeter

type Rectangle struct {
    Width float64
    Height float64
}

func Perimeter(width,height float64)float64{
	result := width*2+height*2
	return result
}

func Area(width,height float64)float64{
	result := width*height
	return result
}
func AreaStruct(rectangle Rectangle) float64 {
    return rectangle.Width * rectangle.Height
}