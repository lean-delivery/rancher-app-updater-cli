package api

import (
	"_context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func (a *api) HelmAppGetInfo() (helmAppInfo helmAppGetInfoResponse){

	apiEndpoint := _context.Connectivity.GetApiEndpoint("helmAppInfo", map[string]string{
		"CLUSTER_ID": _context.Cluster.ClusterId,
		"PROJECT_ID": _context.Cluster.ProjectId,
		"APP_NAME": _context.Cluster.HelmAppId,
	})

	resp, err := a.get(apiEndpoint)

	if err == nil {
		if err := json.Unmarshal(resp.Body(), &helmAppInfo); err != nil {
			log.Errorf("Error parsing response body into AppInfo struct: %v", err)
		}
	} else {
		log.Errorf("Error sending GET request: %v", resp.Request)
	}

	return
}