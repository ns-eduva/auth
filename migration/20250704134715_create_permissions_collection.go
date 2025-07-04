package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nsevenpack/mignosql"
)

var CreatePermissionsCollection = mignosql.Migration{
	Name: "20250704120000_create_permissions_collection",

	Up: func(db *mongo.Database) error {
		ctx := context.Background()

		validator := bson.M{
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"name", "description"},
				"properties": bson.M{
					"name": bson.M{
						"bsonType":   "string",
						"minLength":  1,
						"description": "Nom de la permission",
					},
					"description": bson.M{
						"bsonType":   "string",
						"minLength":  1,
						"description": "Description de la permission",
					},
				},
			},
		}

		opts := options.CreateCollection().SetValidator(validator)
		err := db.CreateCollection(ctx, "permissions", opts)
		if err != nil {
			if !mongo.IsDuplicateKeyError(err) {
				return err
			}
		}

		indexModel := mongo.IndexModel{
			Keys:    bson.M{"name": 1},
			Options: options.Index().SetUnique(true).SetName("idx_unique_permission_name"),
		}

		_, err = db.Collection("permissions").Indexes().CreateOne(ctx, indexModel)
		return err
	},
}
