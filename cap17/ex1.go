package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"name"`
	Age   int    `json:"age"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	j, err := json.Marshal(users)

	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println(string(j))
	}
}
