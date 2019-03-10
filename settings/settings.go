package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConnSettings struct {
	ApiHost    string `json:"api_host"`
	ApiPort    string `json:"api_port"`
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbName     string `json:"db_name"`
	DbUserName string `json:"db_user_name"`
	DbPassword string `json:"db_password"`
}

// loads .json settings file from disk and parse it, also checks if all
// required fields are presents
func LoadFromFile(fileName *string) (*ConnSettings, error) {
	fileData, err := ioutil.ReadFile(*fileName)
	if err != nil {
		return nil, err
	} else {
		connSet := ConnSettings{}
		if jsonErr := json.Unmarshal(fileData, &connSet); jsonErr != nil {
			return nil, jsonErr
		}

		if connSet.HaveMissingFields() {
			return nil, fmt.Errorf("missing field(s) in settings file")
		}

		return &connSet, nil
	}
}

// if ConnSettings misses required or mandatory fields returns true otherwise false
func (cs *ConnSettings) HaveMissingFields() bool {
	if cs.ApiHost == "" ||
		cs.ApiPort == "" ||
		cs.DbHost == "" ||
		cs.DbName == "" ||
		cs.DbPassword == "" ||
		cs.DbPort == "" ||
		cs.DbUserName == "" {
		return true
	}

	return false
}

// returns host address with port in format ApiHost:ApiPort
func (cs *ConnSettings) GetApiHostWithPort() string {
	return cs.ApiHost + ":" + cs.ApiPort
}
