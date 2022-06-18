package model

import "time"

type HeartBeat struct {
	Id          int64     `json:"id"`
	Desc        string    `json:"desc"`
	Extra       string    `json:"extra"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
