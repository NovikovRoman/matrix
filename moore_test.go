package matrix

import (
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/require"
	"image/color"
	"path/filepath"
	"testing"
)

func TestMatrix_FindDiagonalPathMooreWave(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "lab.png"))
	require.Nil(t, err)

	m := NewMatrixFromImage(im, ColorObstacles{color.Gray{Y: 255}})

	startPoint := PointZero()
	endPoint := NewPoint(19, 19)

	ok, path := m.FindDiagonalPathMooreWave(startPoint, endPoint)
	require.True(t, ok)
	require.Len(t, path, 49)

	saveMatrix(t, "FindDiagonalPathMooreWave_result.png", m, path, 3)

	m = NewMatrix(5, 5, 0)
	m.SetValue(NewPoint(2, 0), Obstacle)
	m.SetValue(NewPoint(2, 1), Obstacle)
	m.SetValue(NewPoint(2, 2), Obstacle)
	m.SetValue(NewPoint(0, 2), Obstacle)
	m.SetValue(NewPoint(1, 2), Obstacle)

	ok, path = m.FindDiagonalPathMooreWave(startPoint, endPoint)
	require.False(t, ok)
	require.Nil(t, path)

	endPoint = NewPoint(4, 4)
	ok, path = m.FindDiagonalPathMooreWave(startPoint, endPoint)
	require.False(t, ok)
	require.Nil(t, path)
}

func TestMatrix_FindDirectPathMooreWave(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "lab.png"))
	require.Nil(t, err)

	m := NewMatrixFromImage(im, ColorObstacles{color.Gray{Y: 255}})

	startPoint := PointZero()
	endPoint := NewPoint(19, 19)

	ok, path := m.FindDirectPathMooreWave(startPoint, endPoint)
	require.True(t, ok)
	require.Len(t, path, 65)

	saveMatrix(t, "FindDirectPathMooreWave_result.png", m, path, 3)

	m = NewMatrix(5, 5, 0)
	m.SetValue(NewPoint(2, 0), Obstacle)
	m.SetValue(NewPoint(2, 1), Obstacle)
	m.SetValue(NewPoint(2, 2), Obstacle)
	m.SetValue(NewPoint(0, 2), Obstacle)
	m.SetValue(NewPoint(1, 2), Obstacle)

	ok, path = m.FindDirectPathMooreWave(startPoint, endPoint)
	require.False(t, ok)
	require.Nil(t, path)

	endPoint = NewPoint(4, 4)
	ok, path = m.FindDirectPathMooreWave(startPoint, endPoint)
	require.False(t, ok)
	require.Nil(t, path)
}

func TestMatrix_MooreWavePropagation(t *testing.T) {
	im, err := imaging.Open(filepath.Join(testdata, "lab.png"))
	require.Nil(t, err)

	m := NewMatrixFromImage(im, ColorObstacles{color.Gray{Y: 255}})
	startPoint := NewPoint(19, 22)
	endPoint := PointZero()
	m.MooreWavePropagation(endPoint, []Point{startPoint}, 5)

	saveMatrix(t, "MooreWavePropagation_limit_result.png", m, Path{}, 20)

	m.Reset()
	ok := m.MooreWave(startPoint, endPoint, 0)
	require.True(t, ok)
}
