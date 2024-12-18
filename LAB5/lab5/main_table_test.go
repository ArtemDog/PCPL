package main

import (
	"math"
	"testing"
)

func TestCalculateRootsTable(t *testing.T) {
	tests := []struct {
		name           string
		a, b, c        float64
		wantX1, wantX2 float64
		expectError    bool
	}{
		{"TwoRealRoots", 1, -3, 2, 2, 1, false},
		{"OneRealRoot", 1, 2, 1, -1, -1, false},
		{"NoRealRoots", 1, 0, 1, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x1, x2, err := CalculateRoots(tt.a, tt.b, tt.c)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error %v, got %v", tt.expectError, err)
				return
			}

			if !tt.expectError {
				if math.Abs(x1-tt.wantX1) > 0.001 || math.Abs(x2-tt.wantX2) > 0.001 {
					t.Errorf("CalculateRoots(%v, %v, %v) = (%v, %v), want (%v, %v)", tt.a, tt.b, tt.c, x1, x2, tt.wantX1, tt.wantX2)
				}
			}
		})
	}
}
