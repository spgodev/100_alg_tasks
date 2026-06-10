package slice

import (
	"reflect"
	"testing"
)

func TestSliceWithoutNegatives(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "смешанные числа",
			src: []int{-1, 0, 2, -3, 4},
			want: []int{0, 2, 4},
		},
		{
			name: "все отрицательные",
			src: []int{-1, -2},
			want: []int{},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceWithoutNegatives(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceWithoutNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
