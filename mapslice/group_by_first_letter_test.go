package mapslice

import (
	"reflect"
	"testing"
)

func TestGroupByFirstLetter(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want map[string][]string
	}{
		{
			name: "обычный случай",
			words: []string{"go", "git", "rust"},
			want: map[string][]string{"g": []string{"go", "git"}, "r": []string{"rust"}},
		},
		{
			name: "есть пустая строка",
			words: []string{"", "a"},
			want: map[string][]string{"": []string{""}, "a": []string{"a"}},
		},
		{
			name: "nil слайс",
			words: nil,
			want: map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupByFirstLetter(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByFirstLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
