package cmdline

import "flag"

const (
	settingsFileFlagName string = "s"
	settingsFileFlagHint string = `name of settings file

example settings file:
{
  "api_host": "127.0.0.1",
  "api_port": "8080",
  "db_host": "10.10.10.10",
  "db_port": "5432",
  "db_name": "database name",
  "db_user_name": "user name",
  "db_password": "user password"
}
`
)

func GetSettingsFileName() *string {
	configFile := flag.String(settingsFileFlagName, "", settingsFileFlagHint)
	flag.Parse()
	return configFile
}

func PrintUsage() {
	flag.Usage()
}
