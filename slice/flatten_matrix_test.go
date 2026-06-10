package slice

import (
	"reflect"
	"testing"
)

func TestFlattenMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "прямоугольная матрица",
			matrix: [][]int{[]int{1, 2}, []int{3, 4}},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "строки разной длины",
			matrix: [][]int{[]int{1}, []int{2, 3}, []int{}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "nil матрица",
			matrix: nil,
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenMatrix(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlattenMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
