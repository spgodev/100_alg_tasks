package slice

import (
	"testing"
)

func TestMaxValue(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "обычный случай",
			src: []int{3, 1, 7, 2},
			want: 7,
		},
		{
			name: "все отрицательные",
			src: []int{-5, -2, -9},
			want: -2,
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxValue(tt.src)
			if got != tt.want {
				t.Errorf("MaxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
