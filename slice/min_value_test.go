package slice

import (
	"testing"
)

func TestMinValue(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "обычный случай",
			src: []int{3, 1, 7, 2},
			want: 1,
		},
		{
			name: "все отрицательные",
			src: []int{-5, -2, -9},
			want: -9,
		},
		{
			name: "nil слайс",
			src: nil,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinValue(tt.src)
			if got != tt.want {
				t.Errorf("MinValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
