package dao

import (
	"time"

	"github.com/mt1976/purse/core"
	dm "github.com/mt1976/purse/datamodel"
)

func SymbolMetricsHistory_StoreRefresh(inMetric dm.SymbolMetrics) error {

	currentTime := time.Now()

	var metHistory dm.SymbolMetricsHistory
	metHistory.Symbol = inMetric.Symbol
	metHistory.Date = currentTime.Format(core.DATEFORMATYMD)
	metHistory.Price = inMetric.Price
	metHistory.RawPrice = inMetric.RawPrice
	metHistory.Bid = inMetric.Bid
	metHistory.RawBid = inMetric.RawBid
	metHistory.Offer = inMetric.Offer
	metHistory.RawOffer = inMetric.RawOffer

	metHistory.URI = inMetric.URI
	metHistory.Type = "Refresh"

	metHistory.TSKey = metHistory.Symbol + core.IDSep + metHistory.Type + core.IDSep + currentTime.Format(core.DATETIMEUNQ)

	err := SymbolMetricsHistory_StoreSystem(metHistory)

	return err
}

func SymbolMetricsHistory_StoreSnapshot(inMetric dm.SymbolMetrics) error {

	currentTime := time.Now()

	var metHistory dm.SymbolMetricsHistory
	metHistory.Symbol = inMetric.Symbol
	metHistory.Date = currentTime.Format(core.DATEFORMATYMD)
	metHistory.Price = inMetric.Price
	metHistory.RawPrice = inMetric.RawPrice
	metHistory.Bid = inMetric.Bid
	metHistory.RawBid = inMetric.RawBid
	metHistory.Offer = inMetric.Offer
	metHistory.RawOffer = inMetric.RawOffer
	metHistory.URI = inMetric.URI
	metHistory.Type = "Snapshot"

	metHistory.TSKey = metHistory.Symbol + core.IDSep + metHistory.Type + core.IDSep + currentTime.Format(core.DATEFORMATYMD)

	err := SymbolMetricsHistory_StoreSystem(metHistory)

	return err
}
