package _d

type Shape interface {
	Perimeter() float64         // 周长
	Area() float64              // 面积
	ContainPoint(p *Point) bool // 包含点
}
