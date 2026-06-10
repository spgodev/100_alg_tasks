package mapslice

import (
	"testing"
)

func TestHasKey(t *testing.T) {
	tests := []struct {
		name string
		m map[string]int
		key string
		want bool
	}{
		{
			name: "ключ есть",
			m: map[string]int{"go": 1},
			key: "go",
			want: true,
		},
		{
			name: "ключа нет",
			m: map[string]int{"go": 1},
			key: "rust",
			want: false,
		},
		{
			name: "nil map",
			m: nil,
			key: "x",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasKey(tt.m, tt.key)
			if got != tt.want {
				t.Errorf("HasKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
