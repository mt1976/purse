package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/portfolio.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Portfolio (portfolio)
// Endpoint 	        : Portfolio (Code)
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

// Portfolio_GetList() returns a list of all Portfolio records
func Portfolio_GetList() (int, []dm.Portfolio, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Portfolio_SQLTable)
	count, portfolioList, _, _ := portfolio_Fetch(tsql)
	
	return count, portfolioList, nil
}



// Portfolio_GetByID() returns a single Portfolio record
func Portfolio_GetByID(id string) (int, dm.Portfolio, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Portfolio_SQLTable)
	tsql = tsql + " WHERE " + dm.Portfolio_SQLSearchID + "='" + id + "'"
	_, _, portfolioItem, _ := portfolio_Fetch(tsql)

	return 1, portfolioItem, nil
}



// Portfolio_DeleteByID() deletes a single Portfolio record
func Portfolio_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Portfolio_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Portfolio_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Portfolio_Store() saves/stores a Portfolio record to the database
func Portfolio_Store(r dm.Portfolio,req *http.Request) error {

	err := portfolio_Save(r,Audit_User(req))

	return err
}

// Portfolio_StoreSystem() saves/stores a Portfolio record to the database
func Portfolio_StoreSystem(r dm.Portfolio) error {
	
	err := portfolio_Save(r,Audit_Host())

	return err
}

// portfolio_Save() saves/stores a Portfolio record to the database
func portfolio_Save(r dm.Portfolio,usr string) error {

    var err error

// If there are fields below, create the methods in dao\Portfolio_Impl.go






























  r.Type  = portfolio_Type_OverrideStore (r)
  r.DefaultModel  = portfolio_DefaultModel_OverrideStore (r)



	logs.Storing("Portfolio",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Portfolio_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Portfolio_SYSId, r.SYSId)
ts = addData(ts, dm.Portfolio_Code, r.Code)
ts = addData(ts, dm.Portfolio_Name, r.Name)
ts = addData(ts, dm.Portfolio_Type, r.Type)
ts = addData(ts, dm.Portfolio_DefaultModel, r.DefaultModel)
ts = addData(ts, dm.Portfolio_Description, r.Description)
ts = addData(ts, dm.Portfolio_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Portfolio_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Portfolio_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Portfolio_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Portfolio_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Portfolio_SYSUpdatedHost, r.SYSUpdatedHost)




	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Portfolio_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Portfolio_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// portfolio_Fetch read all employees
func portfolio_Fetch(tsql string) (int, []dm.Portfolio, dm.Portfolio, error) {

	var recItem dm.Portfolio
	var recList []dm.Portfolio

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Portfolio_SYSId, "0")
   recItem.Code  = get_String(rec, dm.Portfolio_Code, "")
   recItem.Name  = get_String(rec, dm.Portfolio_Name, "")
   recItem.Type  = get_String(rec, dm.Portfolio_Type, "")
   recItem.DefaultModel  = get_String(rec, dm.Portfolio_DefaultModel, "")
   recItem.Description  = get_String(rec, dm.Portfolio_Description, "")
   recItem.SYSCreated  = get_String(rec, dm.Portfolio_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Portfolio_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Portfolio_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Portfolio_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Portfolio_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Portfolio_SYSUpdatedHost, "")




// If there are fields below, create the methods in dao\Portfolio_Impl.go






























  recItem.Type  = portfolio_Type_OverrideFetch (recItem)
  recItem.DefaultModel  = portfolio_DefaultModel_OverrideFetch (recItem)

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Portfolio_NewID(r dm.Portfolio) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

