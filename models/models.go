package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ObjectID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID       string             `json:"uuid,omitempty" bson:"uuid,omitempty"`
	FullName string             `json:"fullName,omitempty" bson:"fullName,omitempty"`
	Address  string             `json:"address,omitempty" bson:"address,omitempty"`
}

type Product struct {
	ObjectID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID       string             `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Price    float64            `json:"price,omitempty" bson:"price,omitempty"`
}

type ErrorModel struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type ShoppingCart struct {
	ObjectID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId      string             `json:"userId,omitempty" bson:"userId,omitempty"`
	ProductIds  []string           `json:"productIds,omitempty" bson:"productIds,omitempty"`
	TotalAmount float64            `json:"totalAmount,omitempty" bson:"totalAmount,omitempty"`
}

type RemoveProduct struct {
	UserId     string   `json:"userId,omitempty" bson:"userId,omitempty"`
	ProductIds []string `json:"productIds,omitempty" bson:"productIds,omitempty"`
}
