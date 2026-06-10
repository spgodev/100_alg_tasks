package slice

import (
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		src []int
		k int
		want []int
	}{
		{
			name: "сдвиг на один",
			src: []int{1, 2, 3, 4},
			k: 1,
			want: []int{4, 1, 2, 3},
		},
		{
			name: "k больше длины",
			src: []int{1, 2, 3},
			k: 5,
			want: []int{2, 3, 1},
		},
		{
			name: "nil слайс",
			src: nil,
			k: 2,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RotateRight(tt.src, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
