package slice

import (
	"reflect"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		name string
		src []int
		k int
		want []int
	}{
		{
			name: "сдвиг на два",
			src: []int{1, 2, 3, 4, 5},
			k: 2,
			want: []int{3, 4, 5, 1, 2},
		},
		{
			name: "k больше длины",
			src: []int{1, 2, 3},
			k: 4,
			want: []int{2, 3, 1},
		},
		{
			name: "пустой слайс",
			src: []int{},
			k: 3,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RotateLeft(tt.src, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
