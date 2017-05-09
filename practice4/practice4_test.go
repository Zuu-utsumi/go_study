package main

import (
	"testing"
)

func TestCalcReynoldsNum(t *testing.T) {
	expect := 439.02
	d, v, rho, mu := 200.0, 15.0, 3.0, 20.5
	actual, _ := CalcReynoldsNum(d, v, rho, mu)
	if actual != expect {
		t.Errorf("Expected %f but got %f", expect, actual)
	}

	// when mu is 0
	// avoid zero division
	expect2 := 0.0
	d, v, rho, mu = 200.0, 15.0, 3.0, 0.0
	actual2, e := CalcReynoldsNum(d, v, rho, mu)
	if e == nil && actual2 != expect2 {
		t.Errorf("Expected %f but got %f", expect2, actual2)
	}
}
