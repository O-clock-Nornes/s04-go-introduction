package main

import (
	"fmt"
	"io"
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
	State                   string
	Years                   int
	Health                  int
	PrimaryAbilityRessource int
}

func loadFromFile(name string) (*Player, error) {
	NameofFile := name
	NameofFileWithExt := fmt.Sprintf("%s.yml", NameofFile)

	file, err := os.Open(NameofFileWithExt)
	if os.IsNotExist(err) {
		fmt.Printf("File %s not found, creating and saving player.\n", NameofFileWithExt)
	} else if err != nil {
		log.Fatalf("Error when opening file: %v", err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error when reading in file : %v", err)
		return &Player{}, err
	}
	var p Player
	err = yaml.Unmarshal([]byte(content), &p)
	if err != nil {
		log.Fatalf("error when unmarshal: %v", err)
	}

	Players[p.Name] = &p
	return &p, nil
}
func playerLoad(name string) *Player {
	p, exists := Players[name]
	if exists {
		return p
	} else {
		var p *Player
		p, err := loadFromFile(name)
		if err != nil {
			// Ici on crée un nouvel utilisateur vierge
			p = &Player{
				Name:   name,
				Health: 10,
				State:  "needPseudo",
			}
			p.save()
		}
		Players[name] = p
		return p
	}
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
