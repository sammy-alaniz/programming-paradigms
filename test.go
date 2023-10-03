package main

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "test2023test", ""))
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close()

	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session := driver.NewSession(sessionConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("CREATE (n:Person {name: $name}) RETURN n.name", map[string]interface{}{"name": "Sammy Doe"})
		if err != nil {
			return nil, err
		}
		record, err := result.Single()
		if err != nil {
			return nil, err
		}
		return record.Values[0], nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
