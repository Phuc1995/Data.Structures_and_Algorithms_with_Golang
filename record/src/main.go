package main

import (
	"encoding/json"
	"fmt"
)

type (
	Gopher struct {
		Name      string `json: "name, omitempty"`
		Age       int `json:"age"`
		Gender bool
		NickNames []string
		Secrets   map[string]interface{}
	}

	address struct {
		Country string
		City    string
	}
)

func main() {
	zeroGopher := Gopher{
		Name: "gopher",
		Age: 13,
		Gender:false,
		NickNames: []string{"jack", "sparrow"},
		Secrets: map[string]interface{}{
			"real name" : "jack",
			"real age" : 4444,
		},
	}

	b, err := json.Marshal(zeroGopher)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(b))
}

func ChangeName(gopher Gopher) {
	gopher.Name = "Test"
}
func (g *Gopher) ChangeName(newName string) {
	g.Name = newName
}

func (g Gopher) GetName() string {
	return "Name-" + g.Name
}
