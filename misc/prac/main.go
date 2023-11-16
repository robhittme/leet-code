package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
    Name string `json:"name"`
    Email string `json:"email"`
}
func main() {
    person := Person{
        Name: "rob",
        Email: "robhittme@gmail.com",
    }

    jsonBytes, err := json.Marshal(person)
    if err != nil {
        log.Fatal(err)
        return
    }

    fmt.Println(string(jsonBytes))
    fmt.Println("finished")
}
