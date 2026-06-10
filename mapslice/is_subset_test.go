package mapslice

import (
	"testing"
)

func TestIsSubset(t *testing.T) {
	tests := []struct {
		name string
		subset []int
		set []int
		want bool
	}{
		{
			name: "является подмножеством",
			subset: []int{1, 2, 2},
			set: []int{2, 1, 3},
			want: true,
		},
		{
			name: "не является",
			subset: []int{1, 4},
			set: []int{1, 2, 3},
			want: false,
		},
		{
			name: "пустой subset",
			subset: []int{},
			set: []int{1},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSubset(tt.subset, tt.set)
			if got != tt.want {
				t.Errorf("IsSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
