package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Players = make(map[string]*Player)
)

type Player struct {
	Name                    string //PrimaryID // not nullable
	Username                string
	Years                   int
	Health                  int
	PrimaryAbilityRessource int
}

func (p Player) display() string {
	return fmt.Sprintf("Name: %v\n Username: %v\n Years: %v\n Health: %v\n PrimaryAbilityRessource: %v\n ", p.Name, p.Username, p.Years, p.Health, p.PrimaryAbilityRessource)
}

func (p Player) del() {
	delete(Players, p.Name)
	NameofFile := p.Name
	NameofFileWithExt := fmt.Sprintf("%s.yml", NameofFile)

	err := os.Remove(NameofFileWithExt)
	if err != nil {
		log.Fatalf("Error when delete file: %v", err)
	}

	fmt.Printf("file succesfuly destroyed: %s\n", NameofFileWithExt)
}

func (p Player) save() {
	fileName := fmt.Sprintf("%s.yml", p.Name)
	dataYaml, err := yaml.Marshal(&p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(dataYaml))

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error when creating file : %v", err)
	}
	os.WriteFile(fileName, dataYaml, 0600)
	defer file.Close()
}
