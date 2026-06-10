package mapslice

import (
	"reflect"
	"testing"
)

func TestBuildIntSet(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want map[int]bool
	}{
		{
			name: "есть повторы",
			src: []int{1, 2, 1},
			want: map[int]bool{1: true, 2: true},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: map[int]bool{},
		},
		{
			name: "отрицательные",
			src: []int{-1, 0},
			want: map[int]bool{-1: true, 0: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildIntSet(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildIntSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
