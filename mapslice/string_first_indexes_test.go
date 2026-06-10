package mapslice

import (
	"reflect"
	"testing"
)

func TestStringFirstIndexes(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want map[string]int
	}{
		{
			name: "есть повторы",
			words: []string{"a", "b", "a"},
			want: map[string]int{"a": 0, "b": 1},
		},
		{
			name: "пустая строка",
			words: []string{"", "x", ""},
			want: map[string]int{"": 0, "x": 1},
		},
		{
			name: "nil слайс",
			words: nil,
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringFirstIndexes(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringFirstIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}
