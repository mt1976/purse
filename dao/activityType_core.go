package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/activitytype.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ActivityType (activitytype)
// Endpoint 	        : ActivityType (Code)
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

// ActivityType_GetList() returns a list of all ActivityType records
func ActivityType_GetList() (int, []dm.ActivityType, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.ActivityType_SQLTable)
	count, activitytypeList, _, _ := activitytype_Fetch(tsql)
	
	return count, activitytypeList, nil
}



// ActivityType_GetByID() returns a single ActivityType record
func ActivityType_GetByID(id string) (int, dm.ActivityType, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.ActivityType_SQLTable)
	tsql = tsql + " WHERE " + dm.ActivityType_SQLSearchID + "='" + id + "'"
	_, _, activitytypeItem, _ := activitytype_Fetch(tsql)

	return 1, activitytypeItem, nil
}



// ActivityType_DeleteByID() deletes a single ActivityType record
func ActivityType_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.ActivityType_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.ActivityType_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// ActivityType_Store() saves/stores a ActivityType record to the database
func ActivityType_Store(r dm.ActivityType,req *http.Request) error {

	err := activitytype_Save(r,Audit_User(req))

	return err
}

// ActivityType_StoreSystem() saves/stores a ActivityType record to the database
func ActivityType_StoreSystem(r dm.ActivityType) error {
	
	err := activitytype_Save(r,Audit_Host())

	return err
}

// activitytype_Save() saves/stores a ActivityType record to the database
func activitytype_Save(r dm.ActivityType,usr string) error {

    var err error

// If there are fields below, create the methods in dao\ActivityType_Impl.go























	logs.Storing("ActivityType",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = ActivityType_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.ActivityType_SYSId, r.SYSId)
ts = addData(ts, dm.ActivityType_Code, r.Code)
ts = addData(ts, dm.ActivityType_Name, r.Name)
ts = addData(ts, dm.ActivityType_Factor, r.Factor)
ts = addData(ts, dm.ActivityType_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.ActivityType_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.ActivityType_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.ActivityType_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.ActivityType_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.ActivityType_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.ActivityType_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	ActivityType_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// activitytype_Fetch read all employees
func activitytype_Fetch(tsql string) (int, []dm.ActivityType, dm.ActivityType, error) {

	var recItem dm.ActivityType
	var recList []dm.ActivityType

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.ActivityType_SYSId, "0")
   recItem.Code  = get_String(rec, dm.ActivityType_Code, "")
   recItem.Name  = get_String(rec, dm.ActivityType_Name, "")
   recItem.Factor  = get_String(rec, dm.ActivityType_Factor, "")
   recItem.SYSCreated  = get_String(rec, dm.ActivityType_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.ActivityType_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.ActivityType_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.ActivityType_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.ActivityType_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.ActivityType_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\ActivityType_Impl.go





















	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func ActivityType_NewID(r dm.ActivityType) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

