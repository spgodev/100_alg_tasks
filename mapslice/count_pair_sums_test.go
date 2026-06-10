package mapslice

import (
	"testing"
)

func TestCountPairSums(t *testing.T) {
	tests := []struct {
		name string
		src []int
		target int
		want int
	}{
		{
			name: "несколько пар",
			src: []int{1, 2, 3, 4, 5},
			target: 6,
			want: 2,
		},
		{
			name: "одинаковые числа",
			src: []int{3, 3, 3},
			target: 6,
			want: 3,
		},
		{
			name: "пар нет",
			src: []int{1, 2},
			target: 10,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountPairSums(tt.src, tt.target)
			if got != tt.want {
				t.Errorf("CountPairSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
