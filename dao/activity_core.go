package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/activity.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Activity (activity)
// Endpoint 	        : Activity (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:35
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

// Activity_GetList() returns a list of all Activity records
func Activity_GetList() (int, []dm.Activity, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Activity_SQLTable)
	count, activityList, _, _ := activity_Fetch(tsql)
	
	return count, activityList, nil
}



// Activity_GetByID() returns a single Activity record
func Activity_GetByID(id string) (int, dm.Activity, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Activity_SQLTable)
	tsql = tsql + " WHERE " + dm.Activity_SQLSearchID + "='" + id + "'"
	_, _, activityItem, _ := activity_Fetch(tsql)

	return 1, activityItem, nil
}



// Activity_DeleteByID() deletes a single Activity record
func Activity_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Activity_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Activity_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Activity_Store() saves/stores a Activity record to the database
func Activity_Store(r dm.Activity,req *http.Request) error {

	err := activity_Save(r,Audit_User(req))

	return err
}

// Activity_StoreSystem() saves/stores a Activity record to the database
func Activity_StoreSystem(r dm.Activity) error {
	
	err := activity_Save(r,Audit_Host())

	return err
}

// activity_Save() saves/stores a Activity record to the database
func activity_Save(r dm.Activity,usr string) error {

    var err error

// If there are fields below, create the methods in dao\Activity_Impl.go







































	logs.Storing("Activity",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Activity_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Activity_SYSId, r.SYSId)
ts = addData(ts, dm.Activity_Code, r.Code)
ts = addData(ts, dm.Activity_Symbol, r.Symbol)
ts = addData(ts, dm.Activity_Date, r.Date)
ts = addData(ts, dm.Activity_Type, r.Type)
ts = addData(ts, dm.Activity_Portfolio, r.Portfolio)
ts = addData(ts, dm.Activity_Amount, r.Amount)
ts = addData(ts, dm.Activity_Price, r.Price)
ts = addData(ts, dm.Activity_Notes, r.Notes)
ts = addData(ts, dm.Activity_Holding, r.Holding)
ts = addData(ts, dm.Activity_ValueAMT, r.ValueAMT)
ts = addData(ts, dm.Activity_ValueCCY, r.ValueCCY)
ts = addData(ts, dm.Activity_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Activity_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Activity_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Activity_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Activity_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Activity_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Activity_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Activity_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// activity_Fetch read all employees
func activity_Fetch(tsql string) (int, []dm.Activity, dm.Activity, error) {

	var recItem dm.Activity
	var recList []dm.Activity

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Activity_SYSId, "0")
   recItem.Code  = get_String(rec, dm.Activity_Code, "")
   recItem.Symbol  = get_String(rec, dm.Activity_Symbol, "")
   recItem.Date  = get_String(rec, dm.Activity_Date, "")
   recItem.Type  = get_String(rec, dm.Activity_Type, "")
   recItem.Portfolio  = get_String(rec, dm.Activity_Portfolio, "")
   recItem.Amount  = get_String(rec, dm.Activity_Amount, "")
   recItem.Price  = get_String(rec, dm.Activity_Price, "")
   recItem.Notes  = get_String(rec, dm.Activity_Notes, "")
   recItem.Holding  = get_String(rec, dm.Activity_Holding, "")
   recItem.ValueAMT  = get_String(rec, dm.Activity_ValueAMT, "")
   recItem.ValueCCY  = get_String(rec, dm.Activity_ValueCCY, "")
   recItem.SYSCreated  = get_String(rec, dm.Activity_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Activity_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Activity_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Activity_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Activity_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Activity_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\Activity_Impl.go





































	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Activity_NewID(r dm.Activity) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

