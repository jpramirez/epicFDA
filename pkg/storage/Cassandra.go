package cassandra

import (
  "github.com/gocql/gocql"
  "fmt"
  "time"
  "io/ioutil"
  "path/filepath"
  "os"
  models "github.com/jpramirez/epicFDA/pkg/models"

)


type StorageCassandra struct {
	Config models.Config
	Session *gocql.Session
}

// Session holds the connection to Cassandra


func (S *StorageCassandra) Init() {

	var err error
	cluster := gocql.NewCluster(S.Config.CassandraHost...)
	cluster.Keyspace = S.Config.CassandraKeySpace
	S.Session, err = cluster.CreateSession()
	if err != nil {
	  panic(err)
	}
	fmt.Println("cassandra init done")
  }



  //MigrateDatabase will create the kespace and tables
  func (S *StorageCassandra) MigrateDatabase (){
		// connect to the cluster
	fmt.Println("MigrateDatabase Started")
	cluster := gocql.NewCluster(S.Config.CassandraHost[0]) //replace PublicIP with the IP addresses used by your cluster.
	
	fmt.Println("Step 1" + S.Config.CassandraHost[0] )

	//cluster.Consistency = gocql.Quorum
	//cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10

	fmt.Println("Step 2")
	// No Authenticator for now.
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.


	session, err := cluster.CreateSession()

	fmt.Println("Step 3")
	if err != nil {
			fmt.Println(err)
			return
	}
	defer session.Close()
	

	// Cycle Through the ables in SQL folder 
	var files []string
	
    root := "extras/database/"
	
	fmt.Println(root)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if filepath.Ext(path) != ".sql" {
			return nil
		}
        files = append(files, path)
        return nil
    })
    if err != nil {
        fmt.Println(err)
    }
    for _, sqlfile := range files {
		fmt.Println(sqlfile)
		file, err := os.Open(sqlfile)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	
	
	 	 b, err := ioutil.ReadAll(file)
	 	 err = session.Query(string(b)).Exec()
	 	 if err != nil {
		 	 fmt.Println(err)
		 	 return
	  	}
    }


 
  }