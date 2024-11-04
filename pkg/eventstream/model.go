/*
* This file contains templates that define the structure of Kafka messages
* in JSON format.
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

type KafkaData struct {
	UserID      string      `bson:"user_id" json:"user_id"`
	MovieID     string      `bson:"movie_id" json:"movie_id"`
	Preferences Preferences `bson:"preferences" json:"preferences"`
	Next        bool        `bson:"next" json:"next"`
}

type Preferences struct {
	GenreScore       []Score `bson:"genre_score" json:"genre_score"`
	ProtagonistScore []Score `bson:"protagonist_score" json:"protagonist_score"`
	DirectorScore    []Score `bson:"director_score" json:"director_score"`
}

type Score struct {
	Name  string  `bson:"name" json:"name"`
	Score float64 `bson:"score" json:"score"`
}
