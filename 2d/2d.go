package _d

type Shape interface {
	Perimeter() float64         // 周长
	Area() float64              // 面积
	ContainPoint(p *Point) bool // 包含点
}

// 叉积
func multiply(p1, p2, p0 *Point) float64 {
	return (p1.X-p0.X)*(p2.Y-p0.Y) - (p2.X-p0.X)*(p1.Y-p0.Y)
}
