package database

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

// CassandraSession represents a connection to a Cassandra cluster
type CassandraSession struct {
	Session *gocql.Session
}

// NewCassandraSession creates and returns a new Cassandra session
func NewCassandraSession(hosts []string, keyspace string) (*CassandraSession, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = time.Second * 5

	session, err := cluster.CreateSession()
	if err != nil {
		log.Printf("Error creating Cassandra session: %v", err)
		return nil, err
	}

	return &CassandraSession{Session: session}, nil
}

// Close closes the Cassandra session
func (c *CassandraSession) Close() {
	if c.Session != nil {
		c.Session.Close()
	}
}

// ExecuteQuery executes a CQL query and returns an error if any
func (c *CassandraSession) ExecuteQuery(query string, values ...interface{}) error {
	return c.Session.Query(query, values...).Exec()
}

// ExecuteQueryWithResult executes a CQL query and scans the result into the provided struct
func (c *CassandraSession) ExecuteQueryWithResult(query string, dest interface{}, values ...interface{}) error {
	return c.Session.Query(query, values...).Scan(dest)
}
