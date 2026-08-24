// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-present Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package handler

import (
	"code.vikunja.io/api/pkg/web"
)

// WebHandler defines the webhandler object
// This does web stuff, aka returns json etc. Uses CRUDable Methods to get the data
type WebHandler struct {
	EmptyStruct func() CObject
}

// CObject is the definition of our object, holds the structs
type CObject interface {
	web.CRUDable
	web.Permissions
}

// maxPermissionClearer is implemented by models carrying the requesting user's
// permission in their response body. The generic pipeline never resolves that
// field for the body (v1 reports it in the x-max-permission header, v2 sets it
// per handler), and its zero value is a real permission (read), so an untouched
// field would claim read-only access instead of "not computed".
type maxPermissionClearer interface {
	ClearMaxPermission()
}

func clearMaxPermission(obj CObject) {
	if c, is := obj.(maxPermissionClearer); is {
		c.ClearMaxPermission()
	}
}
