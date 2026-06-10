package slice

import (
	"reflect"
	"testing"
)

func TestIntersectInts(t *testing.T) {
	tests := []struct {
		name string
		a []int
		b []int
		want []int
	}{
		{
			name: "есть пересечение",
			a: []int{1, 2, 3, 2},
			b: []int{2, 4, 3},
			want: []int{2, 3},
		},
		{
			name: "пересечения нет",
			a: []int{1, 2},
			b: []int{3, 4},
			want: []int{},
		},
		{
			name: "один слайс nil",
			a: nil,
			b: []int{1},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectInts(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
