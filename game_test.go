package main_test

import (
	"testing"

	"github.com/selfup/goOne"
)

func TestItStartsWithXAtZero(t *testing.T) {
	bike := main.BikeOne{}
	bikeX := bike.X

	if bikeX != 0 {
		t.Fatalf("expected 0, got %s instead", bike.X)
	}
}

func TestItStartsWithYAtZero(t *testing.T) {
	bike := main.BikeOne{}
	bikeY := bike.Y

	if bikeY != 0 {
		t.Fatalf("expected 0, got %s instead", bike.Y)
	}
}
