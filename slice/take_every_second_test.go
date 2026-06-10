package slice

import (
	"reflect"
	"testing"
)

func TestTakeEverySecond(t *testing.T) {
	tests := []struct {
		name string
		src []int
		want []int
	}{
		{
			name: "нечётная длина",
			src: []int{10, 20, 30, 40, 50},
			want: []int{10, 30, 50},
		},
		{
			name: "чётная длина",
			src: []int{1, 2, 3, 4},
			want: []int{1, 3},
		},
		{
			name: "пустой слайс",
			src: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TakeEverySecond(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TakeEverySecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
