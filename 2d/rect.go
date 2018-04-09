package _d

/**
 * Rect 矩形
 */
type Rectangle struct {
	P [4]*Point
}

/**
 * Points 获取所有点
 */
func (r *Rectangle) Points() []*Point {
	return []*Point{r.P[0], r.P[1], r.P[2], r.P[3]}
}

/**
 * Vectors 获取所有向量 顺序是从 [0] 开始
 */
func (r *Rectangle) Vectors() []*Vector {
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
 * Perimeter 计算周长
 */
func (r *Rectangle) Perimeter() float64 {
	perimeter := 0.0
	for _, line := range r.Vectors() {
		perimeter += line.Length()
	}
	return perimeter
}

/**
 * Area 计算面积
 */
func (r *Rectangle) Area() float64 {
	/**
	 * 假设 A、B、C、D 四个点
	 * S□ = S(ABC) + S(ACD)
	 */
	s := 0.0
	A := r.P[0]
	for index := 1; index <= len(r.P)-1; index++ {
		s += (&Triangle{[3]*Point{A, r.P[index], r.P[index+1]}}).Area()
	}
	return s
}

/**
 * ContainPoint 矩形包含某点
 */
func (r *Rectangle) ContainPoint(point *Point) bool {
	j := len(r.P) - 1
	containPoint := false
	for index := 0; index < len(r.P); index++ {
		if r.P[index].Y < point.Y && r.P[j].Y >= point.Y ||
			r.P[j].Y < point.Y && r.P[index].Y >= point.Y {
			if r.P[index].X+(point.Y-r.P[index].Y)/
				(r.P[j].Y-r.P[index].Y)*
				(r.P[j].X-r.P[index].X) < point.X {
				containPoint = !containPoint
			}
		}
		j = index
	}
	return containPoint
}
