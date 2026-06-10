package mapslice

import (
	"reflect"
	"testing"
)

func TestOnlyKeys(t *testing.T) {
	tests := []struct {
		name string
		m map[string]int
		keys []string
		want map[string]int
	}{
		{
			name: "часть ключей есть",
			m: map[string]int{"a": 1, "b": 2, "c": 3},
			keys: []string{"c", "a"},
			want: map[string]int{"a": 1, "c": 3},
		},
		{
			name: "ключей нет",
			m: map[string]int{"a": 1},
			keys: []string{"x"},
			want: map[string]int{},
		},
		{
			name: "nil keys",
			m: map[string]int{"a": 1},
			keys: nil,
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OnlyKeys(tt.m, tt.keys)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnlyKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
