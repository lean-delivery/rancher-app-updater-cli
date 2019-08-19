# Update 

## How to use it

Preconditions

Helm charts should contain "container.image.tag"

Set following environment variables

* RANCHER_TOKEN
* RANCHER_BASE_URL
* RANCHER_API_VERSION
* K8S_CLUSTER_NAME
* K8S_PROJECT_NAME
* HELM_APP_NAME
* BUILD_NUMBER
* SERVICE_NAME
* IMAGE_NAME
* TAG_NAME

Run cli
```bash
$ charts-updater-cli
```

## Docker Images

docker pull leandelivery/rancher-app-updater-cli:latest

https://hub.docker.com/r/leandelivery/rancher-app-updater-cli
