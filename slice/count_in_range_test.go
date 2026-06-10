package slice

import (
	"testing"
)

func TestCountInRange(t *testing.T) {
	tests := []struct {
		name string
		src []int
		left int
		right int
		want int
	}{
		{
			name: "обычный случай",
			src: []int{1, 2, 3, 4, 5},
			left: 2,
			right: 4,
			want: 3,
		},
		{
			name: "левая граница больше правой",
			src: []int{1, 2, 3},
			left: 5,
			right: 1,
			want: 0,
		},
		{
			name: "нет подходящих",
			src: []int{10, 20},
			left: 1,
			right: 5,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountInRange(tt.src, tt.left, tt.right)
			if got != tt.want {
				t.Errorf("CountInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
