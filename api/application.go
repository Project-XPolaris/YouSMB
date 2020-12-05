package api

import "github.com/allentom/haruka"

func RunWebApi(addr string) {
	e := haruka.NewEngine()
	e.Router.AddHandler("/config", getConfigHandler)
	e.Router.AddHandler("/config/global", getGlobalConfigHandler)
	e.Router.AddHandler("/folders/add", addFolderConfig)
	e.Router.AddHandler("/folders/remove", removeFolderConfig)
	e.Router.AddHandler("/folders/update", updateFolderConfig)
	e.RunAndListen(addr)
}
