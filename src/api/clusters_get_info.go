package api

import (
	"_context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func (a *api) ClustersGetInfo(clusterName string) (clustersInfo ClustersGetInfoResponse){

	apiEndpoint := _context.Connectivity.GetApiEndpoint("clustersInfo", map[string]string{
		"CLUSTER_NAME": clusterName,
	})

	resp, err := a.get(apiEndpoint)

	if err == nil {
		if err := json.Unmarshal(resp.Body(), &clustersInfo); err != nil {
			log.Errorf("Error parsing response body for ClustersInfo struct: %v", err)
		}
	} else {
		log.Errorf("Error sending GET request: %v\n%v", resp.Request, err)
	}

	return
}