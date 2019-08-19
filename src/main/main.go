package main

import (
	c "_context"
	"api"
	"fmt"
	"os"
	"workflow"
)

// TODO:
// 1) Implement over https://github.com/spf13/viper
// 2) Implement proper logging over https://github.com/sirupsen/logrus
var RANCHER_TOKEN =  os.Getenv("RANCHER_TOKEN")
var RANCHER_BASE_URL = os.Getenv("RANCHER_BASE_URL")
var RANCHER_API_VERSION = os.Getenv("RANCHER_API_VERSION")
var K8S_CLUSTER_NAME = os.Getenv("K8S_CLUSTER_NAME")
var K8S_PROJECT_NAME = os.Getenv("K8S_PROJECT_NAME")
var HELM_APP_NAME = os.Getenv("HELM_APP_NAME")
var BUILD_NUMBER = os.Getenv("BUILD_NUMBER")
var SERVICE_NAME =  os.Getenv("SERVICE_NAME")
var IMAGE_NAME =  os.Getenv("IMAGE_NAME")
var TAG_NAME = os.Getenv("TAG_NAME")


func main() {
	InitContext()

	upgradeStatus := workflow.UpgradeApp()

	fmt.Println(upgradeStatus)
}

func InitContext() {
	c.InitConnectivityContext(RANCHER_BASE_URL, RANCHER_API_VERSION, RANCHER_TOKEN)

	c.InitApplicationContext(
		SERVICE_NAME,
		IMAGE_NAME,
		fmt.Sprintf("%s-%s", TAG_NAME, BUILD_NUMBER),
	)

	api.Init()

	c.InitClusterContext(
		K8S_CLUSTER_NAME,
		K8S_PROJECT_NAME,
		api.Api.ClustersGetInfo(K8S_CLUSTER_NAME).Data[0].ID,
	)

	c.Cluster.InitProjectContext(
		api.Api.ProjectsGetInfo(c.Cluster.ClusterId, K8S_PROJECT_NAME).Data[0].ID,
	)

	c.Cluster.InitHelmAppContext(HELM_APP_NAME)
}