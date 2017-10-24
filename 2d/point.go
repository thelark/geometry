package _d

import (
	"strings"
	"strconv"
)

/**
 * 2D 点
 */
type Point struct {
	X float64 // X
	Y float64 // Y
}

// 判断点是否在向量左侧
func (p *Point) IsLeftInVector(v *Vector) {

}

// TODO: 判断点是否在形状内部
func (p *Point) IsInShape(s Shape) bool {
	return s.ContainPoint(p)
}

func ParsePoint(str string) *Point {
	pointStr := strings.Split(str, ",")
	x, err := strconv.ParseFloat(pointStr[0], 64)
	if err != nil {
	}
	y, err := strconv.ParseFloat(pointStr[1], 64)
	if err != nil {
	}
	return &Point{x, y}
}
