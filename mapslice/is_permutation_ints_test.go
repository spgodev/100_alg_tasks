package mapslice

import (
	"testing"
)

func TestIsPermutationInts(t *testing.T) {
	tests := []struct {
		name string
		a []int
		b []int
		want bool
	}{
		{
			name: "перестановка",
			a: []int{1, 2, 2},
			b: []int{2, 1, 2},
			want: true,
		},
		{
			name: "разная длина",
			a: []int{1, 2},
			b: []int{1},
			want: false,
		},
		{
			name: "разные частоты",
			a: []int{1, 2, 2},
			b: []int{1, 1, 2},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPermutationInts(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("IsPermutationInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
