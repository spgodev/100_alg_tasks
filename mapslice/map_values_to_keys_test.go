package mapslice

import (
	"reflect"
	"testing"
)

func TestMapValuesToKeys(t *testing.T) {
	tests := []struct {
		name string
		m map[string]int
		want map[int][]string
	}{
		{
			name: "несколько ключей на значение",
			m: map[string]int{"b": 1, "a": 1, "c": 2},
			want: map[int][]string{1: []string{"a", "b"}, 2: []string{"c"}},
		},
		{
			name: "пустая map",
			m: map[string]int{},
			want: map[int][]string{},
		},
		{
			name: "nil map",
			m: nil,
			want: map[int][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapValuesToKeys(tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapValuesToKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
