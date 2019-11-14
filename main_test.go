package main

import "testing"

func TestSum(t *testing.T) {
	if add(2, 5) != 7 {
		t.Fail()
	}

	if add(20, 5) != 25 {
		t.Fail()
	}

	if add(26, 5) != 31 {
		t.Fail()
	}

}
func TestProduct(t *testing.T) {
	if multiply(2, 5) != 10 {
		t.Fail()
	}
	if multiply(2, 100) != 200 {
		t.Fail()
	}
	if multiply(222, 3) != 666 {
		t.Fail()
	}
}
