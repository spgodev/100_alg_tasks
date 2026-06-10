package slice

import (
	"reflect"
	"testing"
)

func TestRunningMax(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "обычный случай",
			src: []int{2, 1, 5, 3},
			want: []int{2, 2, 5, 5},
		},
		{
			name: "убывание",
			src: []int{5, 4, 3},
			want: []int{5, 5, 5},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RunningMax(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunningMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
