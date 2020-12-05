package api

import (
	"github.com/allentom/haruka"
	"net/http"
	"yousmb/smb"
)

// get config
var getConfigHandler haruka.RequestHandler = func(context *haruka.Context) {
	context.JSON(smb.Current)
}

// get global
var getGlobalConfigHandler haruka.RequestHandler = func(context *haruka.Context) {
	var globalSection *smb.Section
	for _, section := range smb.Current.Sections {
		if section.Name == "global" {
			globalSection = section
		}
	}

	context.JSON(globalSection)
}

// add folder
type SMBFolderRequestBody struct {
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties"`
}

var addFolderConfig haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody SMBFolderRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusBadRequest)
		return
	}
	err = smb.Current.AddFolder(requestBody.Name, requestBody.Properties)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusBadRequest)
		return
	}
	context.JSON(map[string]interface{}{
		"result": "success",
	})
}

// remove folder
var removeFolderConfig haruka.RequestHandler = func(context *haruka.Context) {
	name := context.GetQueryString("name")
	err := smb.Current.RemoveFolder(name)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusBadRequest)
		return
	}
	context.JSON(map[string]interface{}{
		"result": "success",
	})
}

// update folder
var updateFolderConfig haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody SMBFolderRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusBadRequest)
		return
	}
	err = smb.Current.UpdateFolder(requestBody.Name, requestBody.Properties)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusBadRequest)
		return
	}
	context.JSON(map[string]interface{}{
		"result": "success",
	})
}
