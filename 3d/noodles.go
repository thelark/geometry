package _d

/**
 * 面
 */
type Noodles struct {
	P [4]*Point
}

/**
 * 获取所有点
 */
func (n *Noodles) Points() []*Point {
	return []*Point{n.P[0], n.P[1], n.P[2], n.P[3]}
}

/**
 * 获取所有边 顺序是从 [0] 开始
 */
func (n *Noodles) Lines() []*Vector {
	lines := make([]*Vector, 0)
	for k, point := range n.Points() {
		i := k + 1
		if k == len(n.P)-1 {
			i = 0
		}
		lines = append(lines, &Vector{point, n.P[i]})
	}
	return lines
}
/**
 * 计算周长
 */
func (n *Noodles) Perimeter() float64 {
	perimeter := 0.0
	for _, line := range n.Lines() {
		perimeter += line.Length()
	}
	return perimeter
}