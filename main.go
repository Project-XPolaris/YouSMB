package main

import (
	"log"
	"yousmb/api"
	"yousmb/application"
	"yousmb/smb"
)

func main() {
	err := application.LoadAppConfig()
	if err != nil {
		log.Fatalln(err)
	}
	err = smb.LoadConfig(application.Config.SmbConfigPath)
	if err != nil {
		log.Fatalln(err)
	}
	api.RunWebApi(application.Config.Addr)
}
