package api

import "github.com/allentom/haruka"

func AbortErrorWithStatus(err error, context *haruka.Context, status int) {
	context.Writer.Header().Set("ContentType", "application/json")
	context.Writer.WriteHeader(status)
	context.JSON(haruka.JSON{
		"success": false,
		"reason":  err.Error(),
	})
}
