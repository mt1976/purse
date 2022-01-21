package dao

import (
	core "github.com/mt1976/purse/core"
	dm "github.com/mt1976/purse/datamodel"
)

// DataSource_GetList() returns a list of all DataSource records
func Symbol_GetListBySource(id string) (int, []dm.Symbol, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Symbol_SQLTable)
	tsql = tsql + " WHERE " + dm.Symbol_DataSource + "='" + id + "'"
	count, datasourceList, _, _ := symbol_Fetch(tsql)

	return count, datasourceList, nil
}

// DataSource_GetList() returns a list of all DataSource records
func Symbol_GetListByType(id string) (int, []dm.Symbol, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Symbol_SQLTable)
	tsql = tsql + " WHERE " + dm.Symbol_Type + "='" + id + "'"
	count, datasourceList, _, _ := symbol_Fetch(tsql)

	return count, datasourceList, nil
}
