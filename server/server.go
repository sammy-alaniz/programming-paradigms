package main

import (
	"fmt"
)

func main() {

	fmt.Println("Start of sever!")

	var cd ConnectionData

	cd.uri = "neo4j://localhost:7687"
	cd.username = "neo4j"
	cd.password = "test2023test"

	cd.init()

	var hr HttpRest

	hr.cd = &cd
	hr.start_http_request_handlers()
	cd.shutdown()
}
