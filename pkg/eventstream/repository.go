package eventstream

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type KafkaRepository struct {
	Collection *mongo.Collection
	CTX        context.Context
}

func (kafka *KafkaRepository) UpdateUserInfo(data KafkaData) error {
	filter := bson.M{"user_id": data.UserID}

	ensureFields := bson.M{
		"$setOnInsert": bson.M{
			"preferences.genre_score":       []bson.M{},
			"preferences.protagonist_score": []bson.M{},
			"preferences.director_score":    []bson.M{},
		},
	}
	_, err := kafka.Collection.UpdateOne(kafka.CTX, filter, ensureFields, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error ensuring fields exist: %v", err)
	}

	update := bson.M{
		"$inc": bson.M{
			"preferences.genre_score.$[genre].score":             data.Preferences.GenreScore[0].Score,
			"preferences.protagonist_score.$[protagonist].score": data.Preferences.ProtagonistScore[0].Score,
			"preferences.director_score.$[director].score":       data.Preferences.DirectorScore[0].Score,
		},
	}

	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []any{
			bson.M{"genre.name": data.Preferences.GenreScore[0].Name},
			bson.M{"protagonist.name": data.Preferences.ProtagonistScore[0].Name},
			bson.M{"director.name": data.Preferences.DirectorScore[0].Name},
		},
	})

	results, err := kafka.Collection.UpdateOne(kafka.CTX, filter, update, arrayFilters)
	if err != nil {
		return err
	}
	if results.ModifiedCount == 0 {
		updatePush := bson.M{
			"$push": bson.M{
				"preferences.genre_score": bson.M{
					"name":  data.Preferences.GenreScore[0].Name,
					"score": data.Preferences.GenreScore[0].Score,
				},
				"preferences.protagonist_score": bson.M{
					"name":  data.Preferences.ProtagonistScore[0].Name,
					"score": data.Preferences.ProtagonistScore[0].Score,
				},
				"preferences.director_score": bson.M{
					"name":  data.Preferences.DirectorScore[0].Name,
					"score": data.Preferences.DirectorScore[0].Score,
				},
			},
		}

		_, err = kafka.Collection.UpdateOne(kafka.CTX, filter, updatePush)
		if err != nil {
			return err
		}
	}
	return nil
}
