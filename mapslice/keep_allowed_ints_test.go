package mapslice

import (
	"reflect"
	"testing"
)

func TestKeepAllowedInts(t *testing.T) {
	tests := []struct {
		name string
		src []int
		allowed map[int]bool
		want []int
	}{
		{
			name: "часть разрешена",
			src: []int{1, 2, 3, 4},
			allowed: map[int]bool{2: true, 4: true},
			want: []int{2, 4},
		},
		{
			name: "ничего не разрешено",
			src: []int{1, 2},
			allowed: map[int]bool{},
			want: []int{},
		},
		{
			name: "nil map",
			src: []int{1, 2},
			allowed: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeepAllowedInts(tt.src, tt.allowed)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeepAllowedInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
