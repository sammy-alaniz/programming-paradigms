package main

import (
	"fmt"
	// "net/http"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, World!")
	// })

	// http.ListenAndServe(":8080", nil)

	var driver neo4j.Driver

	var address = "neo4j://localhost:7687"
	var username = "neo4j"
	var password = "test2023test"

	driver, err := neo4j.NewDriver(address, neo4j.BasicAuth(username, password, ""))

    if err != nil {
        panic(err)
    }

    fmt.Println("Driver established!")

	defer driver.Close()
}
