package mapslice

import (
	"reflect"
	"testing"
)

func TestLastIndexMap(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want map[int]int
	}{
		{
			name: "есть повторы",
			src: []int{5, 7, 5, 9},
			want: map[int]int{5: 2, 7: 1, 9: 3},
		},
		{
			name: "без повторов",
			src: []int{1, 2},
			want: map[int]int{1: 0, 2: 1},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: map[int]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LastIndexMap(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastIndexMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
