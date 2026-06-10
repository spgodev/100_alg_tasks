package slice

import (
	"testing"
)

func TestSumNonNegative(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "смешанные числа",
			src: []int{-2, 0, 3, 5},
			want: 8,
		},
		{
			name: "только отрицательные",
			src: []int{-1, -2},
			want: 0,
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumNonNegative(tt.src)
			if got != tt.want {
				t.Errorf("SumNonNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}
