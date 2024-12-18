package main

import (
	"math"
	"testing"
)

func TestCalculateRoots(t *testing.T) {
	tests := []struct {
		a, b, c        float64
		wantX1, wantX2 float64
		expectError    bool
	}{
		{1, -3, 2, 2, 1, false},  // Два действительных корня
		{1, 2, 1, -1, -1, false}, // Один действительный корень (D = 0)
		{1, 0, 1, 0, 0, true},    // Нет действительных корней (D < 0)
		{1, 2, 3, 0, 0, false},   // Провленный тест
	}

	for _, tt := range tests {
		x1, x2, err := CalculateRoots(tt.a, tt.b, tt.c)
		if (err != nil) != tt.expectError {
			t.Errorf("CalculateRoots(%v, %v, %v) ожидал ошибку %v, но получил %v", tt.a, tt.b, tt.c, tt.expectError, err != nil)
			continue
		}
		if !tt.expectError {
			if math.Abs(x1-tt.wantX1) > 0.001 || math.Abs(x2-tt.wantX2) > 0.001 {
				t.Errorf("CalculateRoots(%v, %v, %v) = (%v, %v), ожидалось (%v, %v)", tt.a, tt.b, tt.c, x1, x2, tt.wantX1, tt.wantX2)
			}
		}
	}
}
