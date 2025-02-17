package apigateway

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/stangirard/yatas/internal/logger"
	"github.com/stangirard/yatas/internal/yatas"
)

func CheckIfStagesProtectedByAcl(checkConfig yatas.CheckConfig, stages []types.Stage, testName string) {
	logger.Info(fmt.Sprint("Running ", testName))
	var check yatas.Check
	check.InitCheck("ApiGateways are protected by an ACL", "Check if all stages are protected by ACL", testName)
	for _, stage := range stages {
		if stage.WebAclArn != nil && *stage.WebAclArn != "" {
			Message := "Stage " + *stage.StageName + " is protected by ACL"
			result := yatas.Result{Status: "OK", Message: Message, ResourceID: *stage.StageName}
			check.AddResult(result)
		} else {
			Message := "Stage " + *stage.StageName + " is not protected by ACL"
			result := yatas.Result{Status: "FAIL", Message: Message, ResourceID: *stage.StageName}
			check.AddResult(result)
		}
	}
	checkConfig.Queue <- check
}
