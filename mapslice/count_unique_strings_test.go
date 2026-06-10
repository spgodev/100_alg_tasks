package mapslice

import (
	"testing"
)

func TestCountUniqueStrings(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want int
	}{
		{
			name: "есть повторы",
			words: []string{"a", "b", "a"},
			want: 2,
		},
		{
			name: "разный регистр",
			words: []string{"Go", "go"},
			want: 2,
		},
		{
			name: "пустой слайс",
			words: []string{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountUniqueStrings(tt.words)
			if got != tt.want {
				t.Errorf("CountUniqueStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
