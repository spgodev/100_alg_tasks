package mapslice

import (
	"testing"
)

func TestCanMakePairSum(t *testing.T) {
	tests := []struct {
		name string
		src []int
		target int
		want bool
	}{
		{
			name: "пара есть",
			src: []int{2, 7, 11},
			target: 9,
			want: true,
		},
		{
			name: "нельзя один индекс дважды",
			src: []int{3, 1, 4},
			target: 6,
			want: false,
		},
		{
			name: "два одинаковых значения",
			src: []int{3, 3},
			target: 6,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanMakePairSum(tt.src, tt.target)
			if got != tt.want {
				t.Errorf("CanMakePairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
