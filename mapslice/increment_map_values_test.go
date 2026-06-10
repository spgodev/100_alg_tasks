package mapslice

import (
	"reflect"
	"testing"
)

func TestIncrementMapValues(t *testing.T) {
	tests := []struct {
		name string
		m map[string]int
		delta int
		want map[string]int
	}{
		{
			name: "обычный случай",
			m: map[string]int{"a": 1, "b": 2},
			delta: 3,
			want: map[string]int{"a": 4, "b": 5},
		},
		{
			name: "отрицательная delta",
			m: map[string]int{"x": 5},
			delta: -2,
			want: map[string]int{"x": 3},
		},
		{
			name: "nil map",
			m: nil,
			delta: 1,
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IncrementMapValues(tt.m, tt.delta)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncrementMapValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
