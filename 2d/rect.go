package _d

/**
 * 矩形
 */
type Rect struct {
	P [4]*Point
}

/**
 * 获取所有点
 */
func (r *Rect) Points() []*Point {
	return []*Point{r.P[0], r.P[1], r.P[2], r.P[3]}
}

/**
 * 获取所有边 顺序是从 [0] 开始
 */
func (r *Rect) Lines() []*Vector {
	lines := make([]*Vector, 0)
	for k, point := range r.Points() {
		i := k + 1
		if k == len(r.P)-1 {
			i = 0
		}
		lines = append(lines, &Vector{point, r.P[i]})
	}
	return lines
}

/**
 * 计算周长
 */
func (r *Rect) Perimeter() float64 {
	perimeter := 0.0
	for _, line := range r.Lines() {
		perimeter += line.Length()
	}
	return perimeter
}

/**
 * 计算面积
 */
func (r *Rect) Area() float64 {
	/**
	 * 假设 A、B、C、D 四个点
	 * S□ = S(ABC) + S(ACD)
	 */
	s := 0.0
	A := r.P[0]
	for index := 1; index <= len(r.P)-1; index++ {
		s += (&Delta{[3]*Point{A, r.P[index], r.P[index+1]}}).Area()
	}
	return s
}
