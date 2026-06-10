package slice

import (
	"testing"
)

func TestContainsValue(t *testing.T) {
	tests := []struct {
		name string
		src []int
		target int
		want bool
	}{
		{
			name: "значение есть",
			src: []int{1, 2, 3},
			target: 2,
			want: true,
		},
		{
			name: "значения нет",
			src: []int{1, 2, 3},
			target: 5,
			want: false,
		},
		{
			name: "nil слайс",
			src: nil,
			target: 1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsValue(tt.src, tt.target)
			if got != tt.want {
				t.Errorf("ContainsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
