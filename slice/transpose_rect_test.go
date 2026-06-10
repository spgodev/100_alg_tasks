package slice

import (
	"reflect"
	"testing"
)

func TestTransposeRect(t *testing.T) {
	tests := []struct {
		name string
		matrix [][]int
		want [][]int
	}{
		{
			name: "два на три",
			matrix: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
			want: [][]int{[]int{1, 4}, []int{2, 5}, []int{3, 6}},
		},
		{
			name: "одна строка",
			matrix: [][]int{[]int{1, 2}},
			want: [][]int{[]int{1}, []int{2}},
		},
		{
			name: "пустая матрица",
			matrix: [][]int{},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TransposeRect(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransposeRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
