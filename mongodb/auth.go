package mongodb

import (
	"context"
	"time"

	"github.com/silco-dev/vander/structs"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) GetUser(token string) (*structs.User, error) {
	var user *structs.User

	col := db.client.Database("vander").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := bson.M{
		"token": token,
	}

	cursor, err := col.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		cursor.Decode(&user)
	}
	return user, err
}
