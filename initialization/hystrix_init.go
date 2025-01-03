package Init

import (
	"github.com/afex/hystrix-go/hystrix"
)

func HyInit() {
	hystrix.ConfigureCommand("QueryHystrix", hystrix.CommandConfig{
		Timeout:                3000,
		MaxConcurrentRequests:  500,
		ErrorPercentThreshold:  60,
		SleepWindow:            5000,
		RequestVolumeThreshold: 30,
	})
}
