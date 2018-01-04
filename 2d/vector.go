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

/**
 * 向量加
 */
func (v *Vector) Add(t *Vector) *Vector {
	return &Vector{
		From: &Point{X: v.From.X + t.From.X, Y: v.From.Y + t.From.Y},
		To:   &Point{X: v.To.X + t.To.X, Y: v.To.Y + t.To.Y},
	}
}

/**
 * 向量减
 */
func (v *Vector) Sub(t *Vector) *Vector {
	return &Vector{
		From: &Point{X: v.From.X - t.From.X, Y: v.From.Y - t.From.Y},
		To:   &Point{X: v.To.X - t.To.X, Y: v.To.Y - t.To.Y},
	}
}

/**
 * 向量与向量是否相交
 */
func (v *Vector) Intersect(t *Vector) bool {
	return v.ToVector().To.X/t.ToVector().To.X != v.ToVector().To.Y/t.ToVector().To.Y
}

func (v *Vector) ToVector() *Vector {
	return &Vector{
		From: Origin,
		To:   &Point{X: v.To.X - v.From.X, Y: v.To.Y - v.From.Y},
	}
}

/**
 * 叉积
 */
func (v *Vector) Cross() float64 {
	return math.Abs(v.From.X*v.To.Y - v.From.Y*v.To.X)
}
