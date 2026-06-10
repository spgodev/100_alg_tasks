package mapslice

import (
	"reflect"
	"testing"
)

func TestStringFrequency(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want map[string]int
	}{
		{
			name: "обычный случай",
			words: []string{"a", "b", "a"},
			want: map[string]int{"a": 2, "b": 1},
		},
		{
			name: "регистр важен",
			words: []string{"Go", "go"},
			want: map[string]int{"Go": 1, "go": 1},
		},
		{
			name: "пустой слайс",
			words: []string{},
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringFrequency(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
