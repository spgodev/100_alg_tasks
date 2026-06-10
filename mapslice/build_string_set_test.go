package mapslice

import (
	"reflect"
	"testing"
)

func TestBuildStringSet(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want map[string]bool
	}{
		{
			name: "есть повторы",
			words: []string{"a", "b", "a"},
			want: map[string]bool{"a": true, "b": true},
		},
		{
			name: "пустая строка",
			words: []string{"", "x"},
			want: map[string]bool{"": true, "x": true},
		},
		{
			name: "nil слайс",
			words: nil,
			want: map[string]bool{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildStringSet(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildStringSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
