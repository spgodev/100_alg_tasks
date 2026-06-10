package slice

import (
	"reflect"
	"testing"
)

func TestClampSlice(t *testing.T) {
	tests := []struct {
		name string
		src []int
		min int
		max int
		want []int
	}{
		{
			name: "обычный случай",
			src: []int{-5, 0, 5, 10},
			min: 0,
			max: 7,
			want: []int{0, 0, 5, 7},
		},
		{
			name: "все внутри",
			src: []int{1, 2, 3},
			min: 0,
			max: 5,
			want: []int{1, 2, 3},
		},
		{
			name: "пустой слайс",
			src: []int{},
			min: 0,
			max: 1,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClampSlice(tt.src, tt.min, tt.max)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClampSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
