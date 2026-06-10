package slice

import (
	"reflect"
	"testing"
)

func TestIndexesOf(t *testing.T) {
	tests := []struct {
		name string
		src []int
		target int
		want []int
	}{
		{
			name: "несколько вхождений",
			src: []int{1, 2, 1, 3, 1},
			target: 1,
			want: []int{0, 2, 4},
		},
		{
			name: "нет вхождений",
			src: []int{1, 2},
			target: 3,
			want: []int{},
		},
		{
			name: "nil слайс",
			src: nil,
			target: 0,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IndexesOf(tt.src, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexesOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
