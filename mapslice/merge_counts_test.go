package mapslice

import (
	"reflect"
	"testing"
)

func TestMergeCounts(t *testing.T) {
	tests := []struct {
		name string
		a map[string]int
		b map[string]int
		want map[string]int
	}{
		{
			name: "есть общий ключ",
			a: map[string]int{"go": 2, "js": 1},
			b: map[string]int{"go": 3, "py": 4},
			want: map[string]int{"go": 5, "js": 1, "py": 4},
		},
		{
			name: "первая nil",
			a: nil,
			b: map[string]int{"a": 1},
			want: map[string]int{"a": 1},
		},
		{
			name: "обе пустые",
			a: map[string]int{},
			b: map[string]int{},
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeCounts(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
