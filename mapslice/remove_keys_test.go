package mapslice

import (
	"reflect"
	"testing"
)

func TestRemoveKeys(t *testing.T) {
	tests := []struct {
		name string
		m map[string]int
		keys []string
		want map[string]int
	}{
		{
			name: "удалить часть",
			m: map[string]int{"a": 1, "b": 2, "c": 3},
			keys: []string{"b"},
			want: map[string]int{"a": 1, "c": 3},
		},
		{
			name: "ключа нет",
			m: map[string]int{"a": 1},
			keys: []string{"x"},
			want: map[string]int{"a": 1},
		},
		{
			name: "nil map",
			m: nil,
			keys: []string{"a"},
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveKeys(tt.m, tt.keys)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
