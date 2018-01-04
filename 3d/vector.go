package _d

import "math"

type Vector struct {
	From *Point
	To   *Point
}

/**
 * 向量长度
 */
func (v *Vector) Length() float64 {
	return math.Sqrt(math.Pow(float64(v.To.X-v.From.X), 2) + math.Pow(float64(v.To.Y-v.From.Y), 2) + math.Pow(float64(v.To.Z-v.From.Z), 2))
}

/**
 * 转换为从原点触发的向量
 */
func (v *Vector) ToVector() *Vector {
	return &Vector{
		From: Origin,
		To:   &Point{X: v.To.X - v.From.X, Y: v.To.Y - v.From.Y, Z: v.To.Z - v.From.Z},
	}
}
