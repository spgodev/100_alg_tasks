package mapslice

import (
	"reflect"
	"testing"
)

func TestMapFromPairs(t *testing.T) {
	tests := []struct {
		name string
		keys []string
		values []int
		want map[string]int
	}{
		{
			name: "одинаковая длина",
			keys: []string{"a", "b"},
			values: []int{1, 2},
			want: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "values короче",
			keys: []string{"a", "b", "c"},
			values: []int{1},
			want: map[string]int{"a": 1},
		},
		{
			name: "повтор ключа",
			keys: []string{"a", "a"},
			values: []int{1, 2},
			want: map[string]int{"a": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapFromPairs(tt.keys, tt.values)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapFromPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
