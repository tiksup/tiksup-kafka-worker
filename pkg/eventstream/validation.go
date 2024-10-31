/*
* This file contains a function in charge of validating id
* for the mongodb database
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

import "go.mongodb.org/mongo-driver/bson/primitive"

func IsValidObjectID(s string) string {
	objID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		randomId := "c5ecece22fb00fa4da29cf01"
		return randomId
	}
	return objID.Hex()
}
