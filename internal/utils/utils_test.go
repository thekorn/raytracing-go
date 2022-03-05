package utils

import "testing"

func TestGetDefaultRandomNumber(t *testing.T) {
	n := GetDefaultRandomNumber()
	if n < 0 || n > 1 {
		t.Errorf("default random number should be in [0,1]")
	}
}

func TestClamp(t *testing.T) {
	n := Clamp(0, 1, 2)
	if n != 1 {
		t.Errorf("clamped to the wrong number")
	}

	n = Clamp(1, 1, 2)
	if n != 1 {
		t.Errorf("clamped to the wrong number")
	}

	n = Clamp(1.5, 1, 2)
	if n != 1.5 {
		t.Errorf("clamped to the wrong number")
	}

	n = Clamp(2, 1, 2)
	if n != 2 {
		t.Errorf("clamped to the wrong number")
	}

	n = Clamp(3, 1, 2)
	if n != 2 {
		t.Errorf("clamped to the wrong number")
	}
}
