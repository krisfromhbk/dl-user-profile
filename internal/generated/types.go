// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package generated

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	Dance_link_authScopes = "dance_link_auth.Scopes"
)

// GetUserInfoResponse defines model for GetUserInfoResponse.
type GetUserInfoResponse struct {
	Age               string      `json:"age"`
	City              string      `json:"city"`
	DanceExperience   string      `json:"danceExperience"`
	DanceStyle        string      `json:"danceStyle"`
	Media             []MediaItem `json:"media"`
	Name              string      `json:"name"`
	Note              *string     `json:"note"`
	ProfilePictureUrl string      `json:"profilePictureUrl"`
}

// GetUserInfoSelectItemsResponse defines model for GetUserInfoSelectItemsResponse.
type GetUserInfoSelectItemsResponse struct {
	Cities           []SlugTitleItem `json:"cities"`
	DanceExperiences []SlugTitleItem `json:"danceExperiences"`
	DanceLevels      []SlugTitleItem `json:"danceLevels"`
	DanceStyles      []SlugTitleItem `json:"danceStyles"`
	Statuses         []SlugTitleItem `json:"statuses"`
}

// MediaItem defines model for MediaItem.
type MediaItem struct {
	Id         int    `json:"id"`
	PreviewUrl string `json:"previewUrl"`
	Type       string `json:"type"`
	Url        string `json:"url"`
}

// SlugTitleItem defines model for SlugTitleItem.
type SlugTitleItem struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

// UpdateUserInfoRequest defines model for UpdateUserInfoRequest.
type UpdateUserInfoRequest = map[string]interface{}

// Version defines model for Version.
type Version = int64

// DeleteMediaJSONBody defines parameters for DeleteMedia.
type DeleteMediaJSONBody struct {
	Id int `json:"id"`
}

// UploadMediaMultipartBody defines parameters for UploadMedia.
type UploadMediaMultipartBody = openapi_types.File

// UpdateNoteJSONBody defines parameters for UpdateNote.
type UpdateNoteJSONBody struct {
	Note string `json:"note"`
}

// DeleteMediaJSONRequestBody defines body for DeleteMedia for application/json ContentType.
type DeleteMediaJSONRequestBody DeleteMediaJSONBody

// UploadMediaMultipartRequestBody defines body for UploadMedia for multipart/form-data ContentType.
type UploadMediaMultipartRequestBody = UploadMediaMultipartBody

// UpdateNoteJSONRequestBody defines body for UpdateNote for application/json ContentType.
type UpdateNoteJSONRequestBody UpdateNoteJSONBody

// UpdateUserInfoJSONRequestBody defines body for UpdateUserInfo for application/json ContentType.
type UpdateUserInfoJSONRequestBody = UpdateUserInfoRequest