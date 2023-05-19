package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"hello bond",
			"heyyy james",
		},
	}

	u2 := user{
		First: "Moneypenny",
		Last:  "Miss",
		Age:   27,
		Sayings: []string{
			"hello money",
			"heyyy penny",
		},
	}

	u3 := user{
		First: "Nimo",
		Last:  "Yash",
		Age:   21,
		Sayings: []string{
			"hello nimo",
			"heyyy yash",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
