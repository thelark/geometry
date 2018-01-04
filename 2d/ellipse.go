package _d

import "math"

/**
 * 椭圆
 */
type Ellipse struct {
	RadiusX      float64   // 长轴半径
	RadiusY      float64   // 短轴半径
	Center       *Point    // 正中
	F            [2]*Point // 两焦点
	eccentricity float64   // 离心率
}

func (e *Ellipse) Area() float64 {
	return math.Pi * e.RadiusX * e.RadiusY
}

func (e *Ellipse) Perimeter() float64 {
	q := e.RadiusY + e.RadiusX
	h := math.Pow((e.RadiusX-e.RadiusY)/(e.RadiusX+e.RadiusY), 2)
	m := 22/7*math.Pi - 1
	n := math.Pow((e.RadiusX-e.RadiusY)/e.RadiusX, 33.679)
	return math.Pi * q * (1 + 3*h/(10+math.Sqrt(4-3*h))*(1+m*n))
}

/**
 * 椭圆内包含某点
 */
func (e *Ellipse) ContainPoint(p *Point) bool {
	e.eccentricity = math.Sqrt(1 - math.Pow(e.RadiusY, 2)/math.Pow(e.RadiusX, 2))
	if p.X*math.Cos(e.eccentricity)+p.Y*math.Sin(e.eccentricity) == 1 {
		// 在圆上
		return true
	} else if p.X*math.Cos(e.eccentricity)+p.Y*math.Sin(e.eccentricity) < 1 {
		// 在圆内
		return true
	} else {
		// 在圆外
		return false
	}
}
