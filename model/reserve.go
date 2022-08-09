package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Table struct {
	Id          string `json:"Id"`
	MaxPosition int    `json:"maxPosition"`
}

type Reserve struct {
	Id          *primitive.ObjectID `bson:"_id,omitempty"`
	DateTime    string              `json:"datetime" bson:"datetime"`
	TotalPeople int64               `json:"totalPeople" bson:"total_people"`
	PhoneNumber string              `json:"phoneNumber" bson:"phone_number"`
	TableId     string              `json:"tableId" bson:"table_id"`
}
