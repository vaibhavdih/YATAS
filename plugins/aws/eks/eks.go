package eks

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/stangirard/yatas/internal/yatas"
)

func RunChecks(wa *sync.WaitGroup, s aws.Config, c *yatas.Config, queue chan []yatas.Check) {
	var checkConfig yatas.CheckConfig
	checkConfig.Init(s, c)
	var checks []yatas.Check
	svc := eks.NewFromConfig(s)
	clusters := GetClusters(svc)
	go yatas.CheckTest(checkConfig.Wg, c, "AWS_EKS_001", CheckIfLoggingIsEnabled)(checkConfig, clusters, "AWS_EKS_001")
	go func() {
		for t := range checkConfig.Queue {
			checks = append(checks, t)
			if c.CheckProgress.Bar != nil {
				c.CheckProgress.Bar.Increment()
				time.Sleep(time.Millisecond * 100)
			}
			checkConfig.Wg.Done()

		}
	}()
	checkConfig.Wg.Wait()
	queue <- checks
}
