package dto

import "time"

type (
	FindFilter struct {
		Code       string
		Tag        string
		Name       string
		CreatorId  string
		MembersIds []string
		StartDate  time.Time
		EndDate    time.Time
	}
)
