package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	fmt.Println("Start of sever!")

	// var cd ConnectionData

	// cd.uri = "neo4j://localhost:7687"
	// cd.username = "neo4j"
	// cd.password = "test2023test"

	// cd.init()
	// cd.shutdown()

	u := uuid.New()
	fmt.Println(u)

	start_http_request_handlers()
}
