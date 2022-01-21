package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/holdingtype.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : HoldingType (holdingtype)
// Endpoint 	        : HoldingType (Code)
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

// HoldingType_GetList() returns a list of all HoldingType records
func HoldingType_GetList() (int, []dm.HoldingType, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.HoldingType_SQLTable)
	count, holdingtypeList, _, _ := holdingtype_Fetch(tsql)
	
	return count, holdingtypeList, nil
}



// HoldingType_GetByID() returns a single HoldingType record
func HoldingType_GetByID(id string) (int, dm.HoldingType, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.HoldingType_SQLTable)
	tsql = tsql + " WHERE " + dm.HoldingType_SQLSearchID + "='" + id + "'"
	_, _, holdingtypeItem, _ := holdingtype_Fetch(tsql)

	return 1, holdingtypeItem, nil
}



// HoldingType_DeleteByID() deletes a single HoldingType record
func HoldingType_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.HoldingType_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.HoldingType_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// HoldingType_Store() saves/stores a HoldingType record to the database
func HoldingType_Store(r dm.HoldingType,req *http.Request) error {

	err := holdingtype_Save(r,Audit_User(req))

	return err
}

// HoldingType_StoreSystem() saves/stores a HoldingType record to the database
func HoldingType_StoreSystem(r dm.HoldingType) error {
	
	err := holdingtype_Save(r,Audit_Host())

	return err
}

// holdingtype_Save() saves/stores a HoldingType record to the database
func holdingtype_Save(r dm.HoldingType,usr string) error {

    var err error

// If there are fields below, create the methods in dao\HoldingType_Impl.go























	logs.Storing("HoldingType",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = HoldingType_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.HoldingType_SYSId, r.SYSId)
ts = addData(ts, dm.HoldingType_Code, r.Code)
ts = addData(ts, dm.HoldingType_Name, r.Name)
ts = addData(ts, dm.HoldingType_Factor, r.Factor)
ts = addData(ts, dm.HoldingType_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.HoldingType_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.HoldingType_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.HoldingType_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.HoldingType_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.HoldingType_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.HoldingType_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	HoldingType_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// holdingtype_Fetch read all employees
func holdingtype_Fetch(tsql string) (int, []dm.HoldingType, dm.HoldingType, error) {

	var recItem dm.HoldingType
	var recList []dm.HoldingType

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.HoldingType_SYSId, "0")
   recItem.Code  = get_String(rec, dm.HoldingType_Code, "")
   recItem.Name  = get_String(rec, dm.HoldingType_Name, "")
   recItem.Factor  = get_String(rec, dm.HoldingType_Factor, "")
   recItem.SYSCreated  = get_String(rec, dm.HoldingType_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.HoldingType_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.HoldingType_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.HoldingType_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.HoldingType_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.HoldingType_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\HoldingType_Impl.go





















	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func HoldingType_NewID(r dm.HoldingType) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

