package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ConnectionData struct {
	uri string
	username string
	password string
	driver neo4j.Driver
	err error
}

func (cd *ConnectionData) init() {
	cd.driver, cd.err = neo4j.NewDriver(cd.uri, neo4j.BasicAuth(cd.username, cd.password, ""))

	if cd.err != nil {
        panic(cd.err)
    }

    fmt.Println("Driver established!")
}

func (cd *ConnectionData) shutdown() {
	if cd.driver == nil {
		fmt.Println("Driver is nil")
	} else {
		cd.driver.Close()
	}
}
