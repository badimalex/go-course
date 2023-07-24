package hw

import "testing"

func TestGeom_CalculateDistance(t *testing.T) {
	tests := []struct {
		name           string
		x1, y1, x2, y2 float64
		wantDistance   float64
	}{
		{
			name: "#1",
			x1:   1, y1: 1, x2: 4, y2: 5,
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := distance(tt.x1, tt.y1, tt.x2, tt.y2); gotDistance != tt.wantDistance {
				t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
