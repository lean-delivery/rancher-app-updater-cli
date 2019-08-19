package api

import "time"

type helmAppGetInfoResponse struct {
	Actions struct {
		Rollback string `json:"rollback"`
		Upgrade  string `json:"upgrade"`
	} `json:"actions"`
	Annotations struct {
		LifecycleCattleIoCreateHelmControllerCQd5Hv string `json:"lifecycle.cattle.io/create.helm-controller_c-qd5hv"`
	} `json:"annotations"`
	Answers map[string]string `json:"answers"`
	AppRevisionID string `json:"appRevisionId"`
	BaseType      string `json:"baseType"`
	Conditions    []struct {
		LastUpdateTime time.Time `json:"lastUpdateTime,omitempty"`
		Status         string    `json:"status"`
		Type           string    `json:"type"`
	} `json:"conditions"`
	Created    time.Time `json:"created"`
	CreatedTS  int64     `json:"createdTS"`
	CreatorID  string    `json:"creatorId"`
	ExternalID string    `json:"externalId"`
	ID         string    `json:"id"`
	Labels     struct {
		CattleIoCreator string `json:"cattle.io/creator"`
	} `json:"labels"`
	Links struct {
		Remove   string `json:"remove"`
		Revision string `json:"revision"`
		Self     string `json:"self"`
		Update   string `json:"update"`
	} `json:"links"`
	MultiClusterAppID    interface{} `json:"multiClusterAppId"`
	Name                 string      `json:"name"`
	NamespaceID          interface{} `json:"namespaceId"`
	ProjectID            string      `json:"projectId"`
	Prune                bool        `json:"prune"`
	State                string      `json:"state"`
	TargetNamespace      string      `json:"targetNamespace"`
	Transitioning        string      `json:"transitioning"`
	TransitioningMessage string      `json:"transitioningMessage"`
	Type                 string      `json:"type"`
	UUID                 string      `json:"uuid"`
}