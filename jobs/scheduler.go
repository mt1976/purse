package jobs

import (
	"fmt"
	"runtime"
	"strings"

	core "github.com/mt1976/purse/core"
	"github.com/mt1976/purse/logs"

	cron "github.com/robfig/cron/v3"
)

//CONST_CONFIG_FILE is cheese

// Start Initialises the Required Jobs
func Start() {
	//funcName = "Start"
	//logit("JobHandler", "CRON SCHEDULER")
	// Verbose c := cron.New(cron.WithLogger(
	//	cron.VerbosePrintfLogger(log.New(os.Stdout, "", log.LstdFlags))))
	c := cron.New()

	// Insert Jobs Here
	//var period string
	//var runType string

	//logit("info", c.Location().String())

	//logit("cron Locale", c.Location().String())

	HeartBeat_Register(c)

	if !core.IsChildInstance {
		//		RatesFXSpot_Register(c)
	}

	SessionHouseKeeping_Register(c)

	BarchartSTD_Register(c)
	BarchartCrypto_Register(c)
	StaticRates_Register(c)
	MetricsSnapshot_Register(c)

	c.Start()
	//fmt.Println(len(c.Entries()))
	//for i, e := range c.Entries() {
	//	fmt.Print(i)
	//		fmt.Printf("e: %v\n", e)
	//		fmt.Println(e.WrappedJob)
	//	}
}

func logit(actionType string, data string) {
	_, caller, _, _ := runtime.Caller(1)
	outcall := strings.Split(caller, "/")
	depth := len(outcall) - 1
	depth2 := depth - 1
	//log.Println(len(outcall), depth, depth2)
	callerName := outcall[depth2] + "/" + outcall[depth]
	op := fmt.Sprintf("%v '%v' {%v}", actionType, data, callerName)
	logs.Message("Scheduler", op)
}
