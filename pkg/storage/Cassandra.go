package cassandra

import (
  "github.com/gocql/gocql"
  "fmt"
)



// Session holds the connection to Cassandra
var Session *gocql.Session


func init() {
	var err error
	cluster := gocql.NewCluster("192.168.64.130")
	cluster.Keyspace = "epicfda"
	Session, err = cluster.CreateSession()
	if err != nil {
	  panic(err)
	}
	fmt.Println("cassandra init done")
  }