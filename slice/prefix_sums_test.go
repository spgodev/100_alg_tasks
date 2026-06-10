package slice

import (
	"reflect"
	"testing"
)

func TestPrefixSums(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "обычный случай",
			src: []int{1, 2, 3, 4},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "с отрицательными",
			src: []int{2, -1, 3},
			want: []int{2, 1, 4},
		},
		{
			name: "nil слайс",
			src: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrefixSums(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrefixSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
