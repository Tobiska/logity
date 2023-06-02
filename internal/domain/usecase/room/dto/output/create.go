package output

type CreateRoomOutputDto struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Tag        string `json:"tag"`
	CountUsers int    `json:"count_users"`
}
