package slice

import (
	"reflect"
	"testing"
)

func TestPadRight(t *testing.T) {
	tests := []struct {
		name string
		src []int
		size int
		value int
		want []int
	}{
		{
			name: "нужно дополнить",
			src: []int{1, 2},
			size: 5,
			value: 0,
			want: []int{1, 2, 0, 0, 0},
		},
		{
			name: "длина уже достаточна",
			src: []int{1, 2, 3},
			size: 2,
			value: 9,
			want: []int{1, 2, 3},
		},
		{
			name: "nil слайс",
			src: nil,
			size: 3,
			value: 7,
			want: []int{7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PadRight(tt.src, tt.size, tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PadRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
