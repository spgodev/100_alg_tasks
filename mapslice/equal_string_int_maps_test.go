package mapslice

import (
	"testing"
)

func TestEqualStringIntMaps(t *testing.T) {
	tests := []struct {
		name string
		a map[string]int
		b map[string]int
		want bool
	}{
		{
			name: "равные map",
			a: map[string]int{"a": 1, "b": 2},
			b: map[string]int{"b": 2, "a": 1},
			want: true,
		},
		{
			name: "разные значения",
			a: map[string]int{"a": 1},
			b: map[string]int{"a": 2},
			want: false,
		},
		{
			name: "nil и пустая равны",
			a: nil,
			b: map[string]int{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EqualStringIntMaps(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("EqualStringIntMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
