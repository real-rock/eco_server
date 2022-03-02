package request

import (
	"time"
)

type RegisterRequest struct {
	Email    string    `json:"email" example:"example@economicus.kr"`
	Password string    `json:"password" example:"some password"`
	Name     string    `json:"name" example:"user name"`
	Nickname string    `json:"nickname" example:"user nickname"`
	Birth    time.Time `time_format:"2006-01-02T15:04:05.000Z" json:"birth,omitempty" example:"2016-03-31T00:00:000.Z"`
}
