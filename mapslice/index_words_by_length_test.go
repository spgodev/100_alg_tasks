package mapslice

import (
	"reflect"
	"testing"
)

func TestIndexWordsByLength(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want map[int][]int
	}{
		{
			name: "несколько длин",
			words: []string{"a", "bb", "c"},
			want: map[int][]int{1: []int{0, 2}, 2: []int{1}},
		},
		{
			name: "пустая строка",
			words: []string{"", "x"},
			want: map[int][]int{0: []int{0}, 1: []int{1}},
		},
		{
			name: "пустой слайс",
			words: []string{},
			want: map[int][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IndexWordsByLength(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexWordsByLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
