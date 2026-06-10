package mapslice

import (
	"testing"
)

func TestCountUniqueInts(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "есть повторы",
			src: []int{1, 2, 1, 3},
			want: 3,
		},
		{
			name: "все одинаковые",
			src: []int{5, 5, 5},
			want: 1,
		},
		{
			name: "nil слайс",
			src: nil,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountUniqueInts(tt.src)
			if got != tt.want {
				t.Errorf("CountUniqueInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
