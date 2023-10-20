package models

import "time"

type Cliente struct {
	Id   int64    `json:"clientId"`
	Data Info     `json:"inform"`
	Date InfoData `json:"date"`
}

type Info struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type InfoData struct {
	Created time.Time `json:"created"`
}

type Null struct {
}
