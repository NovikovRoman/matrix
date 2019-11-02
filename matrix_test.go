package matrix

import (
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"path/filepath"
	"testing"
)

const testdata = "testdata"

func TestNewMatrix(t *testing.T) {
	m := NewMatrix(10, 5, 12.3)
	require.Equal(t, m.Width(), 10)
	require.Equal(t, m.Height(), 5)

	pZero := NewPoint(0, 0)
	m.SetValue(pZero, -1)
	m.Reset()

	for y, row := range m {
		for x := range row {
			point := NewPoint(x, y)
			if pZero.Eq(point) {
				require.True(t, pZero.Eq(PointZero()))
				continue
			}

			require.Equal(t, m.Value(point), 0.0)
		}
	}
}

func saveMatrix(t *testing.T, filename string, m Matrix, path Path, k float64) {
	imSet := image.NewNRGBA(image.Rect(0, 0, m.Width()-1, m.Height()-1))

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	for y, row := range m {
		for x := range row {
			p := NewPoint(x, y)
			if path.Contains(p) {
				imSet.Set(x, y, red)
				continue
			}

			value := m.Value(p)
			if value < 0 {
				imSet.Set(x, y, color.White)
				continue
			}

			if value == 0 {
				imSet.Set(x, y, color.Black)
				continue
			}

			value *= k
			c := color.RGBA{R: uint8(value), G: uint8(value), B: uint8(value), A: 255}
			imSet.Set(x, y, c)
		}
	}

	err := imaging.Save(imSet, filepath.Join(testdata, filename))
	require.Nil(t, err)
}
