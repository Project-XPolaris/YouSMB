package api

import "github.com/allentom/haruka"

func AbortErrorWithStatus(err error, context *haruka.Context, status int) {
	context.JSONWithStatus(haruka.JSON{
		"success": false,
		"reason":  err.Error(),
	}, status)
}
