package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursname"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcom to JSON in go!")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 799, "LearningCodeOnline.in", "abc123", []string{"web-dev", "js", "react"}},
		{"Vue Bootcamp", 1299, "LearningCodeOnline.in", "bcd", []string{"web-dev", "js", "vue"}},
		{"Angular Bootcamp", 899, "LearningCodeOnline.in", "qwerty123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursname": "Vue Bootcamp",
			"Price": 1299,
			"website": "LearningCodeOnline.in",
			"tags": ["web-dev","js","vue"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID!")
	}

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("%v: %v [%T]\n", k, v, v)
	}
}
