package log

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"logity/config"
	"logity/internal/domain/entity/log"
	"time"
)

var (
	ErrCreateLog = fmt.Errorf("error create and relate log, check input user id and room ids")
)

type Repository struct {
	driver neo4j.DriverWithContext
	cfg    *config.Neo4j
}

func NewRepository(driver neo4j.DriverWithContext, cfg *config.Neo4j) *Repository {
	return &Repository{
		driver: driver,
		cfg:    cfg,
	}
}

func (r *Repository) GetLogs(_ context.Context) ([]*log.Log, error) {
	panic("implement GetLogs!!!")
}

func (r *Repository) CreateLogText(ctx context.Context, l *log.LogText, roomIds []string) error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: r.cfg.Database,
	})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx,
			`CREATE (l:Log:Text {id: randomUUID(), type: 'text', text: $text, created_at: $createdAt})
					WITH l
					
					MATCH (u:User {id: $userId})
					WITH u, l
					
					CREATE (l)<-[:OWNED]-(u)
					WITH u, l
					
					MATCH (r:Room) WHERE r.id in $roomIds
					CREATE (r)<-[:PUSHED]-(l)`,
			map[string]any{
				"userId":    l.Owner.Id,
				"roomIds":   roomIds,
				"text":      l.Text,
				"createdAt": l.CreatedAt.Format(time.RFC3339),
			})
		if err != nil {
			return nil, err
		}

		resultQuery, err := res.Consume(ctx)
		if err != nil {
			return nil, err
		}
		if resultQuery.Counters().RelationshipsCreated() != 1+len(roomIds) || resultQuery.Counters().NodesCreated() != 1 {
			return nil, ErrCreateLog
		}

		return nil, nil
	})

	if err != nil {
		return fmt.Errorf("error create log: %w", err)
	}
	return nil
}

func (r *Repository) CreateLogPhoto(ctx context.Context, l *log.LogPhoto, roomIds []string) error {
	panic("implement CreateLogPhoto me!!!")
}

func (r *Repository) CreateLogPicture(ctx context.Context, l *log.LogPicture, roomIds []string) error {
	panic("implement CreateLogPicture me!!!")
}
