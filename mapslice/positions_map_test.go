package mapslice

import (
	"reflect"
	"testing"
)

func TestPositionsMap(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want map[int][]int
	}{
		{
			name: "есть повторы",
			src: []int{1, 2, 1, 3, 2},
			want: map[int][]int{1: []int{0, 2}, 2: []int{1, 4}, 3: []int{3}},
		},
		{
			name: "одно значение",
			src: []int{5, 5},
			want: map[int][]int{5: []int{0, 1}},
		},
		{
			name: "nil слайс",
			src: nil,
			want: map[int][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PositionsMap(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PositionsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
