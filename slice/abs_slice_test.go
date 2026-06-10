package slice

import (
	"reflect"
	"testing"
)

func TestAbsSlice(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "смешанные числа",
			src: []int{-3, 0, 2},
			want: []int{3, 0, 2},
		},
		{
			name: "все положительные",
			src: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AbsSlice(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AbsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
