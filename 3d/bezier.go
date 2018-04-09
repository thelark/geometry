package _d

import "math"

func Bezier(ps []*Point, precision int) (result []*Point) {
	result = make([]*Point, precision)
	// 贝塞尔曲线控制点数（阶数）
	var number = len(ps)

	// 控制点数不小于 2
	if number < 2 {
		return
	}

	//计算杨辉三角
	var mi = make([]int, number, number)
	mi[0] = 1
	mi[1] = 1
	for i := 3; i <= number; i++ {
		t := make([]int, i-1)
		for j := 0; j < len(t); j++ {
			t[j] = mi[j]
		}

		mi[0], mi[i-1] = 1, 1
		for j := 0; j < i-2; j++ {
			mi[j+1] = t[j] + t[j+1]
		}
	}
	//计算坐标点
	for i := 0; i < precision; i++ {
		var t = float64(i) / float64(precision)
		X := 0.0
		Y := 0.0
		Z := 0.0
		for k := 0; k < number; k++ {
			X += math.Pow(1-t, float64(number-k-1)) *
				ps[k].X *
				math.Pow(t, float64(k)) * float64(mi[k])
			Y += math.Pow(1-t, float64(number-k-1)) *
				ps[k].Y *
				math.Pow(t, float64(k)) * float64(mi[k])
			Z += math.Pow(1-t, float64(number-k-1)) *
				ps[k].Z *
				math.Pow(t, float64(k)) * float64(mi[k])
		}
		result[i] = &Point{X, Y, Z}
	}
	// 加入最后一个点坐标
	result = append(result, ps[len(ps)-1])
	return
}
