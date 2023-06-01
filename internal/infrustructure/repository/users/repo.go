package users

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"logity/config"
	"logity/internal/domain/entity/user"
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

func (r *Repository) CreateUser(ctx context.Context, u *user.User) error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: r.cfg.Database,
	})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `CREATE (u:User {
			id: $userId,
			email: $email,
			phone: $phone,
			username: $username
		})`, map[string]any{
			"userId":   u.Id,
			"email":    u.Email,
			"phone":    u.Phone,
			"username": u.Username,
		})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		return err
	}

	return nil
}
