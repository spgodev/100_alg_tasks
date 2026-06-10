package slice

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesSorted(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "есть дубликаты",
			src: []int{1, 1, 2, 2, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "без дубликатов",
			src: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicatesSorted(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicatesSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
