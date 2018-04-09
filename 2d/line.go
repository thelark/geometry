package _d

import "math"

/**
 * Line 线段
 */
type Line struct {
	P [2]*Point
}

/**
 * 线段长度
 */
func (l *Line) Length() float64 {
	return math.Sqrt(math.Pow(float64(l.P[0].X-l.P[1].X), 2) + math.Pow(float64(l.P[0].Y-l.P[1].Y), 2))
}

// Intersect 判断两线是否相交
func (l *Line) Intersect(tmpL *Line) bool {
	return (math.Max(l.P[0].X, l.P[1].X) >= math.Min(tmpL.P[0].X, tmpL.P[1].X)) && //l中最右的点是否在tmpL最左的点的右边
		(math.Max(tmpL.P[0].X, tmpL.P[1].X) >= math.Min(l.P[0].X, l.P[1].X)) && //tmpL中最右的点是否在l最左的点的右边
	//判断这两条线段在水平层面上是否可能相交
		(math.Max(l.P[0].Y, l.P[1].Y) >= math.Min(tmpL.P[0].Y, tmpL.P[1].Y)) && //l中最上的点是否在tmpL最下的点的上边
		(math.Max(tmpL.P[0].Y, tmpL.P[1].Y) >= math.Min(l.P[0].Y, l.P[1].Y)) && //tmpL中最上的点是否在l最下的点的上边
	//判断这两条线段在垂直层面上是否可能相交
		(multiply(tmpL.P[0], l.P[1], l.P[0])*multiply(l.P[1], tmpL.P[1], l.P[0]) >= 0) &&
	//判断tmpL.P[0],tmpL.P[1]是否分布在l.P[1]两侧（或线上）
		(multiply(l.P[0], tmpL.P[1], tmpL.P[0])*multiply(tmpL.P[1], l.P[1], tmpL.P[0]) >= 0)
	//判断l.P[0],l.P[1]是否分布在tmpL.P[0]两侧（或线上）
}

// GetK 线段斜率
func (l *Line) GetK() float64 {
	return (l.P[1].Y - l.P[0].Y) / (l.P[1].X - l.P[0].X)
}

// GetB 线段偏移量
func (l *Line) GetB() float64 {
	return l.P[0].Y - l.GetK()*l.P[0].X
}

// 根据 Y 获取线段上对应的 P(X,Y)点
func (l *Line) GetPointByY(Y float64) *Point {
	if l.P[1].Y-l.P[0].Y == 0 {
		return &Point{l.P[1].X, Y}
	}
	X := (Y - l.GetB()) / l.GetK()
	return &Point{X, Y}
}

// 根据 X 获取线段上对应的 P(X,Y)点
func (l *Line) GetPointByX(X float64) *Point {
	if l.P[1].X-l.P[0].X == 0 {
		return &Point{X, l.P[1].Y}
	}
	Y := l.GetK()*X + l.GetB()
	return &Point{X, Y}
}
