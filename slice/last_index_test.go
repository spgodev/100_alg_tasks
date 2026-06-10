package slice

import (
	"testing"
)

func TestLastIndex(t *testing.T) {
	tests := []struct {
		name string
		src []int
		target int
		want int
	}{
		{
			name: "несколько вхождений",
			src: []int{1, 2, 3, 2},
			target: 2,
			want: 3,
		},
		{
			name: "одно вхождение",
			src: []int{5, 6, 7},
			target: 5,
			want: 0,
		},
		{
			name: "не найдено",
			src: []int{1, 2},
			target: 9,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LastIndex(tt.src, tt.target)
			if got != tt.want {
				t.Errorf("LastIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
