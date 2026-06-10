package slice

import (
	"testing"
)

func TestCountEven(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "смешанные числа",
			src: []int{1, 2, 3, 4, 6},
			want: 3,
		},
		{
			name: "чётных нет",
			src: []int{1, 3, 5},
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
			got := CountEven(tt.src)
			if got != tt.want {
				t.Errorf("CountEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
