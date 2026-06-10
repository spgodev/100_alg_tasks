package slice

import (
	"testing"
)

func TestCountOdd(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "смешанные числа",
			src: []int{1, 2, 3, 4, 5},
			want: 3,
		},
		{
			name: "отрицательные числа",
			src: []int{-3, -2, -1, 0},
			want: 2,
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountOdd(tt.src)
			if got != tt.want {
				t.Errorf("CountOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
