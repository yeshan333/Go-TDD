package structure

// Shape : 图形，多态，只要实现了这两个方法的类都实现了 Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

/* ---------------------------------------------------- */

// Rectangle : 矩形
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter : 计算矩形周长
func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area : 计算矩形面积
func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

/* ---------------------------------------------------- */

// Circle : 圆形
type Circle struct {
	Redius float64
}

// Perimeter : 计算⚪周长
func (circle Circle) Perimeter() float64 {
	return 2 * 3.14 * circle.Redius
}

// Area : 计算⚪面积
func (circle Circle) Area() float64 {
	return 3.14 * circle.Redius * circle.Redius
}
