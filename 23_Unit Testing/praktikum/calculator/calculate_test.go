package calculator

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}
	if Addition(-1, -2) != -3 {
		t.Error("Expected (-1) + (-2) to equal -3")
	}
}

func TestSubstract(t *testing.T) {
	if Substract(2, 1) != 1 {
		t.Error("Expected 2 - 1 to equal 1")
	}
	if Substract(-1, -2) != 1 {
		t.Error("Expected (-1) - (-2) to equal 1")
	}
}

func TestMult(t *testing.T) {
	if Mult(2, 1) != 2 {
		t.Error("Expected 2 * 1 to equal 2")
	}
	if Mult(-2, -1) != 2 {
		t.Error("Expected -2 * -1 to equal 2")
	}
}

func TestDiv(t *testing.T) {
	if Div(2, 1) != 2 {
		t.Error("Expected 2 / 1 to equal 2")
	}
	if Div(-2, -1) != 2 {
		t.Error("Expected -2 / -1 to equal 2")
	}
}
