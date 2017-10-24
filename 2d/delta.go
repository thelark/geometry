package _d

/**
 * 三角形
 */
type Delta struct {
	P [3]*Point
}

/**
 * 获取所有点
 */
func (d *Delta) Points() []*Point {
	return []*Point{d.P[0], d.P[1], d.P[2]}
}

/**
 * 获取所有边 顺序是从 [0] 开始
 */
func (d *Delta) Lines() []*Vector {
	lines := make([]*Vector, 0)
	for k, point := range d.Points(){
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
func (d *Delta) Perimeter() float64 {
	perimeter := 0.0
	for _, line := range d.Lines() {
		perimeter += line.Length()
	}
	return perimeter
}

/**
 * 计算面积
 */
func (d *Delta) Area() float64 {
	/**
	 * 设有 A、B、C 三点
	 * SΔ = (1 / 2) * (A.X * B.Y + B.X * C.Y + C.X * A.Y - A.X * C.Y - B.X * A.Y - C.X * B.Y)
	 */
	return 1 / 2 * (d.P[0].X*d.P[1].Y + d.P[1].X*d.P[2].Y + d.P[2].X*d.P[0].Y - d.P[0].X*d.P[2].Y - d.P[1].X*d.P[0].Y - d.P[2].X*d.P[1].Y)
}
/**
 * 三角形包含某点
 */
func (d *Delta) ContainPoint(point *Point) bool {
	j := len(d.P) - 1
	containPoint := false
	for index := 0; index < len(d.P); index++ {
		if d.P[index].Y < point.Y && d.P[j].Y >= point.Y ||
			d.P[j].Y < point.Y && d.P[index].Y >= point.Y {
			if d.P[index].X + (point.Y - d.P[index].Y) /
				(d.P[j].Y - d.P[index].Y) *
				(d.P[j].X - d.P[index].X) < point.X {
				containPoint = !containPoint
			}
		}
		j = index
	}
	return containPoint
}