package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/datasource.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DataSource (datasource)
// Endpoint 	        : DataSource (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:36
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"log"
	
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/purse/core"
	das  "github.com/mt1976/purse/das"
	
	
	dm   "github.com/mt1976/purse/datamodel"
	logs   "github.com/mt1976/purse/logs"
)

// DataSource_GetList() returns a list of all DataSource records
func DataSource_GetList() (int, []dm.DataSource, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataSource_SQLTable)
	count, datasourceList, _, _ := datasource_Fetch(tsql)
	
	return count, datasourceList, nil
}



// DataSource_GetByID() returns a single DataSource record
func DataSource_GetByID(id string) (int, dm.DataSource, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataSource_SQLTable)
	tsql = tsql + " WHERE " + dm.DataSource_SQLSearchID + "='" + id + "'"
	_, _, datasourceItem, _ := datasource_Fetch(tsql)

	return 1, datasourceItem, nil
}



// DataSource_DeleteByID() deletes a single DataSource record
func DataSource_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.DataSource_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.DataSource_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// DataSource_Store() saves/stores a DataSource record to the database
func DataSource_Store(r dm.DataSource,req *http.Request) error {

	err := datasource_Save(r,Audit_User(req))

	return err
}

// DataSource_StoreSystem() saves/stores a DataSource record to the database
func DataSource_StoreSystem(r dm.DataSource) error {
	
	err := datasource_Save(r,Audit_Host())

	return err
}

// datasource_Save() saves/stores a DataSource record to the database
func datasource_Save(r dm.DataSource,usr string) error {

    var err error

// If there are fields below, create the methods in dao\DataSource_Impl.go





























	logs.Storing("DataSource",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = DataSource_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.DataSource_SYSId, r.SYSId)
ts = addData(ts, dm.DataSource_Code, r.Code)
ts = addData(ts, dm.DataSource_Name, r.Name)
ts = addData(ts, dm.DataSource_URI, r.URI)
ts = addData(ts, dm.DataSource_Description, r.Description)
ts = addData(ts, dm.DataSource_Notes, r.Notes)
ts = addData(ts, dm.DataSource_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.DataSource_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.DataSource_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.DataSource_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.DataSource_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.DataSource_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.DataSource_AuthKey, r.AuthKey)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataSource_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	DataSource_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// datasource_Fetch read all employees
func datasource_Fetch(tsql string) (int, []dm.DataSource, dm.DataSource, error) {

	var recItem dm.DataSource
	var recList []dm.DataSource

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.DataSource_SYSId, "0")
   recItem.Code  = get_String(rec, dm.DataSource_Code, "")
   recItem.Name  = get_String(rec, dm.DataSource_Name, "")
   recItem.URI  = get_String(rec, dm.DataSource_URI, "")
   recItem.Description  = get_String(rec, dm.DataSource_Description, "")
   recItem.Notes  = get_String(rec, dm.DataSource_Notes, "")
   recItem.SYSCreated  = get_String(rec, dm.DataSource_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.DataSource_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.DataSource_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.DataSource_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.DataSource_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.DataSource_SYSUpdatedHost, "")
   recItem.AuthKey  = get_String(rec, dm.DataSource_AuthKey, "")
// If there are fields below, create the methods in dao\DataSource_Impl.go



























	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func DataSource_NewID(r dm.DataSource) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

