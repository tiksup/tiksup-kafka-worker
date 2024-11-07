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

type UserInfo struct {
	UserID         string       `json:"user_id"`
	MovieID        string       `json:"movie_id"`
	WatchingTime   float64      `json:"watching_time"`
	WatchingRepeat int          `json:"watching_repeat"`
	Interactions   Interactions `json:"interactions"`
	Next           bool         `json:"next"`
}

type Interactions struct {
	Genre       []string `json:"genre"`
	Protagonist string   `json:"protagonist"`
	Director    string   `json:"director"`
}
