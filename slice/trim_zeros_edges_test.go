package slice

import (
	"reflect"
	"testing"
)

func TestTrimZerosEdges(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "нули по краям",
			src: []int{0, 0, 1, 0, 2, 0},
			want: []int{1, 0, 2},
		},
		{
			name: "нулей по краям нет",
			src: []int{1, 0, 2},
			want: []int{1, 0, 2},
		},
		{
			name: "все нули",
			src: []int{0, 0},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrimZerosEdges(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimZerosEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}
