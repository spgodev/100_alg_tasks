package slice

import (
	"reflect"
	"testing"
)

func TestMoveZerosEnd(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "нули внутри",
			src: []int{0, 1, 0, 2, 3},
			want: []int{1, 2, 3, 0, 0},
		},
		{
			name: "нулей нет",
			src: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "все нули",
			src: []int{0, 0},
			want: []int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MoveZerosEnd(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZerosEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
