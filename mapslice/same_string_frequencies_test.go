package mapslice

import (
	"testing"
)

func TestSameStringFrequencies(t *testing.T) {
	tests := []struct {
		name string
		a []string
		b []string
		want bool
	}{
		{
			name: "одинаковые частоты",
			a: []string{"a", "b", "a"},
			b: []string{"b", "a", "a"},
			want: true,
		},
		{
			name: "разные частоты",
			a: []string{"a", "b"},
			b: []string{"a", "a"},
			want: false,
		},
		{
			name: "оба пустые",
			a: []string{},
			b: []string{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SameStringFrequencies(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("SameStringFrequencies() = %v, want %v", got, tt.want)
			}
		})
	}
}
