package main

import "testing"

func TestAddition(t *testing.T) {
	somme := add(3, 4)
	if somme != 7 {
		t.Fatalf("3 +4 devrait égaler 7 et non pas %v", somme)
	}
}
