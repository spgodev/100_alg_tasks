package slice

import (
	"reflect"
	"testing"
)

func TestDoubleSlice(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "обычный случай",
			src: []int{1, 2, 3},
			want: []int{2, 4, 6},
		},
		{
			name: "с отрицательными",
			src: []int{-1, 0, 2},
			want: []int{-2, 0, 4},
		},
		{
			name: "nil слайс",
			src: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DoubleSlice(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
