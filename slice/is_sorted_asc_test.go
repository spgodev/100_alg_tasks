package slice

import (
	"testing"
)

func TestIsSortedAsc(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want bool
	}{
		{
			name: "отсортирован",
			src: []int{1, 2, 2, 3},
			want: true,
		},
		{
			name: "не отсортирован",
			src: []int{1, 3, 2},
			want: false,
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSortedAsc(tt.src)
			if got != tt.want {
				t.Errorf("IsSortedAsc() = %v, want %v", got, tt.want)
			}
		})
	}
}
