package mapslice

import (
	"reflect"
	"testing"
)

func TestRemoveBannedInts(t *testing.T) {
	tests := []struct {
		name string
		src []int
		banned map[int]bool
		want []int
	}{
		{
			name: "есть запрещённые",
			src: []int{1, 2, 3, 4},
			banned: map[int]bool{2: true, 4: true},
			want: []int{1, 3},
		},
		{
			name: "false не запрещает",
			src: []int{1, 2},
			banned: map[int]bool{2: false},
			want: []int{1, 2},
		},
		{
			name: "nil map",
			src: []int{1, 2},
			banned: nil,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveBannedInts(tt.src, tt.banned)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveBannedInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
