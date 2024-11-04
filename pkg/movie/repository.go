/*
* This file contains a repository responsible for inserting data
* into a history, which represents the history of views in the
* MongoDB database.
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

package movie

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository struct {
	Database *mongo.Database
	CTX      context.Context
}

func (movie *MovieRepository) InsertHistory(userId string, movieId string) error {
	ctx := movie.CTX
	collection := movie.Database.Collection("history")
	data := MovieHistory{
		UserID:  userId,
		MovieID: movieId,
	}

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
