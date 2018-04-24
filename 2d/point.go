package _d

// Origin 原点
var Origin = &Point{X: 0, Y: 0}

/**
 * 2D 点
 */
type Point struct {
	X float64 // X
	Y float64 // Y
}

type PointStatus uint8

const (
	IsInLine  PointStatus = iota // 在线上
	IsInLeft                     // 在线左侧
	IsInRight                    // 在线右侧
)

/**
 * IsInVectorStatus 判断点在向量某侧
 */
func (p *Point) IsInVectorStatus(v *Vector) PointStatus {
	r := (v.From.X-p.X)*(v.To.Y-p.Y) - (v.From.Y-p.Y)*(v.To.X-p.X)
	if r > 0 {
		return IsInLeft
	} else if r < 0 {
		return IsInRight
	} else {
		return IsInLine
	}
}

/**
 * IsInShape 判断点是否在形状内部
 */
func (p *Point) IsInShape(s Shape) bool {
	return s.ContainPoint(p)
}
func (p *Point) Add(t *Point) *Point {
	return &Point{X: p.X + t.X, Y: p.Y + t.Y}
}
func (p *Point) Sub(t *Point) *Point {
	return &Point{X: p.X - t.X, Y: p.Y - t.Y}
}
