package unit1_test

import (
	"study-golang/unit1"
	"testing"
)

func TestUnitOne(t *testing.T) {
	if unit1.UnitOne() != "Unit One" {
		t.Fatal("Wrong number : (")
	}
}

func TestUnitOne2(t *testing.T) {
	if unit1.UnitOne() != "Unit ONe" {
		t.Fatal("Wrong number : (")
	}
}
