package _d

/**
 * 立方体
 */
type Cube struct {
	P [8]*Point
}

/**
 * 假设立方体在第一象限
 * 即 左-下-前 点为原点
 * 那么绘制各点的方向为 从原点出发 上 -> 右 -> 下 -> 后 -> 上 -> 左 -> 下
 */
func (c *Cube) Noodles() []*Noodles {
	retNoodles := make([]*Noodles, 0)
	retNoodles = append(retNoodles, &Noodles{[4]*Point{c.P[0], c.P[1], c.P[2], c.P[3]}})
	retNoodles = append(retNoodles, &Noodles{[4]*Point{c.P[2], c.P[3], c.P[4], c.P[5]}})
	retNoodles = append(retNoodles, &Noodles{[4]*Point{c.P[4], c.P[5], c.P[6], c.P[7]}})
	retNoodles = append(retNoodles, &Noodles{[4]*Point{c.P[6], c.P[7], c.P[0], c.P[1]}})

	retNoodles = append(retNoodles, &Noodles{[4]*Point{c.P[0], c.P[3], c.P[4], c.P[7]}})
	retNoodles = append(retNoodles, &Noodles{[4]*Point{c.P[1], c.P[2], c.P[5], c.P[6]}})
	return retNoodles
}

/**
 * 假设立方体在第一象限
 * 即 左-下-前 点为原点
 * 那么绘制各点的方向为 从原点出发 上 -> 右 -> 下 -> 后 -> 上 -> 左 -> 下
 */
func (c *Cube) Lines() []*Vector {
	retVector := make([]*Vector, 0)
	retVector = append(retVector, &Vector{c.P[0], c.P[1]})
	retVector = append(retVector, &Vector{c.P[1], c.P[2]})
	retVector = append(retVector, &Vector{c.P[2], c.P[3]})
	retVector = append(retVector, &Vector{c.P[3], c.P[4]})
	retVector = append(retVector, &Vector{c.P[4], c.P[5]})
	retVector = append(retVector, &Vector{c.P[5], c.P[6]})
	retVector = append(retVector, &Vector{c.P[6], c.P[7]})
	retVector = append(retVector, &Vector{c.P[7], c.P[0]})

	retVector = append(retVector, &Vector{c.P[1], c.P[6]})
	retVector = append(retVector, &Vector{c.P[2], c.P[5]})
	retVector = append(retVector, &Vector{c.P[3], c.P[0]})
	retVector = append(retVector, &Vector{c.P[4], c.P[7]})
	return retVector
}
