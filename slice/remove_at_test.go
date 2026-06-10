package slice

import (
	"reflect"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	tests := []struct {
		name string
		src []int
		index int
		want []int
	}{
		{
			name: "удалить из середины",
			src: []int{1, 2, 3, 4},
			index: 1,
			want: []int{1, 3, 4},
		},
		{
			name: "удалить последний",
			src: []int{1, 2},
			index: 1,
			want: []int{1},
		},
		{
			name: "индекс вне диапазона",
			src: []int{1, 2},
			index: -1,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveAt(tt.src, tt.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
