package group_actions

import (
	mongodb_handler "que-sera-de-mi/src/config/db"
	group_model "que-sera-de-mi/src/models/group"

	"go.mongodb.org/mongo-driver/mongo"
)

type CreateGroupInput struct {
	Name string `json:"name"`
}

func CreateGroupHandler(groupInput CreateGroupInput) (*mongo.InsertOneResult, error) {
	connection, ctx, err := mongodb_handler.GetConnection()

	resultGroup, err := connection.Collection(group_model.GroupEntityName).InsertOne(ctx, groupInput)

	if err != nil {
		return nil, err
	}

	return resultGroup, nil
}
