package slice

import (
	"reflect"
	"testing"
)

func TestFilterEven(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "смешанные числа",
			src: []int{1, 2, 3, 4},
			want: []int{2, 4},
		},
		{
			name: "чётных нет",
			src: []int{1, 3},
			want: []int{},
		},
		{
			name: "nil слайс",
			src: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterEven(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
