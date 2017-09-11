package _d

/**
 * 多边形
 */
type Polygon struct {
	P []*Point
}

/**
 * 获取所有点
 */
func (p *Polygon) Points() []*Point {
	return p.P
}

/**
 * 获取所有边 顺序是从 [0] 开始
 */
func (p *Polygon) Lines() []*Vector {
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
 * 计算周长
 */
func (p *Polygon) Perimeter() float64 {
	perimeter := 0.0
	for _, line := range p.Lines() {
		perimeter += line.Length()
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
		s += (&Delta{[3]*Point{A, p.P[index], p.P[index+1]}}).Area()
	}
	return s
}

/**
 * 多边形包含某点
 */
func (p *Polygon) ContainPoint(point *Point) bool {

}
