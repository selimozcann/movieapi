package models

type Movie struct {
	Id          int    `json:"id"`
	MovieName   string `json:"moviename"`
	ReleaseYear int    `json:"releaseyear"`
	DirectedBy  string `json:"directedby"`
	Genre       string `json:"genre"`
}
