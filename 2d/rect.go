package _d

/**
 * Rect 矩形
 */
type Rectangle struct {
	P [4]*Point
}

func (r *Rectangle) Top() *Point {
	p := r.P[0]
	for _, v := range r.P {
		if v.Y >= p.Y {
			p = v
		}
	}
	return p
}
func (r *Rectangle) Bottom() *Point {
	p := r.P[0]
	for _, v := range r.P {
		if v.Y <= p.Y {
			p = v
		}
	}
	return p
}
func (r *Rectangle) Left() *Point {
	p := r.P[0]
	for _, v := range r.P {
		if v.X <= p.X {
			p = v
		}
	}
	return p
}
func (r *Rectangle) Right() *Point {
	p := r.P[0]
	for _, v := range r.P {
		if v.X >= p.X {
			p = v
		}
	}
	return p
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
			if r.P[index].X + (point.Y - r.P[index].Y) /
				(r.P[j].Y - r.P[index].Y) *
				(r.P[j].X - r.P[index].X) < point.X {
				containPoint = !containPoint
			}
		}
		j = index
	}
	return containPoint
}

/**
 * Split 矩形分割
 * 分割后按从左到右，从下到上返回
 */
func (r *Rectangle) Split(splitCount int) []*Rectangle {
	YSplit := r.Top().Y - r.Bottom().Y // Y轴分割距
	XSplit := r.Right().X - r.Left().X // X轴分割距
	XMin := r.Left().X
	//XMax := r.Right().X
	YMin := r.Bottom().Y
	//YMax := r.Top().Y
	rects := make([]*Rectangle, 0)
	for i := 0; i < splitCount; i++ {
		for j := 0; j < splitCount; j++ {
			rect := &Rectangle{[4]*Point{
				{XMin + float64(i)*XSplit, YMin + float64(j)*YSplit},
				{XMin + float64(i)*XSplit, YMin + float64(j+1)*YSplit},
				{XMin + float64(i+1)*XSplit, YMin + float64(j+1)*YSplit},
				{XMin + float64(i+1)*XSplit, YMin + float64(j)*YSplit},
			}}
			rects = append(rects, rect)
		}
	}
	return rects
}
