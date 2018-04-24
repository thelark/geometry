package _d

import "math"

/**
 * 多边形
 */
type Polygon struct {
	P []*Point
}

func (po *Polygon) Top() *Point {
	p := po.P[0]
	for _, v := range po.P {
		if v.Y >= p.Y {
			p = v
		}
	}
	return p
}
func (po *Polygon) Bottom() *Point {
	p := po.P[0]
	for _, v := range po.P {
		if v.Y <= p.Y {
			p = v
		}
	}
	return p
}
func (po *Polygon) Left() *Point {
	p := po.P[0]
	for _, v := range po.P {
		if v.X <= p.X {
			p = v
		}
	}
	return p
}
func (po *Polygon) Right() *Point {
	p := po.P[0]
	for _, v := range po.P {
		if v.X >= p.X {
			p = v
		}
	}
	return p
}

/**
 * Points 获取所有点
 */
func (p *Polygon) Points() []*Point {
	return p.P
}

/**
 * Vectors 获取所有向量 顺序是从 [0] 开始
 */
func (p *Polygon) Vectors() []*Vector {
	lines := make([]*Vector, 0)
	for k, point := range p.Points() {
		i := k + 1
		if k == len(p.P)-1 {
			i = 0
		}
		lines = append(lines, &Vector{point, p.P[i]})
	}
	return lines
}

/**
 * Perimeter 计算周长
 */
func (p *Polygon) Perimeter() float64 {
	perimeter := 0.0
	for _, line := range p.Vectors() {
		perimeter += line.Length()
	}
	return perimeter
}

/**
 * Area 计算面积
 */
func (p *Polygon) Area() float64 {
	s := 0.0
	A := p.P[0]
	for index := 1; index <= len(p.P)-1; index++ {
		s += (&Triangle{[3]*Point{A, p.P[index], p.P[index+1]}}).Area()
	}
	return s
}
/**
 * 面积计算法 （较为准确）
 */
func (p *Polygon) _Area() float64 {
	if len(p.P) < 3 {
		return 0
	}
	s := p.P[0].Y * (p.P[len(p.P)-1].X - p.P[1].X)
	for i := 1; i < len(p.P); i++ {
		s += p.P[i].Y * (p.P[i-1].X - p.P[(i+1)%len(p.P)].X)
	}
	return math.Abs(s)
}

/**
 * ContainPoint 多边形包含某点
 */
func (p *Polygon) ContainPoint(point *Point) bool {
	j := len(p.P) - 1
	containPoint := false
	for index := 0; index < len(p.P); index++ {
		if p.P[index].Y < point.Y && p.P[j].Y >= point.Y ||
			p.P[j].Y < point.Y && p.P[index].Y >= point.Y {
			if p.P[index].X + (point.Y - p.P[index].Y) /
				(p.P[j].Y - p.P[index].Y) *
				(p.P[j].X - p.P[index].X) < point.X {
				containPoint = !containPoint
			}
		}
		j = index
	}
	return containPoint
}

/**
 * MappingScaleTrimmingRect 映射缩放倍数的切边矩形
 */
func (p *Polygon) MappingScaleTrimmingRect(scaleX, scaleY float64) *Rectangle {
	// 1倍 矩形
	defaultRect := &Rectangle{
		[4]*Point{
			{p.Left().X, p.Top().Y},
			{p.Right().X, p.Top().Y},
			{p.Right().X, p.Bottom().Y},
			{p.Left().X, p.Bottom().Y},
		},
	}
	scaleRect := &Rectangle{
		[4]*Point{
			{
				defaultRect.P[0].X - 1.0/2*(&Vector{defaultRect.P[0], defaultRect.P[1]}).Length()*(scaleX-1),
				defaultRect.P[0].Y + 1.0/2*(&Vector{defaultRect.P[0], defaultRect.P[3]}).Length()*(scaleY-1),
			},
			{
				defaultRect.P[1].X + 1.0/2*(&Vector{defaultRect.P[0], defaultRect.P[1]}).Length()*(scaleX-1),
				defaultRect.P[1].Y + 1.0/2*(&Vector{defaultRect.P[1], defaultRect.P[2]}).Length()*(scaleY-1),
			},
			{
				defaultRect.P[2].X + 1.0/2*(&Vector{defaultRect.P[2], defaultRect.P[3]}).Length()*(scaleX-1),
				defaultRect.P[2].Y - 1.0/2*(&Vector{defaultRect.P[1], defaultRect.P[2]}).Length()*(scaleY-1),
			},
			{
				defaultRect.P[3].X - 1.0/2*(&Vector{defaultRect.P[2], defaultRect.P[3]}).Length()*(scaleX-1),
				defaultRect.P[3].Y - 1.0/2*(&Vector{defaultRect.P[0], defaultRect.P[3]}).Length()*(scaleY-1),
			},
		},
	}
	return scaleRect
}

/**
 * MappingTrimmingRect 映射切边矩形
 */
func (p *Polygon) MappingTrimmingRect() *Rectangle {
	return &Rectangle{
		[4]*Point{
			{p.Left().X, p.Top().Y},
			{p.Right().X, p.Top().Y},
			{p.Right().X, p.Bottom().Y},
			{p.Left().X, p.Bottom().Y},
		},
	}
}
