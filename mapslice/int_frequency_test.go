package mapslice

import (
	"reflect"
	"testing"
)

func TestIntFrequency(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want map[int]int
	}{
		{
			name: "обычный случай",
			src: []int{1, 2, 1, 3},
			want: map[int]int{1: 2, 2: 1, 3: 1},
		},
		{
			name: "отрицательные",
			src: []int{-1, -1, 0},
			want: map[int]int{-1: 2, 0: 1},
		},
		{
			name: "nil слайс",
			src: nil,
			want: map[int]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntFrequency(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
