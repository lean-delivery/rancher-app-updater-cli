package _context

import (
	"fmt"
	"router"
)

var Connectivity connectivity

func InitConnectivityContext(rancherUrl string, apiVersion string, apiAuthToken string) *connectivity {
	Connectivity = connectivity{
		RancherUrl: rancherUrl,
		ApiVersion: apiVersion,
		ApiAuthToken: apiAuthToken,
	}

	return &Connectivity
}

type connectivity struct {
	ApiVersion string
	RancherUrl string
	ApiAuthToken string
}

func (c *connectivity) GetApiEndpoint(context string, parameters map[string]string) string{
	apiContextPath := router.GetContextPath(context, parameters)
	return fmt.Sprintf("%s/v%s/%s", c.RancherUrl, c.ApiVersion, apiContextPath)
}

func (c *connectivity) ApiAuthTokenProperty() string{
	return c.ApiAuthToken
}