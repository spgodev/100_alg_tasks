package mapslice

import (
	"reflect"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want []string
	}{
		{
			name: "есть повторы",
			words: []string{"go", "rust", "go", "c"},
			want: []string{"go", "rust", "c"},
		},
		{
			name: "без повторов",
			words: []string{"a", "b"},
			want: []string{"a", "b"},
		},
		{
			name: "nil слайс",
			words: nil,
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueStrings(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
