package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages without ruby", languages)

	// loops over the maps in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// loops over the maps in golang
	for _, value := range languages {
		fmt.Printf("Value is %v\n", value)
	}
}
