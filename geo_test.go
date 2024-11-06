package main

import "testing"

func TestCoordNajac(t *testing.T) {
	coord, err := getCoordFromCity("najac")
	if err != nil {
		t.Fatalf("Impossible d'extraire les coordonnées de Najac : %v", err)
	}
	coordAttendue := Coordinates{Lat: 44.22018, Lon: 1.9809309}
	if coord != coordAttendue {
		t.Fatalf("Les coordonnées attendues pour Najac (%v) ne sont pas les bonnes : %v", coordAttendue, coord)
	}
}
