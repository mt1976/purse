package jobs

import (
	"log"

	//"github.com/mt1976/common"

	application "github.com/mt1976/purse/application"
	core "github.com/mt1976/purse/core"
	dao "github.com/mt1976/purse/dao"
	dm "github.com/mt1976/purse/datamodel"
	logs "github.com/mt1976/purse/logs"
	cron "github.com/robfig/cron/v3"
)

func MetricsSnapshot_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "METRICS_SNAPSHOT"
	j.Name = "METRICS_SNAPSHOT"
	j.Period = "*/11 7-19 * * 1-5"
	j.Description = "Update FX Spot rate from barchart.com"
	j.Type = core.Aquirer
	return j
}

func MetricsSnapshot_Register(c *cron.Cron) {
	application.Schedule_Register(MetricsSnapshot_Job())
	c.AddFunc(MetricsSnapshot_Job().Period, func() { MetricsSnapshot_Run() })
}

// RunJobRollover is a Rollover function
func MetricsSnapshot_Run() {
	//logs.StartJob(RatesFXSpot_Job().Name)

	logs.StartJob(MetricsSnapshot_Job().Name)
	var message string
	/// CONTENT STARTS

	noItems, symbolList, err := dao.SymbolMetrics_GetList()
	if err != nil {
		log.Println(err.Error())
	}

	for i := 0; i < noItems; i++ {
		//fmt.Printf("%s\n", symbolList[i])
		symbolData := symbolList[i]
		dao.SymbolMetricsHistory_StoreSnapshot(symbolData)
	}

	if err != nil {
		message = err.Error()
	}

	/// CONTENT ENDS
	application.Schedule_Update(MetricsSnapshot_Job(), message)
	logs.EndJob(MetricsSnapshot_Job().Name)
}
