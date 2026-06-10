package slice

import (
	"testing"
)

func TestDiagonalSum(t *testing.T) {
	tests := []struct {
		name string
		matrix [][]int
		want int
	}{
		{
			name: "квадратная матрица",
			matrix: [][]int{[]int{1, 2}, []int{3, 4}},
			want: 5,
		},
		{
			name: "прямоугольная матрица",
			matrix: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
			want: 6,
		},
		{
			name: "пустая матрица",
			matrix: [][]int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DiagonalSum(tt.matrix)
			if got != tt.want {
				t.Errorf("DiagonalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
