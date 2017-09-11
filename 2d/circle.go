package _d

import "math"

/**
 * 圆
 */
type Circle struct {
	R      float64 // 半径
	Center *Point  // 圆心
}

/**
 * 计算周长 C = 2 * π * R
 */
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

/**
 * 计算面积 S = π * R ^ 2
 */
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2)
}
