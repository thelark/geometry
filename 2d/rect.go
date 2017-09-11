package _d

/**
 * 矩形
 */
type Rect struct {
	A *Point
	B *Point
	C *Point
	D *Point
}

/**
 * 获取所有点
 */
func (r *Rect) Points() []*Point {
	return []*Point{r.A, r.B, r.C, r.D}
}

/**
 * 计算周长
 */
func (r *Rect) Perimeter() float64 {
	return (&Vector{r.A, r.B}).Length() +
		(&Vector{r.B, r.C}).Length() +
		(&Vector{r.C, r.D}).Length() +
		(&Vector{r.D, r.A}).Length()
}

/**
 * 计算面积
 */
func (r *Rect) Area() float64 {
	/**
	 * 假设 A、B、C、D 四个点
	 * S□ = S(ABC) + S(ACD)
	 */
	return (&Triangle{r.A, r.B, r.C}).Area() + (&Triangle{r.A, r.C, r.D}).Area()
}
