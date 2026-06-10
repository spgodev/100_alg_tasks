package mapslice

import (
	"testing"
)

func TestAreDisjoint(t *testing.T) {
	tests := []struct {
		name string
		a []int
		b []int
		want bool
	}{
		{
			name: "нет общих",
			a: []int{1, 2},
			b: []int{3, 4},
			want: true,
		},
		{
			name: "есть общее",
			a: []int{1, 2},
			b: []int{2, 3},
			want: false,
		},
		{
			name: "один пустой",
			a: []int{},
			b: []int{1},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AreDisjoint(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("AreDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
