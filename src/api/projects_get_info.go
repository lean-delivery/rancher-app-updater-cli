package api

import (
	"_context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func (a *api) ProjectsGetInfo(clusterId, projectName string) (projectsInfo ProjectsGetInfoResponse){

	apiEndpoint := _context.Connectivity.GetApiEndpoint("projectsInfo", map[string]string{
		"CLUSTER_ID": clusterId,
		"PROJECT_NAME": projectName,
	})

	resp, err := a.get(apiEndpoint)

	if err == nil {
		if err := json.Unmarshal(resp.Body(), &projectsInfo); err != nil {
			log.Errorf("Error parsing response body for ProjectsInfo struct: %v", err)
		}
	} else {
		log.Errorf("Error sending GET request: %v", resp.Request)
	}

	return
}