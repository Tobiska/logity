package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"logity/config"
)

func NewDriverNeo4j(cfg *config.Neo4j) (neo4j.DriverWithContext, error) {
	driver, err := neo4j.NewDriverWithContext(cfg.Host, neo4j.BasicAuth(cfg.Username, cfg.Password, ""))
	if err != nil {
		return nil, err
	}
	return driver, nil
}
