package angle

import (
	"math"
	"fmt"
)

/**
 * 定义 度
 */
type Angle float64

const (
	RightAngle Angle = 90
	StraightAngle Angle = 180
)

/**
 * 度转弧度
 */
func (α Angle) ToRadian() Radian {
	return Radian(α * math.Pi / 180)
}

/**
 * 夹角
 */
func (α Angle) Sub(β Angle) Angle {
	return α - β
}

/**
 * 合角
 */
func (α Angle) Add(β Angle) float64 {
	return float64(α + β)
}

/**
 * 补角
 */
func (α Angle) Supplementary() Angle {
	return StraightAngle - α
}

/**
 * 余角
 */
func (α Angle) Complementary() Angle {
	return RightAngle - α
}

func (α Angle) ToString() string {

	return fmt.Sprintf("%f°", α)
}
