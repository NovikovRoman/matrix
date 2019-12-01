package matrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPoint(t *testing.T) {
	points := Points{}

	p := NewPoint(0, 0)
	p2 := NewPoint(3, 2)
	points.Add(p, p2)

	p.Add(p2)
	require.Equal(t, p.X, 3)
	require.Equal(t, p.Y, 2)
	require.Equal(t, p.String(), "(3, 2)")
	x, y := p.Separate()
	require.Equal(t, x, 3)
	require.Equal(t, y, 2)
	points.Add(p) // не добавится, тк уже есть точка с такими координатами

	p.Mul(3)
	require.Equal(t, p.X, 9)
	require.Equal(t, p.Y, 6)
	require.Equal(t, p.String(), "(9, 6)")
	x, y = p.Separate()
	require.Equal(t, x, 9)
	require.Equal(t, y, 6)
	points.Add(p)

	p.Sub(p2)
	require.Equal(t, p.X, 6)
	require.Equal(t, p.Y, 4)
	require.Equal(t, p.String(), "(6, 4)")
	x, y = p.Separate()
	require.Equal(t, x, 6)
	require.Equal(t, y, 4)
	points.Add(p)

	p.Div(3)
	require.Equal(t, p.X, 2)
	require.Equal(t, p.Y, 1)
	require.Equal(t, p.String(), "(2, 1)")
	x, y = p.Separate()
	require.Equal(t, x, 2)
	require.Equal(t, y, 1)
	points.Add(p)

	require.Len(t, points, 5)
	require.True(t, points[4].Eq(p))
	points.Remove(PointZero(), PointZero(), NewPoint(3, 2), NewPoint(3, 2))
	require.Len(t, points, 3)
	require.True(t, points[2].Eq(p))
}
