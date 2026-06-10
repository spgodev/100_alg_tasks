package slice

import (
	"reflect"
	"testing"
)

func TestRepeatElements(t *testing.T) {
	tests := []struct {
		name string
		src []int
		times int
		want []int
	}{
		{
			name: "повторить два раза",
			src: []int{1, 2, 3},
			times: 2,
			want: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name: "times равен нулю",
			src: []int{1, 2},
			times: 0,
			want: []int{},
		},
		{
			name: "nil слайс",
			src: nil,
			times: 3,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RepeatElements(tt.src, tt.times)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepeatElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
