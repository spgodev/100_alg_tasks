package mapslice

import (
	"reflect"
	"testing"
)

func TestSetIntersectionStrings(t *testing.T) {
	tests := []struct {
		name string
		a []string
		b []string
		want []string
	}{
		{
			name: "есть пересечение",
			a: []string{"a", "b", "a", "c"},
			b: []string{"c", "a"},
			want: []string{"a", "c"},
		},
		{
			name: "нет пересечения",
			a: []string{"a"},
			b: []string{"b"},
			want: []string{},
		},
		{
			name: "один пустой",
			a: []string{},
			b: []string{"a"},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetIntersectionStrings(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIntersectionStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
