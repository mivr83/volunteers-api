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

	// check cmd line
	configFile := cmdline.GetSettingsFileName()
	if *configFile == "" {
		cmdline.PrintUsage()
		return
	}

	// process config file
	connSettings, err := settings.LoadFromFile(configFile)
	if err != nil {
		log.Printf("failed to process settings file -> %s\n", err)
		cmdline.PrintUsage()
		return
	}

	// connect to DB
	if dbErr := db.OpenDbConnection(connSettings); dbErr != nil {
		log.Printf("failed to connect to database -> %s", dbErr)
		return
	}

	// ignore db close error
	defer db.PostgresDb.Close()

	// start api service
	ims := session.Create()
	restapi.InitAndServe(connSettings, ims)
}
