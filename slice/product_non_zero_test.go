package slice

import (
	"testing"
)

func TestProductNonZero(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "есть нули",
			src: []int{2, 0, 3, 0, 4},
			want: 24,
		},
		{
			name: "все нули",
			src: []int{0, 0},
			want: 0,
		},
		{
			name: "отрицательное значение",
			src: []int{-2, 3},
			want: -6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProductNonZero(tt.src)
			if got != tt.want {
				t.Errorf("ProductNonZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
