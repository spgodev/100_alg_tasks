package slice

import (
	"testing"
)

func TestIsSortedDesc(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want bool
	}{
		{
			name: "отсортирован",
			src: []int{5, 4, 4, 1},
			want: true,
		},
		{
			name: "не отсортирован",
			src: []int{5, 3, 4},
			want: false,
		},
		{
			name: "один элемент",
			src: []int{7},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSortedDesc(tt.src)
			if got != tt.want {
				t.Errorf("IsSortedDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}
