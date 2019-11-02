package matrix

import (
	"image"
	"image/color"
)

const Obstacle = -1.0
const Empty = 0.0

// матрица строки x столбцы
type Matrix [][]float64

type ColorObstacles []color.Color

func (c ColorObstacles) Exists(color color.Color) bool {
	for _, cc := range c {
		if cc == color {
			return true
		}
	}
	return false
}

// Создать матрицу и заполнить значениями value
func NewMatrix(width int, height int, value float64) Matrix {
	m := make([][]float64, height)

	for y := range m {
		m[y] = make([]float64, width)
		if value != 0 {
			for x := range m[y] {
				m[y][x] = value
			}
		}
	}

	return m
}

// Создать матрицу из image.Gray
func NewMatrixFromImage(im image.Image, colorObstacles ColorObstacles) Matrix {
	m := make([][]float64, im.Bounds().Max.Y)

	for y := range m {
		m[y] = make([]float64, im.Bounds().Max.X)
		for x := range m[y] {
			if colorObstacles.Exists(im.At(x, y)) {
				m[y][x] = Obstacle
			}
		}
	}

	return m
}

func (m Matrix) Height() int {
	return len(m)
}

func (m Matrix) Width() int {
	return len(m[0])
}

func (m Matrix) Value(p Point) float64 {
	return m[p.Y][p.X]
}

func (m Matrix) SetValue(p Point, v float64) {
	m[p.Y][p.X] = v
}

func (m Matrix) Reset() {
	for y, row := range m {
		for x := range row {
			if m[y][x] > 0 {
				m[y][x] = 0
			}
		}
	}
}

// Точка является препятствием
func (m Matrix) IsObstacle(p Point) bool {
	return !p.In(m) || m.Value(p) == Obstacle
}

func (m Matrix) nextPathPoint(path *Path, startPoint Point, checkPoint Point, direct bool) {
	if startPoint.Eq(checkPoint) {
		return
	}

	newPoint := NewPoint(-1, -1)
	//проверяем окружные 8 клеток
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {

			if x == 0 && y == 0 { // checkPoint не рассматриваем
				continue
			}

			p := NewPoint(checkPoint.X+x, checkPoint.Y+y)
			// вне матрицы
			if !p.In(m) || direct && p.TouchCorner(checkPoint) {
				continue
			}

			value := m.Value(p)
			// свободная клетка или коэффициент больше проверочной точки
			if value <= Empty || value > Empty && m.Value(checkPoint) < value {
				continue
			}

			if (newPoint.X == -1 && newPoint.Y == -1) || value < m.Value(newPoint) {
				newPoint = p
			}
		}
	}

	if newPoint.X == -1 || newPoint.Y == -1 {
		newPoint = startPoint
	}

	*path = append(*path, newPoint)
	m.nextPathPoint(path, startPoint, newPoint, direct)
}

/*
Точки соприкасаются углами и углы между препятствий

Есть препятствие:

X | ⬛

⬛ | X

нет препятствий:

X | ⬜

⬜ | X

и здесь тоже:

X | ⬛

⬜ | X

и здесь:

X | ⬜

⬛ | X

где,

⬜ - пусто

⬛ - препятствие

X - проверяемые точки
*/
func (m Matrix) TouchCornerObstacle(p1 Point, p2 Point) bool {
	if !p1.TouchCorner(p2) {
		return false
	}

	pT1 := NewPoint(p1.X, p2.Y)
	pT2 := NewPoint(p2.X, p1.Y)
	return m.IsObstacle(pT1) && m.IsObstacle(pT2)
}
