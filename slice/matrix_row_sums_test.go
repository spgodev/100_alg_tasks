package slice

import (
	"reflect"
	"testing"
)

func TestMatrixRowSums(t *testing.T) {
	tests := []struct {
		name string
		matrix [][]int
		want []int
	}{
		{
			name: "обычный случай",
			matrix: [][]int{[]int{1, 2}, []int{3, 4}},
			want: []int{3, 7},
		},
		{
			name: "есть пустая строка",
			matrix: [][]int{[]int{}, []int{5}},
			want: []int{0, 5},
		},
		{
			name: "nil матрица",
			matrix: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MatrixRowSums(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixRowSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
