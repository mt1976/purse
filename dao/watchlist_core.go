package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/watchlist.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Watchlist (watchlist)
// Endpoint 	        : Watchlist (Id)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:39
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

// Watchlist_GetList() returns a list of all Watchlist records
func Watchlist_GetList() (int, []dm.Watchlist, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Watchlist_SQLTable)
	count, watchlistList, _, _ := watchlist_Fetch(tsql)
	
	return count, watchlistList, nil
}



// Watchlist_GetByID() returns a single Watchlist record
func Watchlist_GetByID(id string) (int, dm.Watchlist, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Watchlist_SQLTable)
	tsql = tsql + " WHERE " + dm.Watchlist_SQLSearchID + "='" + id + "'"
	_, _, watchlistItem, _ := watchlist_Fetch(tsql)

	return 1, watchlistItem, nil
}



// Watchlist_DeleteByID() deletes a single Watchlist record
func Watchlist_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Watchlist_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Watchlist_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Watchlist_Store() saves/stores a Watchlist record to the database
func Watchlist_Store(r dm.Watchlist,req *http.Request) error {

	err := watchlist_Save(r,Audit_User(req))

	return err
}

// Watchlist_StoreSystem() saves/stores a Watchlist record to the database
func Watchlist_StoreSystem(r dm.Watchlist) error {
	
	err := watchlist_Save(r,Audit_Host())

	return err
}

// watchlist_Save() saves/stores a Watchlist record to the database
func watchlist_Save(r dm.Watchlist,usr string) error {

    var err error

// If there are fields below, create the methods in dao\Watchlist_Impl.go



























	logs.Storing("Watchlist",fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = Watchlist_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Watchlist_SYSId, r.SYSId)
ts = addData(ts, dm.Watchlist_Id, r.Id)
ts = addData(ts, dm.Watchlist_Code, r.Code)
ts = addData(ts, dm.Watchlist_Name, r.Name)
ts = addData(ts, dm.Watchlist_Symbol, r.Symbol)
ts = addData(ts, dm.Watchlist_Notes, r.Notes)
ts = addData(ts, dm.Watchlist_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Watchlist_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Watchlist_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Watchlist_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Watchlist_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Watchlist_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Watchlist_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Watchlist_Delete(r.Id)
	das.Execute(tsql)



	return err

}


// watchlist_Fetch read all employees
func watchlist_Fetch(tsql string) (int, []dm.Watchlist, dm.Watchlist, error) {

	var recItem dm.Watchlist
	var recList []dm.Watchlist

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Watchlist_SYSId, "0")
   recItem.Id  = get_String(rec, dm.Watchlist_Id, "")
   recItem.Code  = get_String(rec, dm.Watchlist_Code, "")
   recItem.Name  = get_String(rec, dm.Watchlist_Name, "")
   recItem.Symbol  = get_String(rec, dm.Watchlist_Symbol, "")
   recItem.Notes  = get_String(rec, dm.Watchlist_Notes, "")
   recItem.SYSCreated  = get_String(rec, dm.Watchlist_SYSCreated, "")
   recItem.SYSUpdated  = get_String(rec, dm.Watchlist_SYSUpdated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Watchlist_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Watchlist_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Watchlist_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Watchlist_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\Watchlist_Impl.go

























	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Watchlist_NewID(r dm.Watchlist) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

