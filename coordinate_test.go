package main

import (
	"testing"
)

func TestLonLatToGoogleMercator(t *testing.T) {
	mercatorPoint := LonLatToGoogleMercator(139.624532, 31.770872)
	if mercatorPoint.X != 3.5580440147278294e+07 {
		t.Errorf("X not equal 3.5580440147278294e+07, result is mercatorPoint.X")
	}
	if mercatorPoint.Y != 1.6304236889698122e+07 {
		t.Errorf("Y not equal 1.6304236889698122e+07, result is mercatorPoint.X")
	}
}
