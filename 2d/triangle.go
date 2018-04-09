package _d

import (
	"thelark.cn/geometry.v1/angle"
	"math"
)

/**
 * 三角形
 */
type Triangle struct {
	P [3]*Point
}

/**
 * 获取所有点
 */
func (d *Triangle) Points() []*Point {
	return []*Point{d.P[0], d.P[1], d.P[2]}
}

/**
 * 获取所有边 顺序是从 [0] 开始
 */
func (d *Triangle) Lines() []*Vector {
	lines := make([]*Vector, 0)
	for k, point := range d.Points() {
		i := k + 1
		if k == len(d.P)-1 {
			i = 0
		}
		lines = append(lines, &Vector{point, d.P[i]})
	}
	return lines
}

/**
 * 计算周长
 */
func (d *Triangle) Perimeter() float64 {
	perimeter := 0.0
	for _, line := range d.Lines() {
		perimeter += line.Length()
	}
	return perimeter
}

/**
 * 计算面积
 */
func (d *Triangle) Area() float64 {
	/**
	 * 设有 A、B、C 三点
	 * SΔ = (1 / 2) * (A.X * B.Y + B.X * C.Y + C.X * A.Y - A.X * C.Y - B.X * A.Y - C.X * B.Y)
	 */
	return 1 / 2 * (d.P[0].X*d.P[1].Y + d.P[1].X*d.P[2].Y + d.P[2].X*d.P[0].Y - d.P[0].X*d.P[2].Y - d.P[1].X*d.P[0].Y - d.P[2].X*d.P[1].Y)
}

/**
 * 三角形包含某点
 */
func (d *Triangle) ContainPoint(point *Point) bool {
	j := len(d.Points()) - 1
	containPoint := false
	for index := 0; index < len(d.Points()); index++ {
		if d.Points()[index].Y < point.Y && d.Points()[j].Y >= point.Y ||
			d.Points()[j].Y < point.Y && d.Points()[index].Y >= point.Y {
			if d.Points()[index].X + (point.Y - d.Points()[index].Y) /
				(d.Points()[j].Y - d.Points()[index].Y) *
				(d.Points()[j].X - d.Points()[index].X) < point.X {
				containPoint = !containPoint
			}
		}
		j = index
	}
	return containPoint
}

/**
 * 直角三角形
 * ∠α = ∠CAB
 * ∠β = ∠ACB
 * B 为直角
 */
type RightTriangle struct {
	Triangle
	α, β    angle.Angle // 另一个角是直角
	A, B, C *Point
}

func (r *RightTriangle) Perimeter() float64 {
	AC := math.Sqrt(math.Pow((&Vector{r.A, r.B}).Length(), 2) + math.Pow((&Vector{r.B, r.C}).Length(), 2))
	return (&Vector{r.A, r.B}).Length() + (&Vector{r.B, r.C}).Length() + AC
}
func (r *RightTriangle) Area() float64 {
	return 1 / 2 * (&Vector{r.A, r.B}).Length() * (&Vector{r.B, r.C}).Length()
}

func (r *RightTriangle) Points() []*Point {
	return []*Point{r.A, r.B, r.C}
}


