package guardduty

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stangirard/yatas/internal/yatas"
)

func RunChecks(wa *sync.WaitGroup, s aws.Config, c *yatas.Config, queue chan []yatas.Check) {

	var checkConfig yatas.CheckConfig
	checkConfig.Init(s, c)
	var checks []yatas.Check
	guardyDetectors := GetDetectors(checkConfig.ConfigAWS)
	go yatas.CheckTest(checkConfig.Wg, c, "AWS_GDT_001", CheckIfGuarddutyEnabled)(checkConfig, "AWS_GDT_001", guardyDetectors)
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
