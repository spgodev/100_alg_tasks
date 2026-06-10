package mapslice

import (
	"reflect"
	"testing"
)

func TestFirstIndexMap(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want map[int]int
	}{
		{
			name: "есть повторы",
			src: []int{5, 7, 5, 9},
			want: map[int]int{5: 0, 7: 1, 9: 3},
		},
		{
			name: "без повторов",
			src: []int{1, 2},
			want: map[int]int{1: 0, 2: 1},
		},
		{
			name: "nil слайс",
			src: nil,
			want: map[int]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstIndexMap(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstIndexMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
