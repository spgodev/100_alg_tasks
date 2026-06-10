package slice

import (
	"testing"
)

func TestCountLessThan(t *testing.T) {
	tests := []struct {
		name string
		src []int
		limit int
		want int
	}{
		{
			name: "обычный случай",
			src: []int{1, 5, 7, 2},
			limit: 5,
			want: 2,
		},
		{
			name: "с отрицательными",
			src: []int{-3, 0, 4},
			limit: 0,
			want: 1,
		},
		{
			name: "пустой слайс",
			src: []int{},
			limit: 10,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountLessThan(tt.src, tt.limit)
			if got != tt.want {
				t.Errorf("CountLessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}
