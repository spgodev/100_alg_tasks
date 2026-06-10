package slice

import (
	"reflect"
	"testing"
)

func TestMergeAlternating(t *testing.T) {
	tests := []struct {
		name string
		a []int
		b []int
		want []int
	}{
		{
			name: "одинаковая длина",
			a: []int{1, 3},
			b: []int{2, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "a длиннее",
			a: []int{1, 3, 5},
			b: []int{2},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "b длиннее",
			a: []int{1},
			b: []int{2, 4, 6},
			want: []int{1, 2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeAlternating(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeAlternating() = %v, want %v", got, tt.want)
			}
		})
	}
}
