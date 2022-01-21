package dao

import (
	core "github.com/mt1976/purse/core"
	dm "github.com/mt1976/purse/datamodel"
	"github.com/mt1976/purse/logs"
)

func holding_Type_OverrideStore(r dm.Holding) string {
	logs.Information("Holding_Type_OverrideStore", r.Type_Lookup)
	return r.Type_Lookup
}

func holding_Portfolio_OverrideStore(r dm.Holding) string {
	logs.Information("Holding_Portfolio_OverrideStore", r.Portfolio_Lookup)
	return r.Portfolio_Lookup
}

//r.Type  = holding_Type_OverrideStore (r)
//r.Portfolio  = holding_Portfolio_OverrideStore (r)

func holding_Type_OverrideFetch(r dm.Holding) string {
	return r.Type
}

func holding_Portfolio_OverrideFetch(r dm.Holding) string {
	return r.Portfolio
}

//recItem.Type  = holding_Type_OverrideFetch (recItem)
//recItem.Portfolio  = holding_Portfolio_OverrideFetch (recItem)

func holding_Instrument_OverrideStore(r dm.Holding) string {
	return r.Instrument_Lookup
}
func holding_Instrument_OverrideFetch(r dm.Holding) string {
	return r.Instrument
}

// Holding_GetList() returns a list of all Holding records
func Holding_GetListByPortfolio(portID string) (int, []dm.Holding, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Holding_SQLTable)
	tsql = tsql + " WHERE " + dm.Holding_Portfolio + "='" + portID + "'"
	count, holdingList, _, _ := holding_Fetch(tsql)

	return count, holdingList, nil
}
