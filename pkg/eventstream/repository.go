/*
* This file contains functions in charge of inserting and
* updating data obtained from Kafka for the MongoDB database.
* Copyright (C) 2024-2025 jsusmachaca
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

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
	ctx := kafka.CTX
	collection := kafka.Collection

	ensureFields := bson.M{
		"$setOnInsert": bson.M{
			"preferences.genre_score":       []bson.M{},
			"preferences.protagonist_score": []bson.M{},
			"preferences.director_score":    []bson.M{},
		},
	}
	_, err := collection.UpdateOne(ctx, filter, ensureFields, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error ensuring fields exist: %v", err)
	}

	// Genre scores
	if err := updateGenreScores(data, filter, ctx, collection); err != nil {
		return err
	}
	// Protagonist scores
	if err := updateProtagonistScores(data, filter, ctx, collection); err != nil {
		return err
	}
	// Director scores
	if err := updateDirectorScores(data, filter, ctx, collection); err != nil {
		return err
	}

	return nil
}

func updateGenreScores(data KafkaData, filter bson.M, ctx context.Context, collection *mongo.Collection) error {
	updateIncGenre := bson.M{
		"$inc": bson.M{
			"preferences.genre_score.$[genre].score": data.Preferences.GenreScore[0].Score,
		},
	}

	arrayFiltersGenre := options.ArrayFilters{
		Filters: []any{
			bson.M{"genre.name": data.Preferences.GenreScore[0].Name},
		},
	}

	updateOptionsGenre := options.Update().SetArrayFilters(arrayFiltersGenre)
	resultsGenre, err := collection.UpdateOne(ctx, filter, updateIncGenre, updateOptionsGenre)
	if err != nil {
		return fmt.Errorf("error incrementing genre scores: %v", err)
	}

	if resultsGenre.ModifiedCount == 0 {
		updatePushGenre := bson.M{
			"$addToSet": bson.M{
				"preferences.genre_score": bson.M{
					"name":  data.Preferences.GenreScore[0].Name,
					"score": data.Preferences.GenreScore[0].Score,
				},
			},
		}

		_, err = collection.UpdateOne(ctx, filter, updatePushGenre)
		if err != nil {
			return fmt.Errorf("error pushing new genre scores: %v", err)
		}
	}
	return nil
}

func updateProtagonistScores(data KafkaData, filter bson.M, ctx context.Context, collection *mongo.Collection) error {
	updateIncProtagonist := bson.M{
		"$inc": bson.M{
			"preferences.protagonist_score.$[protagonist].score": data.Preferences.ProtagonistScore[0].Score,
		},
	}

	arrayFiltersProtagonist := options.ArrayFilters{
		Filters: []any{
			bson.M{"protagonist.name": data.Preferences.ProtagonistScore[0].Name},
		},
	}

	updateOptionsProtagonist := options.Update().SetArrayFilters(arrayFiltersProtagonist)
	resultsProtagonist, err := collection.UpdateOne(ctx, filter, updateIncProtagonist, updateOptionsProtagonist)
	if err != nil {
		return fmt.Errorf("error incrementing protagonist scores: %v", err)
	}

	if resultsProtagonist.ModifiedCount == 0 {
		updatePushProtagonist := bson.M{
			"$addToSet": bson.M{
				"preferences.protagonist_score": bson.M{
					"name":  data.Preferences.ProtagonistScore[0].Name,
					"score": data.Preferences.ProtagonistScore[0].Score,
				},
			},
		}

		_, err = collection.UpdateOne(ctx, filter, updatePushProtagonist)
		if err != nil {
			return fmt.Errorf("error pushing new protagonist scores: %v", err)
		}
	}
	return nil
}

func updateDirectorScores(data KafkaData, filter bson.M, ctx context.Context, collection *mongo.Collection) error {
	updateIncDirector := bson.M{
		"$inc": bson.M{
			"preferences.director_score.$[director].score": data.Preferences.DirectorScore[0].Score,
		},
	}

	arrayFiltersDirector := options.ArrayFilters{
		Filters: []any{
			bson.M{"director.name": data.Preferences.DirectorScore[0].Name},
		},
	}

	updateOptionsDirector := options.Update().SetArrayFilters(arrayFiltersDirector)
	resultsDirector, err := collection.UpdateOne(ctx, filter, updateIncDirector, updateOptionsDirector)
	if err != nil {
		return fmt.Errorf("error incrementing director scores: %v", err)
	}

	if resultsDirector.ModifiedCount == 0 {
		updatePushDirector := bson.M{
			"$addToSet": bson.M{
				"preferences.director_score": bson.M{
					"name":  data.Preferences.DirectorScore[0].Name,
					"score": data.Preferences.DirectorScore[0].Score,
				},
			},
		}

		_, err = collection.UpdateOne(ctx, filter, updatePushDirector)
		if err != nil {
			return fmt.Errorf("error pushing new director scores: %v", err)
		}
	}
	return nil
}
