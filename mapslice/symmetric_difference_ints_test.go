package mapslice

import (
	"reflect"
	"testing"
)

func TestSymmetricDifferenceInts(t *testing.T) {
	tests := []struct {
		name string
		a []int
		b []int
		want []int
	}{
		{
			name: "обычный случай",
			a: []int{1, 2, 3},
			b: []int{3, 4, 5},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "всё общее",
			a: []int{1, 2},
			b: []int{2, 1},
			want: []int{},
		},
		{
			name: "a пустой",
			a: []int{},
			b: []int{1, 1},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SymmetricDifferenceInts(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SymmetricDifferenceInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
