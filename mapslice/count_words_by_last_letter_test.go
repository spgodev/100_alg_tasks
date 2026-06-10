package mapslice

import (
	"reflect"
	"testing"
)

func TestCountWordsByLastLetter(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want map[string]int
	}{
		{
			name: "обычный случай",
			words: []string{"go", "hello", "rust"},
			want: map[string]int{"o": 2, "t": 1},
		},
		{
			name: "пустая строка",
			words: []string{"", "a", ""},
			want: map[string]int{"": 2, "a": 1},
		},
		{
			name: "nil слайс",
			words: nil,
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWordsByLastLetter(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountWordsByLastLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
