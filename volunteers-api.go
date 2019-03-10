package main

import (
	_ "github.com/lib/pq"
	"log"
	"volunteers-api/cmdline"
	"volunteers-api/db"
	"volunteers-api/restapi"
	"volunteers-api/session"
	"volunteers-api/settings"
)

func main() {

	configFile := cmdline.GetSettingsFileName()
	if *configFile == "" {
		cmdline.PrintUsage()
		return
	}

	connSettings, err := settings.LoadFromFile(configFile)
	if err != nil {
		log.Printf("failed to process settings file -> %s\n", err)
		cmdline.PrintUsage()
		return
	}

	if dbErr := db.OpenDbConnection(connSettings); dbErr != nil {
		log.Printf("failed to connect to database -> %s", dbErr)
		return
	}

	defer func() {
		if err := db.PostgresDb.Close(); err != nil {
			log.Printf("failed to close db connection -> %s", err)
		}
	}()

	ims := session.Create()
	restapi.InitAndServe(connSettings, ims)
}
