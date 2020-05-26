package matrix

// Двумерная окрестность фон Неймана порядка 1
//
// ⬜ | ⬛ | ⬜
//
// ⬛ | ⬛ | ⬛
//
// ⬜ | ⬛ | ⬜
func (m Matrix) NeumannWave(startPoint Point, endPoint Point, limitValue float64) bool {
	m.SetValue(startPoint, 0.1)
	return m.NeumannWavePropagation(endPoint, []Point{startPoint}, limitValue)
}

func (m Matrix) findPathNeumannWave(startPoint Point, endPoint Point) (bool, Path) {
	m.SetValue(startPoint, 0.1)
	if m.NeumannWavePropagation(endPoint, []Point{startPoint}, 0) {
		return true, Path{endPoint}
	}

	return false, nil
}

// Поиск пути используя окрестность фон Неймана
func (m Matrix) FindDiagonalPathNeumannWave(startPoint Point, endPoint Point) (bool, Path) {
	ok, path := m.findPathNeumannWave(startPoint, endPoint)
	if ok {
		m.nextPathPoint(&path, startPoint, endPoint, false)
		path.Reverse()
	}
	return ok, path
}

// Поиск пути используя окрестность фон Неймана
func (m Matrix) FindDirectPathNeumannWave(startPoint Point, endPoint Point) (bool, Path) {
	ok, path := m.findPathNeumannWave(startPoint, endPoint)
	if ok {
		m.nextPathPoint(&path, startPoint, endPoint, true)
		path.Reverse()
	}
	return ok, path
}

// Рекурсивная функция
func (m Matrix) NeumannWavePropagation(endPoint Point, points []Point, limitValue float64) bool {
	if len(points) == 0 {
		return false
	}

	var nextPoints []Point
	nextPoints = []Point{}

	//обходим точки
	for _, currPoint := range points {
		v := m.Value(currPoint) + 1 // todo: может регулируемое сделать?

		// достигли лимита
		if limitValue > 0 && v > limitValue {
			return false
		}

		if currPoint.Eq(endPoint) {
			return true
		}

		// проверяем клетки
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {

				if x == 0 && y == 0 { // currPoint не рассматриваем
					continue
				}

				p := NewPoint(currPoint.X+x, currPoint.Y+y)
				// ячейка за гранью или это препятствие или это диагональные точки
				if m.IsObstacle(p) || p.TouchCorner(currPoint) {
					continue
				}

				if m.Value(p) == 0 {
					nextPoints = append(nextPoints, p)

					// обработана/занята
				} else if m.Value(p) <= v {
					continue
				}

				m.SetValue(p, v)
			}
		}
	}

	//повторяем для следующих клеток
	return m.NeumannWavePropagation(endPoint, nextPoints, limitValue)
}
