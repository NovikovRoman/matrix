package matrix

import (
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/require"
	"image/color"
	"path/filepath"
	"testing"
)

func TestMatrix_FindDiagonalPathNeumannWave(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "lab.png"))
	require.Nil(t, err)

	m := NewMatrixFromImage(im, ColorObstacles{color.Gray{Y: 255}})

	startPoint := PointZero()
	endPoint := NewPoint(19, 19)

	ok, path := m.FindDiagonalPathNeumannWave(startPoint, endPoint)
	require.True(t, ok)
	require.Len(t, path, 49)

	saveMatrix(t, "FindDiagonalPathNeumannWave_result.png", m, path,3)
}

func TestMatrix_FindDirectPathNeumannWave(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "lab.png"))
	require.Nil(t, err)

	m := NewMatrixFromImage(im, ColorObstacles{color.Gray{Y: 255}})

	startPoint := PointZero()
	endPoint := NewPoint(19, 19)

	ok, path := m.FindDirectPathNeumannWave(startPoint, endPoint)
	require.True(t, ok)
	require.Len(t, path, 65)

	saveMatrix(t, "FindDirectPathNeumannWave_result.png", m, path, 3)
}

func TestMatrix_NeumannWavePropagation(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "lab.png"))
	require.Nil(t, err)

	m := NewMatrixFromImage(im, ColorObstacles{color.Gray{Y: 255}})
	startPoint := NewPoint(19, 22)
	endPoint := PointZero()
	m.NeumannWavePropagation(endPoint, []Point{startPoint}, 4)

	saveMatrix(t, "NeumannWavePropagation_limit_result.png", m, Path{}, 20)

	m.Reset()
	ok := m.NeumannWave(startPoint, endPoint, 0)
	require.True(t, ok)

}
