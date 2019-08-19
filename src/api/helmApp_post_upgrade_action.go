package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)


func (a *api) HelmAppPostUpgradeAction(apiEndpoint string, bodyStruct HelmAppPostUpgradeActionRequest) (statusCode int){

	body, err := json.Marshal(bodyStruct)

	if err != nil {
		log.Errorf("Error parsing body struct for POST request: %v", err)
		return
	}

	resp, err := a.post(apiEndpoint, string(body))

	if err != nil {
		log.Errorf("Error sending POST request: %v", resp.Request)
		return
	}

	statusCode = resp.StatusCode()

	return
}