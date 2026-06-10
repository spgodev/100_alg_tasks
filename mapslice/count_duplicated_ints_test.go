package mapslice

import (
	"testing"
)

func TestCountDuplicatedInts(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want int
	}{
		{
			name: "два значения с дублями",
			src: []int{1, 2, 1, 3, 2, 2},
			want: 2,
		},
		{
			name: "дублей нет",
			src: []int{1, 2, 3},
			want: 0,
		},
		{
			name: "все одинаковые",
			src: []int{5, 5, 5},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountDuplicatedInts(tt.src)
			if got != tt.want {
				t.Errorf("CountDuplicatedInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
