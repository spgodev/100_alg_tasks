package mapslice

import (
	"reflect"
	"testing"
)

func TestRenameMapKey(t *testing.T) {
	tests := []struct {
		name string
		m map[string]int
		oldKey string
		newKey string
		want map[string]int
	}{
		{
			name: "переименование",
			m: map[string]int{"a": 1, "b": 2},
			oldKey: "a",
			newKey: "c",
			want: map[string]int{"b": 2, "c": 1},
		},
		{
			name: "старого ключа нет",
			m: map[string]int{"a": 1},
			oldKey: "x",
			newKey: "y",
			want: map[string]int{"a": 1},
		},
		{
			name: "newKey уже есть",
			m: map[string]int{"a": 1, "b": 2},
			oldKey: "a",
			newKey: "b",
			want: map[string]int{"b": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RenameMapKey(tt.m, tt.oldKey, tt.newKey)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenameMapKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
