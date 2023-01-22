package main

import (
	"fmt"
	"persona-backend-restfulapi/utils"
	"testing"
)

func TestGetRecord(t *testing.T) {

	t.Run("returns array from json", func(t *testing.T) {
		//ファイルを読みだす
		array := utils.GetNamesFromJson("personality-json/hobbies.json")

		got := array[0]
		want := "釣り"

		assertField(t, got, want, array)
	})

}

func TestGetPersona(t *testing.T) {
	t.Run("returns persona from persona-element", func(t *testing.T) {
		personality_combination := utils.GetNamesFromJson("personality-combination/choosing-way.json")

		paths := []string{}

		for _, path := range personality_combination {
			paths = append(paths, "personality-json/"+path+".json")
		}

		//get each persona-element
		hobbies := utils.GetNamesFromJson(paths[1])
		jobs := utils.GetNamesFromJson(paths[0])

		//combine personality
		persona := []string{}
		persona = append(persona, hobbies[0])
		persona = append(persona, jobs[1])

		fmt.Println(persona)
	})
}
func assertField(t testing.TB, got any, want string, record any) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s given, %v", got, want, record)
	}
}
