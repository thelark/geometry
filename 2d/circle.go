package _d

import "math"

/**
 * Circle 圆
 */
type Circle struct {
	R      float64 // 半径
	Center *Point  // 圆心
}

/**
 * Perimeter 计算周长 C = 2 * π * R
 */
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

/**
 * Area 计算面积 S = π * R ^ 2
 */
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

/**
 * ContainPoint 圆内包含某点
 */
func (c *Circle) ContainPoint(p *Point) bool {
	return math.Sqrt(math.Pow(p.X-c.Center.X, 2)+math.Pow(p.Y-c.Center.Y, 2)) < c.R
}
