package mapslice

import (
	"reflect"
	"testing"
)

func TestGroupByParity(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want map[string][]int
	}{
		{
			name: "смешанные числа",
			src: []int{1, 2, 3, 4},
			want: map[string][]int{"even": []int{2, 4}, "odd": []int{1, 3}},
		},
		{
			name: "только чётные",
			src: []int{2, 4},
			want: map[string][]int{"even": []int{2, 4}, "odd": []int{}},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: map[string][]int{"even": []int{}, "odd": []int{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupByParity(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
