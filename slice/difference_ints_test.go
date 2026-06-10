package slice

import (
	"reflect"
	"testing"
)

func TestDifferenceInts(t *testing.T) {
	tests := []struct {
		name string
		a []int
		b []int
		want []int
	}{
		{
			name: "обычный случай",
			a: []int{1, 2, 3, 2},
			b: []int{2},
			want: []int{1, 3},
		},
		{
			name: "ничего не удаляется",
			a: []int{1, 3},
			b: []int{2},
			want: []int{1, 3},
		},
		{
			name: "a пустой",
			a: []int{},
			b: []int{1},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInts(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DifferenceInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
