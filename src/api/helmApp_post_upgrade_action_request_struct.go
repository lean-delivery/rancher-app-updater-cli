package api

type HelmAppPostUpgradeActionRequest struct {
	ExternalId string	`json:"externalId"`
	Answers map[string]string		`json:"answers"`
	ValuesYaml string	`json:"valuesYaml"`
	ForceUpgrade bool	`json:"forceUpgrade"`
}