package funcs

import (
	"fmt"
)

func CheckCollector() {

	output := make(map[string]bool)

	output["CpuMetrics  "] = len(CpuMetrics()) > 0
	output["MemMetrics  "] = len(MemMetrics()) > 0
	output["SwIfInMetrics "] = len(SwIfInMetrics()) > 0
	output["SwIfInSpeedPercentMetrics "] = len(SwIfInSpeedPercentMetrics()) > 0
        output["SwIfOutMetrics "] = len(SwIfOutMetrics()) > 0
        output["SwIfOutSpeedPercentMetrics "] = len(SwIfOutSpeedPercentMetrics()) > 0

	for k, v := range output {
		status := "fail"
		if v {
			status = "ok"
		}
		fmt.Println(k, "...", status)
	}
}
