package game

import (
	"fmt"
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	game := New(10, 10, 10)

	assert.Equal(t, 10, game.Cols)
}

func TestPoints(t *testing.T) {
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
