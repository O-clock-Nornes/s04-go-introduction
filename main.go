package main

import (
	"io"
	"log"
	"os"
)

func main() {

	// D'abord on ouvre le fichier test.txt
	f, err := os.Open("test.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	// Ici on le lit
	data, _ := io.ReadAll(f)
	log.Printf("%v", string(data))

	// On fait un nouveau fichier avec le résultat de notre lecture
	os.WriteFile("test-copie.txt", data, 0600)

}
