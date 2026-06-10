package slice

import (
	"testing"
)

func TestAnyNegative(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want bool
	}{
		{
			name: "есть отрицательное",
			src: []int{1, -2, 3},
			want: true,
		},
		{
			name: "нет отрицательных",
			src: []int{0, 2, 3},
			want: false,
		},
		{
			name: "nil слайс",
			src: nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnyNegative(tt.src)
			if got != tt.want {
				t.Errorf("AnyNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}
