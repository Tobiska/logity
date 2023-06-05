package room

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"logity/config"
	"logity/internal/domain/entity/room"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/room/dto"
)

type Repository struct {
	driver neo4j.DriverWithContext
	cfg    *config.Neo4j
}

func NewRepository(driver neo4j.DriverWithContext, cfg *config.Neo4j) *Repository {
	return &Repository{
		driver: driver,
		cfg: cfg,
	}
}

func (r *Repository) CreateRoom(ctx context.Context, userId string, inputRoom *room.Room) (*room.Room, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,

	})

	resp, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `MATCH (u:User) WHERE u.id = "1" CREATE (u)-[:OWNER]-> (r:Room {id: apoc.create.uuid(), username: "", tag: ""}) RETURN apoc.convert.toJson(r)`,
			map[string]any{
				"userId": userId,
				"username": inputRoom.Name,
				"tag": inputRoom.Tag,
			})
		if err != nil {
			return nil, fmt.Errorf("execute create room query: %w", err)
		}

		res.

	})
	if err != nil {
		return nil, fmt.Errorf("error execute create room: %w", err)
	}

	createdRoom := &room.Room{}
	json.Unmarshal()

	return nil, nil
}

func (r *Repository) GetRoomByCode(ctx context.Context, roomCode string) (*room.Room, error) {
	panic("implement GetRoomByCode")
}
func (r *Repository) UpdateRoom(ctx context.Context, dto dto.UpdateRoomDto) (*room.Room, error) {
	panic("imlement UpdateRoom")
}
func (r *Repository) DeleteRoom(ctx context.Context, roomCode string) (*room.Room, error) {
	panic("implement DeleteRoom")
}

func (r *Repository) FindRoomByFilter(ctx context.Context, filter dto.FindFilter) ([]*room.Room, error) {
	panic("implement FindRoom")
}
func (r *Repository) ShowAllCreatedRoom(ctx context.Context, userId string) ([]*room.Room, error) {
	panic("sdsd")
}
func (r *Repository) ShowAllAttachedRoom(ctx context.Context, userId string) ([]*room.Room, error) {
	panic("dsdsds")
}

func (r *Repository) AttachUserToRoom(ctx context.Context, userId ,roomCode string) (*room.Room, error) {
	panic("dsdsdsd")
}
func (r *Repository) DetachUserFromRoom(ctx context.Context, userId, roomCode string) (*room.Room, error) {
	panic("dsdsdsdsdsds")
}
