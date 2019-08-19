package api

import (
	"encoding/base64"
	"github.com/go-resty/resty"
	c "_context"
)

var Api api

func Init() *api{
	Api.initApiClient()
	return &Api
}

type api struct{
	client *resty.Client
	apiAuthTokenBase64 string
}

func (a *api) initApiClient(){
	a.client = resty.New()
	a.apiAuthTokenBase64 = base64.StdEncoding.EncodeToString([]byte(c.Connectivity.ApiAuthTokenProperty()))
}