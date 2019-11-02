package matrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPath_CommonPoints(t *testing.T) {
	p1 := Path{
		NewPoint(0, 0),
		NewPoint(1, 0),
		NewPoint(0, 1),
		NewPoint(2, 2),
		NewPoint(4, 2),
		NewPoint(1, 12),
		NewPoint(21, 2),
	}
	p2 := Path{
		NewPoint(0, 0),
		NewPoint(1, 0),
		NewPoint(4, 1),
		NewPoint(2, 2),
		NewPoint(4, 6),
		NewPoint(1, 1),
		NewPoint(1, 2),
	}

	p := p1.CommonPoints(p2)
	require.Len(t, p, 3)
	require.True(t, p[0].Eq(NewPoint(0, 0)))
	require.True(t, p[1].Eq(NewPoint(1, 0)))
	require.True(t, p[2].Eq(NewPoint(2, 2)))
}
