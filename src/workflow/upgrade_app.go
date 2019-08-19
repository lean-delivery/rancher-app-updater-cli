package workflow

import (
	c "_context"
	a "api"
	"fmt"
)

func UpgradeApp() string {

	fmt.Println("################################")
	fmt.Println("# STARTING APPLICATION UPGRADE #")
	fmt.Println("################################")

	helmAppInfo := a.Api.HelmAppGetInfo()

	helmAnswers := helmAppInfo.Answers

	answerKey := fmt.Sprintf("%s.container.image.tag", c.HelmApp.Service)
	answerValue := fmt.Sprintf("%s", c.HelmApp.Tag)

	if oldAnswerValue, ok := helmAnswers[answerKey]; ok {

		helmAnswers[answerKey] = answerValue

		fmt.Println("\n# CONFIGURATION DELTA #")
		fmt.Printf("KEY: %s\n", answerKey)
		fmt.Printf("DIFF: '%s' -> '%s'\n", oldAnswerValue, answerValue)

		upgradeEndpoint := helmAppInfo.Actions.Upgrade

		upgradeRequestBody := a.HelmAppPostUpgradeActionRequest{
			ExternalId:   helmAppInfo.ExternalID,
			Answers:      helmAnswers,
			ValuesYaml:   "",
			ForceUpgrade: false,
		}

		statusCode := a.Api.HelmAppPostUpgradeAction(upgradeEndpoint, upgradeRequestBody)

		if statusCode >= 200 && statusCode < 400 {
			return `
##################
# UPGRADE POSTED #
##################`
		} else {
			return `
##################
# UPGRADE FAILED #
##################`
		}
	} else {
		fmt.Printf("\nKEY: %s\n", answerKey)
		return `
####################
# NO SUCH HELM KEY #
####################
`
	}
}