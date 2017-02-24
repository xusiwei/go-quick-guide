package main

import "fmt"
import "log"
import "encoding/json"

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Color:  false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title:  "Cool Hand Luke",
			Year:   1967,
			Color:  true,
			Actors: []string{"Paul Newman"},
		},
		{
			Title: "Bullitt",
			Year:  1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
		},
		// ...
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var dump []Movie
	err = json.Unmarshal(data, &dump)
	if err != nil {
		log.Fatalf("Unmarshal failed: %s", err)
	}
	fmt.Println("\nUnmarshaled:\n", dump)

	data, err = json.MarshalIndent(dump, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("\nIndented:\n%s\n", data)
}
