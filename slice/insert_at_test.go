package slice

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	tests := []struct {
		name string
		src []int
		index int
		value int
		want []int
	}{
		{
			name: "вставка в середину",
			src: []int{1, 2, 4},
			index: 2,
			value: 3,
			want: []int{1, 2, 3, 4},
		},
		{
			name: "вставка в начало",
			src: []int{2, 3},
			index: 0,
			value: 1,
			want: []int{1, 2, 3},
		},
		{
			name: "индекс вне диапазона",
			src: []int{1, 2},
			index: 5,
			value: 9,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertAt(tt.src, tt.index, tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
