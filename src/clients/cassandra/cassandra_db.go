package cassandra

import (
	"github.com/gocql/gocql"
)

var session *gocql.Session

func init() {
	//Pass a list of initial node IP addresses to NewCluster to create a new cluster configuration:
	cluster := gocql.NewCluster("127.0.0.1")

	//Then you can customize more options (see ClusterConfig):
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4

	//When ready, create a session from the configuration. Don't forget to Close the session once you are done with it:
	/*
		session, err := cluster.CreateSession()
		if err != nil {
			panic(err)
		}
		fmt.Println("Cassandra OK!")
		defer session.Close()
	*/
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
