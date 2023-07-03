package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	MovieName   string `json:"moviename" bson:"moviename"`
	ReleaseYear int    `json:"releaseyear" bson:"releaseyear"`
	DirectedBy  string `json:"directedby" bson:"directedby"`
	Genre       string `json:"genre" bson:"genre"`
}
