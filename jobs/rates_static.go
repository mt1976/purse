package jobs

import (
	"fmt"
	"log"
	"strconv"

	//"github.com/mt1976/common"

	application "github.com/mt1976/purse/application"
	core "github.com/mt1976/purse/core"
	dao "github.com/mt1976/purse/dao"
	dm "github.com/mt1976/purse/datamodel"
	logs "github.com/mt1976/purse/logs"
	cron "github.com/robfig/cron/v3"
)

func StaticRates_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "METRICS_STAT"
	j.Name = "METRICS_STAT"
	j.Period = "*/11 7-19 * * 1-5"
	j.Description = "Update FX Spot rate from barchart.com"
	j.Type = core.Aquirer
	return j
}

func StaticRates_Register(c *cron.Cron) {
	application.Schedule_Register(StaticRates_Job())
	c.AddFunc(StaticRates_Job().Period, func() { StaticRates_Run() })
}

// RunJobRollover is a Rollover function
func StaticRates_Run() {
	//logs.StartJob(RatesFXSpot_Job().Name)

	logs.StartJob(StaticRates_Job().Name)
	var message string
	/// CONTENT STARTS

	noItems, symbolList, err := dao.Symbol_GetListByType("STAT")
	if err != nil {
		log.Println(err.Error())
	}

	baseURI := ""

	var rateCard fxRateCard

	for i := 0; i < noItems; i++ {
		//fmt.Printf("%s\n", symbolList[i])
		symbolData := symbolList[i]
		rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate(symbolData, baseURI))
	}

	if err != nil {
		message = err.Error()
	}

	/// CONTENT ENDS
	application.Schedule_Update(StaticRates_Job(), message)
	logs.EndJob(StaticRates_Job().Name)
}

func getStaticRatesRate(inSymbol dm.Symbol, baseURI string) fxRate {

	var thisRate fxRate
	thisRate.bidRate, _ = strconv.ParseFloat(inSymbol.StaticValue, 64)
	thisRate.askRate, _ = strconv.ParseFloat(inSymbol.StaticValue, 64)
	fmt.Printf("thisRate: %v\n", thisRate)

	if inSymbol.Factor != "" {
		factor, _ := strconv.ParseFloat(inSymbol.Factor, 64)
		thisRate.bidRate = factor / thisRate.bidRate
		thisRate.askRate = factor / thisRate.askRate
	}

	fmt.Printf("thisRate: %v\n", thisRate)

	var symb dm.SymbolMetrics
	symb.Symbol = inSymbol.Code
	symb.URI = ""

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
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("AUDUSD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("EURGBP"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("EURJPY"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("EURUSD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("GBPUSD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("NZDUSD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("USDCAD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("USDCHF"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("USDHKD"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("USDJPY"))
// 	rateCard.fxRates = append(rateCard.fxRates, getStaticRatesRate("USDSGD"))
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
