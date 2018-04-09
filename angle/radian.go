package angle

import (
	"math"
	"fmt"
)

/**
 * 定义弧度
 */
type Radian float64

const (
	RightRadian Radian = math.Pi / 2 //
	FlatRadian  Radian = math.Pi
	Fillet      Radian = math.Pi * 2
)

/**
 * 弧度转度
 */
func (r Radian) ToDegree() Angle {
	return Angle(r * 180 / math.Pi)
}

func (r Radian) ToString() string {
	return fmt.Sprintf("%frad", r)
}
