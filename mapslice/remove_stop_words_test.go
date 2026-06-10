package mapslice

import (
	"reflect"
	"testing"
)

func TestRemoveStopWords(t *testing.T) {
	tests := []struct {
		name string
		words []string
		stop map[string]bool
		want []string
	}{
		{
			name: "есть стоп-слова",
			words: []string{"go", "and", "rust"},
			stop: map[string]bool{"and": true},
			want: []string{"go", "rust"},
		},
		{
			name: "false не удаляет",
			words: []string{"a", "b"},
			stop: map[string]bool{"b": false},
			want: []string{"a", "b"},
		},
		{
			name: "nil stop",
			words: []string{"a"},
			stop: nil,
			want: []string{"a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveStopWords(tt.words, tt.stop)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveStopWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
