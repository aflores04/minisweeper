package game

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGame_SetMines(t *testing.T) {
	game := New(5, 5, 5)
	game.SetMines()
	game.SetValues()

	count := 0

	for _, pointRow := range game.Square.PointRows {
		for _, point := range pointRow.Points {
			if point.Mine {
				count++
			}
		}
	}

	assert.Equal(t, 5, count)
}

func TestGame_SetValues(t *testing.T) {
	game := New(3, 3, 3)
	game.SetMines()
	game.SetValues()
}

func TestGame_FindPointByPositionWithError(t *testing.T) {
	game := New(10, 10, 10)
	game.SetMines()

	_, err := game.FindPointByPosition(20, 20)

	assert.Equal(t, "can't find a point", err.Error())
}

func TestGame_FindPointByPosition(t *testing.T) {
	game := New(10, 10, 10)
	game.SetMines()

	point, _ := game.FindPointByPosition(2,2)

	assert.Equal(t, 2, point.Y)
	assert.Equal(t, 2, point.X)
}

func TestNewGame(t *testing.T) {
	game := New(10, 10, 10)

	assert.Equal(t, 10, game.Cols)
}

func TestNewSquare(t *testing.T) {
	tests := []struct {
		name string
		want *Square
	}{
		{
			name: "get new square",
			want: NewSquare(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSquare(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
