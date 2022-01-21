package dao

import (
	dm "github.com/mt1976/purse/datamodel"
	"github.com/mt1976/purse/logs"
)

func portfolio_Type_OverrideStore(r dm.Portfolio) string {
	logs.Information("portfolio_Type_OverrideStore", r.Type_Lookup)
	return r.Type_Lookup
}

func portfolio_DefaultModel_OverrideStore(r dm.Portfolio) string {
	logs.Information("portfolio_Portfolio_OverrideStore", r.DefaultModel_Lookup)
	return r.DefaultModel_Lookup
}

//r.Type  = portfolio_Type_OverrideStore (r)
//r.Portfolio  = portfolio_Portfolio_OverrideStore (r)

func portfolio_Type_OverrideFetch(r dm.Portfolio) string {
	return r.Type
}

func portfolio_DefaultModel_OverrideFetch(r dm.Portfolio) string {
	return r.DefaultModel
}

//recItem.Type  = portfolio_Type_OverrideFetch (recItem)
//recItem.Portfolio  = portfolio_Portfolio_OverrideFetch (recItem)
