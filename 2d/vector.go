package _d

import "math"

/**
 * 向量
 */
type Vector struct {
	From *Point
	To   *Point
}

/**
 * 向量长度
 */
func (v *Vector) Length() float64 {
	return math.Sqrt(math.Pow(float64(v.To.X-v.From.X), 2) + math.Pow(float64(v.To.Y-v.From.Y), 2))
}

