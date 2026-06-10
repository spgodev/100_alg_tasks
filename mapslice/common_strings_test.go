package mapslice

import (
	"reflect"
	"testing"
)

func TestCommonStrings(t *testing.T) {
	tests := []struct {
		name string
		a []string
		b []string
		want []string
	}{
		{
			name: "есть общие",
			a: []string{"go", "rust", "go", "c"},
			b: []string{"c", "go"},
			want: []string{"go", "c"},
		},
		{
			name: "общих нет",
			a: []string{"a"},
			b: []string{"b"},
			want: []string{},
		},
		{
			name: "один nil",
			a: nil,
			b: []string{"a"},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CommonStrings(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
