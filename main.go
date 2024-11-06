package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Commune struct {
	Nom  string `json:"nom"`
	Code string `json:"code"`
	Pop  int    `json:"population"`
}

func main() {
	resp, err := http.Get("https://geo.api.gouv.fr/communes?codePostal=12000")
	if err != nil {
		log.Fatalf("Lors de la récuparation de geo api : %v", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Lors de la récuparation de geo api, le code de retour est pas bon : %v", resp.StatusCode)
	}
	//log.Printf("%v", resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Lecture du body %v", err)
	}
	//log.Printf("%v", string(body))

	var communes []Commune

	err = json.Unmarshal(body, &communes)
	if err != nil {
		log.Fatalf("Probleme pour extraire le fichier json")
	}

	for _, commune := range communes {
		log.Printf("Commune %v avec code postal %v et population de %v personnes", commune.Nom, commune.Code, commune.Pop)
	}
	//log.Printf("%#v", communes)

	//log.Printf("%#v", resp)

}
