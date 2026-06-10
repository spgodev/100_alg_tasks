package mapslice

import (
	"reflect"
	"testing"
)

func TestKeysWithValue(t *testing.T) {
	tests := []struct {
		name string
		m map[string]int
		value int
		want []string
	}{
		{
			name: "несколько ключей",
			m: map[string]int{"b": 2, "a": 2, "c": 3},
			value: 2,
			want: []string{"a", "b"},
		},
		{
			name: "нет ключей",
			m: map[string]int{"a": 1},
			value: 2,
			want: []string{},
		},
		{
			name: "nil map",
			m: nil,
			value: 1,
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeysWithValue(tt.m, tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeysWithValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
