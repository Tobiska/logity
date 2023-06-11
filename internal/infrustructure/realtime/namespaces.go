package realtime

import "fmt"

type ChannelNamespace string

const (
	RoomNamespace      ChannelNamespace = "room"
	RoomsUserNamespace ChannelNamespace = "rooms_user"
)

func GenerateRoomChannel(roomId string) string {
	return fmt.Sprintf("%s:%s", RoomNamespace, roomId)
}

func GenerateRoomsUserChannel(userId string) string {
	return fmt.Sprintf("%s:%s", RoomNamespace, userId)
}
