package main

import (
	"log"
	"strings"
	"testing"
)

func TestSave(t *testing.T) {
	p4 := Player{
		Name:     "Patrick",
		Username: "THE_STAR",
	}
	p4.save()
	p4.del()
}

func TestDisplay(t *testing.T) {
	p := Player{
		Name:     "Patrick",
		Username: "THE_STAR",
	}
	if !strings.Contains(p.display(), "THE_STAR") {
		t.Fatalf("Username THE_STAR should appear in %v", p.display())
	}
	//log.Printf(p.display())
}
func TestLoad(t *testing.T) {
	p := playerLoad("toto")
	log.Printf("%v", p)
}
