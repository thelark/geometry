package angle

import "math"

/**
 * 定义弧度
 */
type Radian float64

/**
 * 弧度转度
 */
func (r Radian) ToDegree() Degree {
	return Degree(r * 180 / math.Pi)
}
