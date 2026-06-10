package slice

import (
	"testing"
)

func TestSumNegative(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "есть отрицательные",
			src: []int{-1, 2, -3, 4},
			want: -4,
		},
		{
			name: "отрицательных нет",
			src: []int{1, 2, 3},
			want: 0,
		},
		{
			name: "nil слайс",
			src: nil,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumNegative(tt.src)
			if got != tt.want {
				t.Errorf("SumNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}
