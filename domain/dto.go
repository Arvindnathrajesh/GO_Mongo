package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

type LinkData struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Url      string             `json:"url" bson:"url,omitempty"`
	ShortUrl string             `json:"shortUrl" bson:"shortUrl,omitempty"`
	UseCase  string             `json:"useCase" bson:"useCase,omitempty"`
}

type UserLinkData struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Url        string             `json:"url" bson:"url,omitempty"`
	UserPhone  string             `json:"userPhone" bson:"userPhone,omitempty"`
	ClickCount int                `json:"clickCount" bson:"clickCount,omitempty"`
}
