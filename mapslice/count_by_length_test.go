package mapslice

import (
	"reflect"
	"testing"
)

func TestCountByLength(t *testing.T) {
	tests := []struct {
		name string
		words []string
		want map[int]int
	}{
		{
			name: "несколько длин",
			words: []string{"a", "bb", "c", "dd"},
			want: map[int]int{1: 2, 2: 2},
		},
		{
			name: "пустые строки",
			words: []string{"", "x", ""},
			want: map[int]int{0: 2, 1: 1},
		},
		{
			name: "пустой слайс",
			words: []string{},
			want: map[int]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountByLength(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountByLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
