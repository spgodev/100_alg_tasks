package slice

import (
	"testing"
)

func TestFirstPositiveIndex(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "первое положительное в середине",
			src: []int{-2, 0, 5, 7},
			want: 2,
		},
		{
			name: "первый элемент",
			src: []int{1, -1},
			want: 0,
		},
		{
			name: "нет положительных",
			src: []int{-1, 0, -3},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstPositiveIndex(tt.src)
			if got != tt.want {
				t.Errorf("FirstPositiveIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
