package mapslice

import (
	"testing"
)

func TestSumMapValues(t *testing.T) {
	tests := []struct {
		name string
		m map[string]int
		want int
	}{
		{
			name: "обычный случай",
			m: map[string]int{"a": 1, "b": 2},
			want: 3,
		},
		{
			name: "есть отрицательные",
			m: map[string]int{"a": 5, "b": -2},
			want: 3,
		},
		{
			name: "nil map",
			m: nil,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumMapValues(tt.m)
			if got != tt.want {
				t.Errorf("SumMapValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
