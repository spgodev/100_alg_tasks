package mapslice

import (
	"testing"
)

func TestCountTrueValues(t *testing.T) {
	tests := []struct {
		name string
		m map[string]bool
		want int
	}{
		{
			name: "смешанные значения",
			m: map[string]bool{"a": true, "b": false, "c": true},
			want: 2,
		},
		{
			name: "true нет",
			m: map[string]bool{"a": false},
			want: 0,
		},
		{
			name: "nil map",
			m: nil,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountTrueValues(tt.m)
			if got != tt.want {
				t.Errorf("CountTrueValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
