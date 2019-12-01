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

type Points []Point

func (p Points) Search(point Point) int {
	for i := range p {
		if p[i].Eq(point) {
			return i
		}
	}
	return -1
}

func (p Points) GetIndex(point Point) int {
	return p.Search(point)
}

func (p *Points) Add(points ...Point) int {
	for _, point := range points {
		if p.Search(point) > -1 {
			continue
		}
		*p = append(*p, point)
	}

	return len(*p)
}

func (p *Points) Remove(points ...Point) int {
	for _, point := range points {
		i := p.Search(point)
		if i == -1 {
			continue
		}

		b := Points{}

		if i == 0 {
			b = (*p)[1:]

		} else if i == len(*p)-1 {
			b = (*p)[:i]

		} else {
			b = (*p)[:i]
			b = append(b, (*p)[i+1:]...)
		}

		*p = b
	}
	return len(*p)
}

func NewPoints(length int) Points {
	return make([]Point, length)
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func PointZero() Point {
	return Point{X: 0, Y: 0}
}
