package slice

import (
	"testing"
)

func TestAllPositive(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want bool
	}{
		{
			name: "все положительные",
			src: []int{1, 2, 3},
			want: true,
		},
		{
			name: "есть ноль",
			src: []int{1, 0, 3},
			want: false,
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllPositive(tt.src)
			if got != tt.want {
				t.Errorf("AllPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
