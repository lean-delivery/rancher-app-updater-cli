package router

import (
	"strings"
)


var contextPathsTemplates = map[string]string{
	"helmAppInfo": "project/{PROJECT_ID}/apps/{APP_NAME}",
	"clustersInfo": "clusters/?name={CLUSTER_NAME}",
	"projectsInfo": "projects/?clusterId={CLUSTER_ID}&name={PROJECT_NAME}",
}


func GetContextPath(context string, parameters map[string]string) (res string){

	res = contextPathsTemplates[context]

	for param, value := range parameters{
		res = strings.ReplaceAll(res, "{" + param + "}", value)
	}

	return res
}