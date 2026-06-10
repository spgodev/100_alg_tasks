package slice

import (
	"testing"
)

func TestCountGreaterThan(t *testing.T) {
	tests := []struct {
		name string
		src []int
		limit int
		want int
	}{
		{
			name: "обычный случай",
			src: []int{1, 5, 7, 2},
			limit: 3,
			want: 2,
		},
		{
			name: "таких нет",
			src: []int{1, 2},
			limit: 5,
			want: 0,
		},
		{
			name: "nil слайс",
			src: nil,
			limit: 0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountGreaterThan(tt.src, tt.limit)
			if got != tt.want {
				t.Errorf("CountGreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}
