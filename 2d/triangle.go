package _d

/**
 * 三角形
 */
type Triangle struct {
	A *Point // 点 A
	B *Point // 点 B
	C *Point // 点 C
}

/**
 * 获取所有点
 */
func (t *Triangle) Points() []*Point {
	return []*Point{t.A, t.B, t.C}
}

/**
 * 计算周长
 */
func (t *Triangle) Perimeter() float64 {
	return (&Vector{t.A, t.B}).Length() +
		(&Vector{t.B, t.C}).Length() +
		(&Vector{t.C, t.D}).Length()
}

/**
 * 计算面积
 */
func (t *Triangle) Area() float64 {
	/**
	 * 设有 A、B、C 三点
	 * SΔ = (1 / 2) * (A.X * B.Y + B.X * C.Y + C.X * A.Y - A.X * C.Y - B.X * A.Y - C.X * B.Y)
	 */
	return 1 / 2 * (t.A.X*t.B.Y + t.B.X*t.C.Y + t.C.X*t.A.Y - t.A.X*t.C.Y - t.B.X*t.A.Y - t.C.X*t.B.Y)
}
