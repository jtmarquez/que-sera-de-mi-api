package group_model

import "go.mongodb.org/mongo-driver/bson/primitive"

const GroupEntityName = "group"

type Group struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}
