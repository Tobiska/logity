package room

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"logity/config"
	"logity/internal/domain/entity/room"
	"logity/internal/domain/usecase/room/dto"
)

var (
	ErrValuesMoreOne = fmt.Errorf("error values more one")
	ErrUserNotExist  = fmt.Errorf("error user doesn't exist")
)

type Repository struct {
	driver neo4j.DriverWithContext
	cfg    *config.Neo4j
}

func NewRoomRepository(driver neo4j.DriverWithContext, cfg *config.Neo4j) *Repository {
	return &Repository{
		driver: driver,
		cfg:    cfg,
	}
}

func (r *Repository) CreateRoom(ctx context.Context, userId string, inputRoom *room.Room) (*room.Room, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: r.cfg.Database,
	})

	resp, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `MATCH (u:User) WHERE u.id = "84c91898-897a-4b7c-81ab-a5b99edc24c2"
		CREATE (u)-[:OWNER]-> (r:Room {id: apoc.create.uuid(), username: "", tag: ""})
		RETURN apoc.convert.toJson(properties(r)), apoc.convert.toJson(properties(u))`,
			map[string]any{
				"userId":   userId,
				"roomName": inputRoom.Name,
				"roomTag":  inputRoom.Tag,
			})
		if err != nil {
			return nil, fmt.Errorf("execute create room query: %w", err)
		}

		if res.Next(ctx) {
			if len(res.Record().Values) > 2 {
				return nil, ErrValuesMoreOne
			}

			infRoom := Room{}

			if err := json.Unmarshal([]byte(res.Record().Values[0].(string)), &infRoom); err != nil {
				return nil, fmt.Errorf("error unmarhsal result room: %w", err)
			}

			infUserOwner := User{}

			if err := json.Unmarshal([]byte(res.Record().Values[1].(string)), &infUserOwner); err != nil {
				return nil, fmt.Errorf("error unmarhsal result creator user: %w", err)
			}

			infRoom.Owner = infUserOwner
			return infRoom, nil
		}

		return nil, ErrUserNotExist
	})

	if err != nil {
		return nil, fmt.Errorf("error execute query at session: %w", err)
	}

	return resp.(Room).toDomain(), nil
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

func (r *Repository) AttachUserToRoom(ctx context.Context, userId, roomId string) (*room.Room, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: r.cfg.Database,
	})

	resp, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `MATCH (u:User) WHERE u.id = $userId
		MATCH (r:Room) WHERE r.id = $roomId 
		MERGE (u)-[:JOINED]->(r) RETURN apoc.convert.toJson(properties(r)), apoc.convert.toJson(properties(u))`,
			map[string]any{
				"userId": userId,
				"roomId": roomId,
			})
		if err != nil {
			return nil, fmt.Errorf("execute create room query: %w", err)
		}

		if res.Next(ctx) {
			if len(res.Record().Values) > 2 {
				return nil, ErrValuesMoreOne
			}

			infRoom := Room{}

			if err := json.Unmarshal([]byte(res.Record().Values[0].(string)), &infRoom); err != nil {
				return nil, fmt.Errorf("error unmarhsal result room: %w", err)
			}

			infUserOwner := User{}

			if err := json.Unmarshal([]byte(res.Record().Values[1].(string)), &infUserOwner); err != nil {
				return nil, fmt.Errorf("error unmarhsal result creator user: %w", err)
			}

			infRoom.Owner = infUserOwner
			return infRoom, nil
		}

		return nil, ErrUserNotExist
	})

	if err != nil {
		return nil, fmt.Errorf("error execute query at session: %w", err)
	}

	return resp.(Room).toDomain(), nil
}

func (r *Repository) InviteUserToRoom(ctx context.Context, userId, roomId string) (*room.Room, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: r.cfg.Database,
	})

	resp, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `MATCH (u:User) WHERE u.id = $userId
		MATCH (r:Room) WHERE r.id = $roomId 
		MERGE (u)-[:JOINED]->(r) RETURN apoc.convert.toJson(properties(r)), apoc.convert.toJson(properties(u))`,
			map[string]any{
				"userId": userId,
				"roomId": roomId,
			})
		if err != nil {
			return nil, fmt.Errorf("execute create room query: %w", err)
		}

		if res.Next(ctx) {
			if len(res.Record().Values) > 2 {
				return nil, ErrValuesMoreOne
			}

			infRoom := Room{}

			if err := json.Unmarshal([]byte(res.Record().Values[0].(string)), &infRoom); err != nil {
				return nil, fmt.Errorf("error unmarhsal result room: %w", err)
			}

			infUserOwner := User{}

			if err := json.Unmarshal([]byte(res.Record().Values[1].(string)), &infUserOwner); err != nil {
				return nil, fmt.Errorf("error unmarhsal result creator user: %w", err)
			}

			infRoom.Owner = infUserOwner
			return infRoom, nil
		}

		return nil, ErrUserNotExist
	})

	if err != nil {
		return nil, fmt.Errorf("error execute query at session: %w", err)
	}

	return resp.(Room).toDomain(), nil
}

func (r *Repository) DetachUserFromRoom(ctx context.Context, userId, roomCode string) (*room.Room, error) {
	panic("dsdsdsdsdsds")
}
