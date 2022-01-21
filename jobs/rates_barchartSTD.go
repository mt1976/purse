package jobs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	//"github.com/mt1976/common"

	tools "github.com/mt1976/mwtgostringtools"
	application "github.com/mt1976/purse/application"
	core "github.com/mt1976/purse/core"
	dao "github.com/mt1976/purse/dao"
	dm "github.com/mt1976/purse/datamodel"
	logs "github.com/mt1976/purse/logs"
	cron "github.com/robfig/cron/v3"
)

const constRateLen = 8

var funcName = ""

type fxRate struct {
	ccyPair   string
	bidRate   float64
	bidString string
	askRate   float64
	askString string
}

type fxRateCard struct {
	fxRates []fxRate
}

func BarchartSTD_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "METRICS_BARCHART_FX"
	j.Name = "METRICS_BARCHART_FX"
	j.Period = "*/10 7-19 * * 1-5"
	j.Description = "Update FX Spot rate from barchart.com"
	j.Type = core.Aquirer
	return j
}

func BarchartSTD_Register(c *cron.Cron) {
	application.Schedule_Register(BarchartSTD_Job())
	c.AddFunc(BarchartSTD_Job().Period, func() { BarchartSTD_Run() })
}

// RunJobRollover is a Rollover function
func BarchartSTD_Run() {
	//logs.StartJob(RatesFXSpot_Job().Name)

	logs.StartJob(BarchartSTD_Job().Name)
	var message string
	/// CONTENT STARTS

	noItems, symbolList, err := dao.Symbol_GetListBySource("BARCHART_STD")
	if err != nil {
		log.Println(err.Error())
	}

	baseURI := ""

	_, ds, _ := dao.DataSource_GetByID("BARCHART_STD")
	if ds.URI != "" {
		baseURI = ds.URI
	}
	var rateCard fxRateCard

	for i := 0; i < noItems; i++ {
		//fmt.Printf("%s\n", symbolList[i])
		symbolData := symbolList[i]
		rateCard.fxRates = append(rateCard.fxRates, getFXrate(symbolData, baseURI))
	}

	if err != nil {
		message = err.Error()
	}

	/// CONTENT ENDS
	application.Schedule_Update(BarchartSTD_Job(), message)
	logs.EndJob(BarchartSTD_Job().Name)
}

func getFXrate(inSymbol dm.Symbol, baseURI string) fxRate {

	inCCYpair := inSymbol.Identifier
	thisRate := fxRate{}
	thisRate.ccyPair = inCCYpair
	thisPair := "%5E" + inCCYpair
	url := fmt.Sprintf(baseURI, thisPair)
	logs.Accessing(url)
	//fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)
	inString := string(html)
	//fmt.Println(inString)
	searchString := "\"bidPrice\":\""
	searchString2 := "\",\"askPrice\":\""
	searchString3 := "\",\"bidSize\":\""

	bidPriceStart := tools.FindStartPos(inString, searchString)
	bidPriceStop := tools.FindPos(inString, searchString2)
	askPriceStart := tools.FindStartPos(inString, searchString2)
	askPriceStop := tools.FindPos(inString, searchString3)
	if askPriceStop == -1 {
		askPriceStop = askPriceStart + 7
	}

	thisRate.bidRate, _ = strconv.ParseFloat(inString[bidPriceStart:bidPriceStop], 64)

	thisRate.askRate, _ = strconv.ParseFloat(inString[askPriceStart:askPriceStop], 64)

	var symb dm.SymbolMetrics
	symb.Symbol = inSymbol.Code
	symb.URI = url

	dps := "14"

	rawrd := "%." + dps + "f"

	symb.RawBid = fmt.Sprintf(rawrd, thisRate.bidRate)
	symb.RawOffer = fmt.Sprintf(rawrd, thisRate.askRate)
	symb.RawPrice = fmt.Sprintf(rawrd, (thisRate.bidRate+thisRate.askRate)/2)

	if inSymbol.DPS != "" {
		dps = inSymbol.DPS
	}
	rd := "%." + dps + "f"
	fmt.Printf("rd: %v\n", rd)
	symb.Bid = fmt.Sprintf(rd, thisRate.bidRate)
	symb.Offer = fmt.Sprintf(rd, thisRate.askRate)
	symb.Price = fmt.Sprintf(rd, (thisRate.bidRate+thisRate.askRate)/2)

	fmt.Printf("symb: %v\n", symb)

	dao.SymbolMetrics_StoreSystem(symb)
	dao.SymbolMetricsHistory_StoreRefresh(symb)

	//	RatesDataStorePut(ratesData)
	return thisRate
}

// var wg sync.WaitGroup

// func getASYCFXrate(inCCYpair string, rateChan chan fxRate) {
// 	defer wg.Done()

// 	thisRate := fxRate{}
// 	thisRate.ccyPair = inCCYpair
// 	thisPair := "%5E" + inCCYpair
// 	url := fmt.Sprintf("https://www.barchart.com/forex/quotes/%s/overview", thisPair)
// 	fmt.Printf("HTML code of %s ...\n", url)
// 	resp, err := http.Get(url)
// 	// handle the error if there is one
// 	if err != nil {
// 		log.Println(err, inCCYpair)
// 		panic(err)
// 	}
// 	// do this now so it won't be forgotten
// 	defer resp.Body.Close()
// 	// reads html as a slice of bytes
// 	html, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println(err, inCCYpair)
// 		panic(err)
// 	}
// 	// show the HTML code as a string %s
// 	//fmt.Printf("%s\n", html)
// 	inString := string(html)

// 	searchString := "\"bidPrice\":\""
// 	searchString2 := "\",\"askPrice\":\""
// 	searchString3 := "\",\"bidSize\":\""

// 	bidPriceStart := tools.FindStartPos(inString, searchString)
// 	bidPriceStop := tools.FindPos(inString, searchString2)
// 	askPriceStart := tools.FindStartPos(inString, searchString2)
// 	askPriceStop := tools.FindPos(inString, searchString3)
// 	if askPriceStop == -1 {
// 		askPriceStop = askPriceStart + 7
// 	}

// 	thisRate.bidRate, _ = strconv.ParseFloat(inString[bidPriceStart:bidPriceStop], 64)
// 	thisRate.bidString = tools.TruncateString(inString[bidPriceStart:bidPriceStop], constRateLen)
// 	//fmt.Println("bidPrice=", bidPrice)

// 	thisRate.askRate, _ = strconv.ParseFloat(inString[askPriceStart:askPriceStop], 64)
// 	thisRate.askString = tools.TruncateString(inString[askPriceStart:askPriceStop], constRateLen)
// 	//fmt.Println("askPrice=", askPrice)

// 	fmt.Printf("thisRate: %v\n", thisRate)

// 	rateChan <- thisRate

// 	//log.Println(thisRate.ccyPair, bidPriceStart, askPriceStart, askPriceStop, thisRate.bidRate, thisRate.askRate)

// }

// func buildRateCard() {
// 	var rateCard fxRateCard
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("AUDUSD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("EURGBP"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("EURJPY"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("EURUSD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("GBPUSD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("NZDUSD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDCAD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDCHF"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDHKD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDJPY"))
// 	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDSGD"))
// }

// func buldFXRVRates(rateCard fxRateCard) string {
// 	outputString := ""
// 	for _, s := range rateCard.fxRates {
// 		//fmt.Println(i, s)
// 		//fmt.Println(i, s)
// 		//fmt.Sprint(s)
// 		abc := fmt.Sprintf("s%ssptD%11sD%11s", s.ccyPair, s.bidString, s.askString)
// 		//fmt.Println(abc)
// 		outputString = outputString + abc + "\n"
// 	}
// 	//log.Println("\n", outputString)
// 	return outputString
// }

// func deliverRVData(name string, record string) {

// 	f, err := os.Create(name)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	_, err2 := f.WriteString(record)

// 	if err2 != nil {
// 		log.Fatal(err2)
// 	}

// }
