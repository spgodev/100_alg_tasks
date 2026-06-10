package mapslice

import (
	"reflect"
	"testing"
)

func TestSetUnionInts(t *testing.T) {
	tests := []struct {
		name string
		a []int
		b []int
		want []int
	}{
		{
			name: "есть пересечение",
			a: []int{1, 2, 1},
			b: []int{2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "a пустой",
			a: []int{},
			b: []int{3, 3},
			want: []int{3},
		},
		{
			name: "оба пустые",
			a: []int{},
			b: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetUnionInts(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUnionInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
