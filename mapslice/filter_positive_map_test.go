package mapslice

import (
	"reflect"
	"testing"
)

func TestFilterPositiveMap(t *testing.T) {
	tests := []struct {
		name string
		m map[string]int
		want map[string]int
	}{
		{
			name: "смешанные значения",
			m: map[string]int{"a": 1, "b": 0, "c": -2},
			want: map[string]int{"a": 1},
		},
		{
			name: "положительных нет",
			m: map[string]int{"x": 0, "y": -1},
			want: map[string]int{},
		},
		{
			name: "nil map",
			m: nil,
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterPositiveMap(tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterPositiveMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
