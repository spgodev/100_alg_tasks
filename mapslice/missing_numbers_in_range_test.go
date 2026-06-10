package mapslice

import (
	"reflect"
	"testing"
)

func TestMissingNumbersInRange(t *testing.T) {
	tests := []struct {
		name string
		src []int
		left int
		right int
		want []int
	}{
		{
			name: "есть пропуски",
			src: []int{1, 3, 5},
			left: 1,
			right: 5,
			want: []int{2, 4},
		},
		{
			name: "пропусков нет",
			src: []int{1, 2, 3},
			left: 1,
			right: 3,
			want: []int{},
		},
		{
			name: "некорректный диапазон",
			src: []int{1},
			left: 5,
			right: 1,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MissingNumbersInRange(tt.src, tt.left, tt.right)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MissingNumbersInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
