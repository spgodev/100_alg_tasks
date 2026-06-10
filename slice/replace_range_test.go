package slice

import (
	"reflect"
	"testing"
)

func TestReplaceRange(t *testing.T) {
	tests := []struct {
		name string
		src []int
		from int
		to int
		value int
		want []int
	}{
		{
			name: "замена диапазона",
			src: []int{1, 2, 3, 4},
			from: 1,
			to: 3,
			value: 9,
			want: []int{1, 9, 9, 4},
		},
		{
			name: "пустой диапазон",
			src: []int{1, 2},
			from: 1,
			to: 1,
			value: 0,
			want: []int{1, 2},
		},
		{
			name: "диапазон вне границ",
			src: []int{1, 2},
			from: -1,
			to: 2,
			value: 0,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceRange(tt.src, tt.from, tt.to, tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReplaceRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
