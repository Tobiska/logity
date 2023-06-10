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
	ErrValuesMoreOne                = fmt.Errorf("error values more one")
	ErrUserNotExist                 = fmt.Errorf("error user doesn't exist")
	ErrUserNotExistOrAlreadyInvited = fmt.Errorf("error user doesn't exist or user already invited")
	ErrInviteNotExist               = fmt.Errorf("error user doesn't invited to room")
	ErrRoomNotExist                 = fmt.Errorf("error room doesn't exist")
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
		res, err := tx.Run(ctx, `MATCH (u:User) WHERE u.id = $userId
		CREATE (u)-[:OWNED]-> (r:Room {id: apoc.create.uuid(), username: $roomName, tag: $roomTag})
		CREATE (u)-[:JOINED]-> (r)
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
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: r.cfg.Database,
	})

	resp, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		newRoom, err := r.getRoomInformation(ctx, tx, roomCode)
		if err != nil {
			return nil, fmt.Errorf("error get room information: %w", err)
		}

		return newRoom, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error execute query at session: %w", err)
	}

	return resp.(*room.Room), nil
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

func (r *Repository) CheckInvite(ctx context.Context, userId, roomId string) error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: r.cfg.Database,
	})

	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if err := r.checkInvite(ctx, tx, userId, roomId); err != nil {
			return nil, fmt.Errorf("error check invite: %w", err)
		}

		return nil, nil
	})

	if err != nil {
		return fmt.Errorf("error execute query at session: %w", err)
	}

	return nil
}

func (r *Repository) checkInvite(ctx context.Context, tx neo4j.ManagedTransaction, userId, roomId string) error {
	resp, err := tx.Run(ctx, `MATCH (u:User) WHERE u.id=$userId
		MATCH (r:Room) WHERE r.id=$roomId
		MATCH (u)-[i:INVITED]->(r) RETURN i`,
		map[string]any{
			"userId": userId,
			"roomId": roomId,
		})
	if err != nil {
		return fmt.Errorf("execute create room query: %w", err)
	}

	if resp.Next(ctx) {
		if len(resp.Record().Values) > 0 {
			return nil
		}
	}
	return ErrInviteNotExist
}

func (r *Repository) AttachUserToRoom(ctx context.Context, userId, roomId string) (*room.Room, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: r.cfg.Database,
	})

	resp, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if err := r.checkInvite(ctx, tx, userId, roomId); err != nil {
			return nil, fmt.Errorf("error check invite to room: %w", err)
		}

		if err := r.attachUserToRoom(ctx, tx, userId, roomId); err != nil {
			return nil, fmt.Errorf("error attach user to room: %w", err)
		}

		r, err := r.getRoomInformation(ctx, tx, roomId)
		if err != nil {
			return nil, fmt.Errorf("error get room information: %w", err)
		}

		return r, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error execute query at session: %w", err)
	}

	return resp.(*room.Room), nil
}

func (r *Repository) getRoomInformation(ctx context.Context, tx neo4j.ManagedTransaction, roomId string) (*room.Room, error) {
	resp, err := tx.Run(ctx, `MATCH (r:Room {id: '319e860d-73a0-4479-a3bc-59a39f727c08'})
OPTIONAL MATCH (r)<-[:OWNED]-(owner)
OPTIONAL  MATCH (r)<-[:JOINED]-(members)    
OPTIONAL  MATCH (r)<-[:INVITED]-(inviters)    
RETURN apoc.convert.toJson(properties(r)), apoc.convert.toJson(properties(owner)), apoc.convert.toJson(apoc.convert.toList(collect(properties(members)))), apoc.convert.toJson(apoc.convert.toList(collect(properties(inviters))))`,
		map[string]any{
			"roomId": roomId,
		})
	if err != nil {
		return nil, fmt.Errorf("error execute get room: %w", err)
	}

	if resp.Next(ctx) {
		if len(resp.Record().Values) != 4 {
			return nil, ErrRoomNotExist
		}

		r := &Room{}
		roomRaw := resp.Record().Values[0]
		if err := json.Unmarshal([]byte(roomRaw.(string)), r); err != nil {
			return nil, fmt.Errorf("error marshal room information: %w", err)
		}

		ownerRaw := resp.Record().Values[1]
		if err := json.Unmarshal([]byte(ownerRaw.(string)), &r.Owner); err != nil {
			return nil, fmt.Errorf("error marshal owner information: %w", err)
		}

		membersRaw := resp.Record().Values[2]
		if err := json.Unmarshal([]byte(membersRaw.(string)), &r.Members); err != nil {
			return nil, fmt.Errorf("error marshal members information: %w", err)
		}

		invitersRaw := resp.Record().Values[3]
		if err := json.Unmarshal([]byte(invitersRaw.(string)), &r.Inviters); err != nil {
			return nil, fmt.Errorf("error marshal inviters information: %w", err)
		}

		return r.toDomain(), nil
	}

	return nil, ErrRoomNotExist
}

func (r *Repository) attachUserToRoom(ctx context.Context, tx neo4j.ManagedTransaction, userId, roomId string) error {
	resp, err := tx.Run(ctx, `MATCH (u:User) WHERE u.id = $userId
		MATCH (r:Room) WHERE r.id = $roomId 
		MATCH (u)-[i:INVITED]->(r)
		DELETE i
		MERGE (u)-[:JOINED]->(r)`,
		map[string]any{
			"userId": userId,
			"roomId": roomId,
		})
	if err != nil {
		return fmt.Errorf("execute create room query: %w", err)
	}
	return resp.Err()
}

func (r *Repository) InviteUserToRoom(ctx context.Context, userId, roomId string) error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: r.cfg.Database,
	})

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `MATCH (u:User) WHERE u.id = $userId
		MATCH (r:Room) WHERE r.id = $roomId 
			MERGE (u)-[:INVITED]->(r)`,
			map[string]any{
				"userId": userId,
				"roomId": roomId,
			})
		if err != nil {
			return nil, fmt.Errorf("execute create room query: %w", err)
		}

		stat, err := res.Consume(ctx)
		if err != nil {
			return nil, fmt.Errorf("error consume: %w", err)
		}

		if stat.Counters().RelationshipsCreated() != 1 {
			return nil, ErrUserNotExistOrAlreadyInvited
		}

		return nil, nil
	})

	if err != nil {
		return fmt.Errorf("error execute query at session: %w", err)
	}

	return nil
}

func (r *Repository) DetachUserFromRoom(ctx context.Context, userId, roomCode string) error {
	panic("dsdsdsdsdsds")
}
