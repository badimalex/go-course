package homework09

import (
	"reflect"
	"testing"
)

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name   string
		people []interface{}
		want   any
	}{
		{
			name: "Empty array",
			want: nil,
		},
		{
			name: "Single user",
			people: []any{
				Customer{30},
			},
			want: Customer{30},
		},
		{
			name: "Multiple users",
			people: []any{
				Customer{25},
				Employee{40},
				Employee{32},
			},
			want: Employee{40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := MaxAge(tt.people...); !reflect.DeepEqual(gotValue, tt.want) {
				t.Errorf("MaxAge() = %v, want %v", gotValue, tt.want)
			}
		})
	}
}
