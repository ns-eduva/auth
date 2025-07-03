package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nsevenpack/mignosql"
)

var CreateUsersCollection = mignosql.Migration{
	Name: "20250703180000_create_users_collection",

	Up: func(db *mongo.Database) error {
		ctx := context.Background()

		validator := bson.M{
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"email", "password"},
				"properties": bson.M{
					"email": bson.M{
						"bsonType":    "string",
						"description": "Email de l'utilisateur",
					},
					"password": bson.M{
						"bsonType":    "string",
						"description": "Mot de passe hashé",
					},
					"crm_id": bson.M{
						"bsonType":    "string",
						"description": "ID CRM optionnel",
					},
					"role_ids": bson.M{
						"bsonType":    "array",
						"description": "Liste des rôles",
					},
					"created_at": bson.M{
						"bsonType":    "date",
						"description": "Date de création",
					},
					"updated_at": bson.M{
						"bsonType":    "date",
						"description": "Date de mise à jour",
					},
				},
			},
		}

		opts := options.CreateCollection().SetValidator(validator)
		err := db.CreateCollection(ctx, "users", opts)
		if err != nil {
			if !mongo.IsDuplicateKeyError(err) {
				return err
			}
		}

		indexModel := mongo.IndexModel{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true).SetName("idx_unique_email"),
		}

		_, err = db.Collection("users").Indexes().CreateOne(ctx, indexModel)
		return err
	},
}
