package slice

import (
	"reflect"
	"testing"
)

func TestUniqueInts(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "есть повторы",
			src: []int{1, 2, 1, 3, 2},
			want: []int{1, 2, 3},
		},
		{
			name: "без повторов",
			src: []int{4, 5},
			want: []int{4, 5},
		},
		{
			name: "nil слайс",
			src: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueInts(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
