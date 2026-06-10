package slice

import (
	"reflect"
	"testing"
)

func TestFilterOdd(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "смешанные числа",
			src: []int{1, 2, 3, 4, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "нечётных нет",
			src: []int{2, 4},
			want: []int{},
		},
		{
			name: "отрицательные",
			src: []int{-3, -2, -1},
			want: []int{-3, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterOdd(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
