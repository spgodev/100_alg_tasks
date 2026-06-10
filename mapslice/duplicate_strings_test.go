package mapslice

import (
	"reflect"
	"testing"
)

func TestDuplicateStrings(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want []string
	}{
		{
			name: "несколько дублей",
			words: []string{"a", "b", "a", "c", "b", "a"},
			want: []string{"a", "b"},
		},
		{
			name: "дублей нет",
			words: []string{"a", "b"},
			want: []string{},
		},
		{
			name: "пустая строка",
			words: []string{"", "", "x"},
			want: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DuplicateStrings(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DuplicateStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
