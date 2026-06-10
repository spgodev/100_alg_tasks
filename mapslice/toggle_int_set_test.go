package mapslice

import (
	"reflect"
	"testing"
)

func TestToggleIntSet(t *testing.T) {
	tests := []struct {
		name string
		set map[int]bool
		values []int
		want map[int]bool
	}{
		{
			name: "переключить существующий и новый",
			set: map[int]bool{1: true, 2: true},
			values: []int{2, 3},
			want: map[int]bool{1: true, 3: true},
		},
		{
			name: "повторное переключение",
			set: map[int]bool{},
			values: []int{1, 1},
			want: map[int]bool{},
		},
		{
			name: "nil set",
			set: nil,
			values: []int{5},
			want: map[int]bool{5: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToggleIntSet(tt.set, tt.values)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToggleIntSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
