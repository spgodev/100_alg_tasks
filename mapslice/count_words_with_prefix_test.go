package mapslice

import (
	"testing"
)

func TestCountWordsWithPrefix(t *testing.T) {
	tests := []struct {
		name string
		words []string
		prefix string
		want int
	}{
		{
			name: "обычный случай",
			words: []string{"go", "git", "rust"},
			prefix: "g",
			want: 2,
		},
		{
			name: "пустой prefix",
			words: []string{"a", ""},
			prefix: "",
			want: 2,
		},
		{
			name: "nil слайс",
			words: nil,
			prefix: "x",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWordsWithPrefix(tt.words, tt.prefix)
			if got != tt.want {
				t.Errorf("CountWordsWithPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
