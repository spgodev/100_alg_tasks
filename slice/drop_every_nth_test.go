package slice

import (
	"reflect"
	"testing"
)

func TestDropEveryNth(t *testing.T) {
	tests := []struct {
		name string
		src []int
		n int
		want []int
	}{
		{
			name: "убрать каждый второй",
			src: []int{1, 2, 3, 4, 5},
			n: 2,
			want: []int{1, 3, 5},
		},
		{
			name: "n больше длины",
			src: []int{1, 2},
			n: 5,
			want: []int{1, 2},
		},
		{
			name: "n отрицательный",
			src: []int{1, 2},
			n: -1,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DropEveryNth(tt.src, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropEveryNth() = %v, want %v", got, tt.want)
			}
		})
	}
}
