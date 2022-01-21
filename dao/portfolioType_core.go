package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/portfoliotype.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : PortfolioType (portfoliotype)
// Endpoint 	        : PortfolioType (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:37
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

// PortfolioType_GetList() returns a list of all PortfolioType records
func PortfolioType_GetList() (int, []dm.PortfolioType, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.PortfolioType_SQLTable)
	count, portfoliotypeList, _, _ := portfoliotype_Fetch(tsql)
	
	return count, portfoliotypeList, nil
}



// PortfolioType_GetByID() returns a single PortfolioType record
func PortfolioType_GetByID(id string) (int, dm.PortfolioType, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.PortfolioType_SQLTable)
	tsql = tsql + " WHERE " + dm.PortfolioType_SQLSearchID + "='" + id + "'"
	_, _, portfoliotypeItem, _ := portfoliotype_Fetch(tsql)

	return 1, portfoliotypeItem, nil
}



// PortfolioType_DeleteByID() deletes a single PortfolioType record
func PortfolioType_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.PortfolioType_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.PortfolioType_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// PortfolioType_Store() saves/stores a PortfolioType record to the database
func PortfolioType_Store(r dm.PortfolioType,req *http.Request) error {

	err := portfoliotype_Save(r,Audit_User(req))

	return err
}

// PortfolioType_StoreSystem() saves/stores a PortfolioType record to the database
func PortfolioType_StoreSystem(r dm.PortfolioType) error {
	
	err := portfoliotype_Save(r,Audit_Host())

	return err
}

// portfoliotype_Save() saves/stores a PortfolioType record to the database
func portfoliotype_Save(r dm.PortfolioType,usr string) error {

    var err error

// If there are fields below, create the methods in dao\PortfolioType_Impl.go





















	logs.Storing("PortfolioType",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = PortfolioType_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.PortfolioType_SYSId, r.SYSId)
ts = addData(ts, dm.PortfolioType_Code, r.Code)
ts = addData(ts, dm.PortfolioType_Name, r.Name)
ts = addData(ts, dm.PortfolioType_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.PortfolioType_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.PortfolioType_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.PortfolioType_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.PortfolioType_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.PortfolioType_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.PortfolioType_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	PortfolioType_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// portfoliotype_Fetch read all employees
func portfoliotype_Fetch(tsql string) (int, []dm.PortfolioType, dm.PortfolioType, error) {

	var recItem dm.PortfolioType
	var recList []dm.PortfolioType

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.PortfolioType_SYSId, "0")
   recItem.Code  = get_String(rec, dm.PortfolioType_Code, "")
   recItem.Name  = get_String(rec, dm.PortfolioType_Name, "")
   recItem.SYSCreated  = get_String(rec, dm.PortfolioType_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.PortfolioType_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.PortfolioType_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.PortfolioType_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.PortfolioType_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.PortfolioType_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\PortfolioType_Impl.go



















	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func PortfolioType_NewID(r dm.PortfolioType) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

