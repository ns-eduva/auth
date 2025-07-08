package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nsevenpack/mignosql"
)

var CreateRolesCollection = mignosql.Migration{
	Name: "20250708120000_create_roles_collection",

	Up: func(db *mongo.Database) error {
		ctx := context.Background()

		validator := bson.M{
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"name", "description", "permission_ids"},
				"properties": bson.M{
					"name": bson.M{
						"bsonType":   "string",
						"minLength":  1,
						"description": "Nom du rôle",
					},
					"description": bson.M{
						"bsonType":   "string",
						"minLength":  1,
						"description": "Description du rôle",
					},
					"permission_ids": bson.M{
						"bsonType": "array",
						"items": bson.M{
							"bsonType": "objectId",
						},
						"description": "Liste des permissions associées au rôle",
					},
				},
			},
		}

		opts := options.CreateCollection().SetValidator(validator)
		err := db.CreateCollection(ctx, "roles", opts)
		if err != nil {
			if !mongo.IsDuplicateKeyError(err) {
				return err
			}
		}

		indexModel := mongo.IndexModel{
			Keys:    bson.M{"name": 1},
			Options: options.Index().SetUnique(true).SetName("idx_unique_role_name"),
		}

		_, err = db.Collection("roles").Indexes().CreateOne(ctx, indexModel)
		return err
	},
}
