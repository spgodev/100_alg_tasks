package slice

import (
	"reflect"
	"testing"
)

func TestReverseCopy(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "обычный случай",
			src: []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			name: "один элемент",
			src: []int{5},
			want: []int{5},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseCopy(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
