package slice

import (
	"testing"
)

func TestSecondMax(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "обычный случай",
			src: []int{3, 1, 5, 2},
			want: 3,
		},
		{
			name: "с дубликатом максимума",
			src: []int{5, 5, 4, 1},
			want: 4,
		},
		{
			name: "нет второго уникального",
			src: []int{7, 7},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SecondMax(tt.src)
			if got != tt.want {
				t.Errorf("SecondMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
