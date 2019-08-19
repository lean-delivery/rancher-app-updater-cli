package api

import (
	"github.com/go-resty/resty"
)

func (a *api) post (apiEndpoint, body string) (resp *resty.Response, err error) {

	resp, err = a.client.
		R().
		SetHeader("Authorization", "Basic "+a.apiAuthTokenBase64).
		SetBody(body).
		EnableTrace().
		Post(apiEndpoint)

	return
}