package matrix

type Path []Point

func (p Path) Reverse() {
	for i := len(p)/2 - 1; i >= 0; i-- {
		opp := len(p) - 1 - i
		p[i], p[opp] = p[opp], p[i]
	}
}

func (p Path) Contains(point Point) bool {
	for _, pathPoint := range p {
		if pathPoint.Eq(point) {
			return true
		}
	}
	return false
}

func (p Path) CommonPoints(path Path) []Point {
	var result []Point
	result = []Point{}

	for _, pathPoint := range path {
		if p.Contains(pathPoint) {
			result = append(result, pathPoint)
		}
	}

	return result
}
