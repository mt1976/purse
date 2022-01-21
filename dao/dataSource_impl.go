package dao

import (
	core "github.com/mt1976/purse/core"
	dm "github.com/mt1976/purse/datamodel"
)

// DataSource_GetList() returns a list of all DataSource records
func DataSource_GetListByID(id string) (int, []dm.DataSource, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataSource_SQLTable)
	tsql = tsql + " WHERE " + dm.DataSource_SQLSearchID + "='" + id + "'"
	count, datasourceList, _, _ := datasource_Fetch(tsql)

	return count, datasourceList, nil
}
