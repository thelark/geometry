package _d

/**
 * 多边形
 */
type Polygon struct {
	P []*Point
}

/**
 * 计算周长
 */
func (p *Polygon) Perimeter() float64 {
	perimeter := 0.0
	for k, point := range p.P {
		i := k
		if k == len(p.P)-1 {
			i = 0
		}
		perimeter += (&Vector{point, p.P[i]}).Length()
	}
	return perimeter
}

/**
 * 计算面积
 */
func (p *Polygon) Area() float64 {
	s := 0.0
	A := p.P[0]
	for index := 1; index <= len(p.P)-1; index++ {
		s += (&Triangle{A, p.P[index], p.P[index+1]}).Area()
	}
	return s
}
