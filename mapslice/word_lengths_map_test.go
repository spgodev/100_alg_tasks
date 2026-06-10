package mapslice

import (
	"reflect"
	"testing"
)

func TestWordLengthsMap(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want map[string]int
	}{
		{
			name: "обычный случай",
			words: []string{"go", "rust"},
			want: map[string]int{"go": 2, "rust": 4},
		},
		{
			name: "есть повтор",
			words: []string{"a", "bb", "a"},
			want: map[string]int{"a": 1, "bb": 2},
		},
		{
			name: "nil слайс",
			words: nil,
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordLengthsMap(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordLengthsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
