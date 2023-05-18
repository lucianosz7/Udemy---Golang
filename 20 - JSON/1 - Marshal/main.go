package main

import(
	"encoding/json"
	"fmt"
	"log"
	"bytes"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint	`json:"idade"`
}

func main()  {
	c := cachorro{"Thor", "Rotweiller", 2}
	fmt.Println(c)

	cachorroJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson)

	fmt.Println(bytes.NewBuffer(cachorroJson))

	c2 := map[string]string{
		"name":"Toby",
		"raca": "Poodle",
	}

	cachorro2Json, erro :=  json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorro2Json))
}