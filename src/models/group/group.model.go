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

func createGroup(ctx context.Context, db mongodb_handler.DbHandler, group Group) error {
	_, err := db.Collection(GroupEntityName).InsertOne(ctx, group)

	if err != nil {
		return err
	}

	return nil
}

func CreateGroupHandler(group Group) error {
	var handler mongodb_handler.DbCommandCallback = func(ctx context.Context, db mongodb_handler.DbHandler) error {
		return createGroup(ctx, db, group)
	}

	err := mongodb_handler.ExecuteDbCommand(handler)

	return err
}
