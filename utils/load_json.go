package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PersonaElement struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetNamesFromJson(path string) []string {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	//配列型のJSONデータを読み込む
	personaElement := make([]*PersonaElement, 0)
	err = json.Unmarshal(file, &personaElement)
	if err != nil {
		fmt.Println("error:", err)
	}

	//jsonを変形させる
	names := []string{}
	for _, element := range personaElement {

		names = append(names, element.Name)
	}
	return names
}
