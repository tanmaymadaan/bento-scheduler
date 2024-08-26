package database

import (
	"fmt"
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

	var session *gocql.Session
	var err error

	for i := 0; i < 5; i++ {
		session, err = cluster.CreateSession()
		if err == nil {
			break
		}
		log.Printf("Failed to connect to Cassandra, retrying in 5 seconds... (attempt %d/5)", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Printf("Error creating Cassandra session after 5 attempts: %v", err)
		return nil, err
	}

	return &CassandraSession{Session: session}, nil
}

func (c *CassandraSession) TestConnection() error {
	// Simple query to test the connection
	if err := c.Session.Query("SELECT now() FROM system.local").Exec(); err != nil {
		return fmt.Errorf("error executing test query: %v", err)
	}
	return nil
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
