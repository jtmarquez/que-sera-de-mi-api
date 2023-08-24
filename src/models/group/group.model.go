package group

import (
	"context"
	mongodb_handler "que-sera-de-mi/src/config/db"
)

const GroupEntityName = "group"

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var GroupData = []Group{
	{ID: "1", Name: "Group Name"},
}

func createGroup(ctx context.Context, db mongodb_handler.DbHandler, group Group) (interface{}, error) {
	resultGroup, err := db.Collection(GroupEntityName).InsertOne(ctx, group)

	if err != nil {
		return Group{}, err
	}

	return resultGroup, nil
}

func CreateGroupHandler(group Group) (interface{}, error) {
	var handler = func(ctx context.Context, db mongodb_handler.DbHandler) (interface{}, error) {
		return createGroup(ctx, db, group)
	}

	result, err := mongodb_handler.ExecuteDbCommand(handler)

	return result, err
}
