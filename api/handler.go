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
	context.JSON(haruka.JSON{
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
	context.JSON(haruka.JSON{
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
	context.JSON(haruka.JSON{
		"result": "success",
	})
}

type AddUserRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var addUserHandler haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody AddUserRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusBadRequest)
		return
	}
	err = smb.DefaultUserManager.Create(requestBody.Username, requestBody.Password)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusBadRequest)
		return
	}
	context.JSON(haruka.JSON{
		"result": "success",
	})
}

var getUserList haruka.RequestHandler = func(context *haruka.Context) {
	userList := make([]SMBUserTemplate, 0)
	for _, user := range smb.DefaultUserManager.Users {
		template := SMBUserTemplate{Username: user.Username}
		userList = append(userList, template)
	}
	context.JSON(haruka.JSON{
		"users": userList,
	})
}

var removeUser haruka.RequestHandler = func(context *haruka.Context) {
	username := context.GetQueryString("username")
	err := smb.DefaultUserManager.RemoveUser(username)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusInternalServerError)
		return
	}
	context.JSON(haruka.JSON{
		"result": "success",
	})
}

var infoHandler haruka.RequestHandler = func(context *haruka.Context) {
	context.JSON(haruka.JSON{
		"name":   "YouSMB Service",
		"status": "running",
	})
}
