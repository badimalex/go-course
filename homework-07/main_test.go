package main

import (
	"sort"
	"testing"
)

func Test_sortInts(t *testing.T) {
	s := []int{9, 8, 5}
	SortInts(s)
	expected := []int{5, 8, 9}

	for i := range expected {
		if s[i] != expected[i] {
			t.Errorf("expected '%d' but got '%d'", expected[i], s[i])
		}
	}
}

func Test_sortString(t *testing.T) {
	var tests = []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "Test №1",
			args: []string{"Go", "Bravo", "Alpha"},
			want: []string{"Alpha", "Bravo", "Go"},
		},
		{
			name: "Test №2",
			args: []string{"Alpha", "Go", "Bravo"},
			want: []string{"Alpha", "Bravo", "Go"},
		},
		{
			name: "Test №3",
			args: []string{},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortString(tt.args)

			for i := range tt.want {
				if tt.args[i] != tt.want[i] {
					t.Errorf("expected '%s' but got '%s'", tt.want[i], tt.args[i])
				}
			}
		})
	}
}

func BenchmarkSortInts(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := []int{9, 8, 5}
		sort.Ints(s)
	}
}

func BenchmarkSortFloat64s(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := []float64{9.1, 8, 5}
		sort.Float64s(s)
	}
}
