package api

import (
	"github.com/go-resty/resty"
)

func (a *api) get(apiEndpoint string) (resp *resty.Response, err error) {

	resp, err = a.client.
		R().
		SetHeader("Authorization", "Basic "+a.apiAuthTokenBase64).
		EnableTrace().
		Get(apiEndpoint)

	return
}