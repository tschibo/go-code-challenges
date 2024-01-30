package happynumber

import (
	"testing"
)

func TestIsHappy_19(t *testing.T) {
	// Helper
	probe := 19
	expected := true

	// Assert
	got := IsHappy(probe)
	if got != expected {
		t.Errorf("Expected %v for %d, got %v", expected, probe, got)
	}
}

func TestIsHappy_4(t *testing.T) {
	// Helper
	probe := 4
	expected := false

	// Assert
	got := IsHappy(probe)
	if got != expected {
		t.Errorf("Expected %v for %d, got %v", expected, probe, got)
	}
}
