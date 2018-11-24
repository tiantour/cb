package cb

import (
	"log"
	"runtime"

	"github.com/gocql/gocql"
)

var (
	// Hosts hosts
	Hosts []string
	// Keyspace keyspace
	Keyspace string

	session *gocql.Session
	err     error
)

func init() {

	Hosts = []string{"127.0.0.1"}
	Keyspace = "ks"

	cluster := gocql.NewCluster(Hosts...)
	cluster.Keyspace = Keyspace
	cluster.Consistency = gocql.Consistency(1)
	cluster.NumConns = runtime.NumCPU()
	session, err = cluster.CreateSession()
	if err != nil {
		log.Panic("start", err)
		return
	}
}
