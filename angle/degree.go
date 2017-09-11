package angle

import "math"

/**
 * 定义 度
 */
type Degree float64

/**
 * 度转弧度
 */
func (d Degree) ToRadian() Radian {
	return Radian(d * math.Pi / 180)
}
