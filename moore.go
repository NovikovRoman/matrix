package matrix

// Распространение волны по окрестности Мура.
//
// Двумерная окрестность Мура порядка 1:
//
// ⬛ | ⬛ | ⬛
//
// ⬛ | X | ⬛
//
// ⬛ | ⬛ | ⬛
//
// endPoint - финишная точка
//
// points - список точек, из которых необходимо продолжить волну
//
//Волна до значения не более waveRadius.
func (m Matrix) MooreWave(startPoint Point, endPoint Point, limitValue float64) bool {
	m.SetValue(startPoint, 0.1)
	return m.MooreWavePropagation(endPoint, []Point{startPoint}, limitValue)
}

func (m Matrix) findPathMooreWave(startPoint Point, endPoint Point) (bool, Path) {
	m.SetValue(startPoint, 0.1)
	if m.MooreWavePropagation(endPoint, []Point{startPoint}, 0) {
		return true, Path{endPoint}
	}

	return false, nil
}

// Поиск пути используя окрестность Мура
func (m Matrix) FindDiagonalPathMooreWave(startPoint Point, endPoint Point) (bool, Path) {
	ok, path := m.findPathMooreWave(startPoint, endPoint)
	if ok {
		m.nextPathPoint(&path, startPoint, endPoint, false)
		path.Reverse()
	}
	return ok, path
}

// Поиск пути используя окрестность Мура
func (m Matrix) FindDirectPathMooreWave(startPoint Point, endPoint Point) (bool, Path) {
	ok, path := m.findPathMooreWave(startPoint, endPoint)
	if ok {
		m.nextPathPoint(&path, startPoint, endPoint, true)
		path.Reverse()
	}
	return ok, path
}

// Рекурсивная функция
func (m Matrix) MooreWavePropagation(endPoint Point, points []Point, limitValue float64) bool {
	if len(points) == 0 {
		return false
	}

	var nextPoints []Point
	nextPoints = []Point{}

	//обходим точки
	for _, currPoint := range points {
		vDirect := m.Value(currPoint) + 1 // todo: может регулироемое сделать?
		vCorner := vDirect + 0.2          // todo: может регулироемое сделать?

		// достигли лимита
		if limitValue > 0 && vDirect > limitValue {
			return false
		}

		if currPoint.Eq(endPoint) {
			return true
		}

		// проверяем окружные 8 клеток
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {

				if x == 0 && y == 0 { // currPoint не рассматриваем
					continue
				}

				p := NewPoint(currPoint.X+x, currPoint.Y+y)
				if m.IsObstacle(p) { // ячейка за гранью или это препятствие
					continue
				}

				if m.TouchCornerObstacle(p, currPoint) {
					continue
				}

				v := vDirect
				if p.TouchCorner(currPoint) {
					v = vCorner
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
	return m.MooreWavePropagation(endPoint, nextPoints, limitValue)
}
