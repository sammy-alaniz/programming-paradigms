package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ConnectionData struct {
	uri      string
	username string
	password string
	driver   neo4j.Driver
	err      error
	uuid     string
}

func (cd *ConnectionData) init() {
	fmt.Println("initalizing neo4j connection")

	cd.uuid = uuid.New().String()
	fmt.Println(cd.uuid)

	cd.driver, cd.err = neo4j.NewDriver(cd.uri, neo4j.BasicAuth(cd.username, cd.password, ""))

	if cd.err != nil {
		panic(cd.err)
	}

	fmt.Println("Driver established!")
}

// in work func
// func (cd *ConnectionData) query(query string) (hash string, results map[string]interface{}, err error) {

// }

func (cd *ConnectionData) shutdown() {
	if cd.driver == nil {
		fmt.Println("Driver is nil")
	} else {
		cd.driver.Close()
	}
}
