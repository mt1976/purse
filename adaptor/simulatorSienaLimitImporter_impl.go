package adaptor

import (
	"strings"

	"github.com/mt1976/mwt-go-dev/core"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Simulator_SienaLimitImporter_ProcessResponse(filename string) error {
	//	logs.Success("StaticImport_Dispatch")
	logs.Success("SienaLimitImporter_ProcessResponse:" + filename)

	// tokenise the filename, get last element
	tokens := strings.Split(filename, "/")
	last := tokens[len(tokens)-1]

	//	uri := dm.Simulator_SienaFundsChecker_PathAction + "/?" + dm.Simulator_SienaFundsChecker_QueryString + "=" + last

	core.Notification_Normal("New Siena Limit Importer Response", last)
	var err error
	return err
}
