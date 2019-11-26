package matrix

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Eq(point Point) bool {
	return p.X == point.X && p.Y == point.Y
}

func (p *Point) In(m Matrix) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}

	return p.X < m.Width() && p.Y < m.Height()
}

func (p *Point) Add(point Point) {
	p.X += point.X
	p.Y += point.Y
}

func (p *Point) Sub(point Point) {
	p.X -= point.X
	p.Y -= point.Y
}

func (p *Point) Mul(k int) {
	p.X *= k
	p.Y *= k
}

func (p *Point) Div(k int) {
	p.X /= k
	p.Y /= k
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p Point) Separate() (int, int) {
	return p.X, p.Y
}

func (p Point) TouchCorner(point Point) bool {
	return math.Abs(float64(p.X-point.X))+math.Abs(float64(p.Y-point.Y)) == 2
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func PointZero() Point {
	return Point{X: 0, Y: 0}
}
