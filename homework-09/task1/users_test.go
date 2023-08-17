package homework09

import "testing"

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name  string
		users []User
		want  int
	}{
		{
			name: "Empty array",
			want: 0,
		},
		{
			name: "Single user",
			users: []User{
				Customer{30},
			},
			want: 30,
		},
		{
			name: "Multiple users",
			users: []User{
				Customer{25},
				Employee{40},
				Employee{32},
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.users...); got != tt.want {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
