package slice

import (
	"reflect"
	"testing"
)

func TestRemoveAll(t *testing.T) {
	tests := []struct {
		name string
		src []int
		target int
		want []int
	}{
		{
			name: "удалить несколько",
			src: []int{1, 2, 3, 2, 4},
			target: 2,
			want: []int{1, 3, 4},
		},
		{
			name: "ничего не удаляется",
			src: []int{1, 3},
			target: 2,
			want: []int{1, 3},
		},
		{
			name: "все удаляются",
			src: []int{5, 5},
			target: 5,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveAll(tt.src, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
