package model

type Customer struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
}