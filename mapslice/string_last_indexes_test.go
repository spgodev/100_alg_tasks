package mapslice

import (
	"reflect"
	"testing"
)

func TestStringLastIndexes(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want map[string]int
	}{
		{
			name: "есть повторы",
			words: []string{"a", "b", "a"},
			want: map[string]int{"a": 2, "b": 1},
		},
		{
			name: "пустая строка",
			words: []string{"", "x", ""},
			want: map[string]int{"": 2, "x": 1},
		},
		{
			name: "пустой слайс",
			words: []string{},
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringLastIndexes(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringLastIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}
