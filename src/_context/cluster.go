package _context

import (
	"fmt"
	"strings"
)

var Cluster cluster

func InitClusterContext(clusterName, projectName, clusterId string) *cluster{

	Cluster = cluster{
		ClusterName: clusterName,
		ProjectName: projectName,
		ClusterId: clusterId,
	}

	return &Cluster
}

func (c *cluster) InitProjectContext(projectId string){
	c.ProjectId = projectId
}

func (c *cluster) InitHelmAppContext(appName string){
	c.HelmAppId = fmt.Sprintf("%s:%s",
		strings.Split(c.ProjectId, ":")[1],
		appName,
		)
}

type cluster struct {
	ClusterName string
	ProjectName string

	ClusterId string
	ProjectId string
	HelmAppId string
}