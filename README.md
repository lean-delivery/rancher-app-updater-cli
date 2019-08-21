# Update 

## How to use it

Preconditions

Helm charts should contain "${service_name}.container.image.tag"
Sample of the helm chart's `values.yaml` file:
```yaml=
yourApplication:
  name: app-name
  container:
    image:
      repository: registry/app-name
      tag: develop
      pullPolicy: IfNotPresent
```

Set the following environment variables

* RANCHER_TOKEN (i.e. token-1j90s21:0h12809h1pi2ohs12s912hp9812hep9128he1298eh12smd)
* RANCHER_BASE_URL
* RANCHER_API_VERSION
* K8S_CLUSTER_NAME
* K8S_PROJECT_NAME
* HELM_APP_NAME
* BUILD_NUMBER
* SERVICE_NAME
* IMAGE_NAME
* TAG_NAME

> When generating the API token in Rancher please leave parameter `scope` untouched.

Call the tool in CLI:
```bash
$ charts-updater-cli
```

## Docker Images

```sh
# https://hub.docker.com/r/leandelivery/rancher-app-updater-cli
docker pull leandelivery/rancher-app-updater-cli:latest
```