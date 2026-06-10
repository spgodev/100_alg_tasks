package slice

import (
	"reflect"
	"testing"
)

func TestFilterPositive(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "смешанные числа",
			src: []int{-1, 0, 2, 3},
			want: []int{2, 3},
		},
		{
			name: "положительных нет",
			src: []int{-1, 0},
			want: []int{},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterPositive(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
