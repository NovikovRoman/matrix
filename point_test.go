package matrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(0, 0)
	p2 := NewPoint(3, 2)

	p.Add(p2)
	require.Equal(t, p.X, 3)
	require.Equal(t, p.Y, 2)
	require.Equal(t, p.String(), "(3, 2)")

	p.Mul(3)
	require.Equal(t, p.X, 9)
	require.Equal(t, p.Y, 6)
	require.Equal(t, p.String(), "(9, 6)")

	p.Sub(p2)
	require.Equal(t, p.X, 6)
	require.Equal(t, p.Y, 4)
	require.Equal(t, p.String(), "(6, 4)")

	p.Div(3)
	require.Equal(t, p.X, 2)
	require.Equal(t, p.Y, 1)
	require.Equal(t, p.String(), "(2, 1)")
}
