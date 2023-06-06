package output

type CreateRoomOutputDto struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Tag        string `json:"tag"`
	CountUsers int    `json:"count_users"`
}
