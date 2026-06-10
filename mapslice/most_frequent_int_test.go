package mapslice

import (
	"testing"
)

func TestMostFrequentInt(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "есть лидер",
			src: []int{1, 2, 2, 3},
			want: 2,
		},
		{
			name: "ничья по частоте",
			src: []int{5, 1, 5, 1},
			want: 1,
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MostFrequentInt(tt.src)
			if got != tt.want {
				t.Errorf("MostFrequentInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
