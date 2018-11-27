package cb

import (
	"log"
	"runtime"

	"github.com/gocql/gocql"
	"github.com/tiantour/conf"
)

var (
	session *gocql.Session
	err     error
)

func init() {
	c := conf.NewCB().Data
	if c.Keyspace == "" {
		log.Fatal("cb conf is null")
	}
	if len(c.IP) == 0 {
		c.IP = []string{"127.0.0.1"}
	}

	cluster := gocql.NewCluster(c.IP...)
	cluster.Keyspace = c.Keyspace
	cluster.Consistency = gocql.Consistency(0)
	cluster.NumConns = runtime.NumCPU()
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
}
