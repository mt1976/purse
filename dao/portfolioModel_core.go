package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/portfoliomodel.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : PortfolioModel (portfoliomodel)
// Endpoint 	        : PortfolioModel (Code)
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

// PortfolioModel_GetList() returns a list of all PortfolioModel records
func PortfolioModel_GetList() (int, []dm.PortfolioModel, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.PortfolioModel_SQLTable)
	count, portfoliomodelList, _, _ := portfoliomodel_Fetch(tsql)
	
	return count, portfoliomodelList, nil
}



// PortfolioModel_GetByID() returns a single PortfolioModel record
func PortfolioModel_GetByID(id string) (int, dm.PortfolioModel, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.PortfolioModel_SQLTable)
	tsql = tsql + " WHERE " + dm.PortfolioModel_SQLSearchID + "='" + id + "'"
	_, _, portfoliomodelItem, _ := portfoliomodel_Fetch(tsql)

	return 1, portfoliomodelItem, nil
}



// PortfolioModel_DeleteByID() deletes a single PortfolioModel record
func PortfolioModel_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.PortfolioModel_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.PortfolioModel_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// PortfolioModel_Store() saves/stores a PortfolioModel record to the database
func PortfolioModel_Store(r dm.PortfolioModel,req *http.Request) error {

	err := portfoliomodel_Save(r,Audit_User(req))

	return err
}

// PortfolioModel_StoreSystem() saves/stores a PortfolioModel record to the database
func PortfolioModel_StoreSystem(r dm.PortfolioModel) error {
	
	err := portfoliomodel_Save(r,Audit_Host())

	return err
}

// portfoliomodel_Save() saves/stores a PortfolioModel record to the database
func portfoliomodel_Save(r dm.PortfolioModel,usr string) error {

    var err error

// If there are fields below, create the methods in dao\PortfolioModel_Impl.go





















	logs.Storing("PortfolioModel",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = PortfolioModel_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.PortfolioModel_SYSId, r.SYSId)
ts = addData(ts, dm.PortfolioModel_Code, r.Code)
ts = addData(ts, dm.PortfolioModel_Name, r.Name)
ts = addData(ts, dm.PortfolioModel_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.PortfolioModel_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.PortfolioModel_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.PortfolioModel_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.PortfolioModel_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.PortfolioModel_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.PortfolioModel_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	PortfolioModel_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// portfoliomodel_Fetch read all employees
func portfoliomodel_Fetch(tsql string) (int, []dm.PortfolioModel, dm.PortfolioModel, error) {

	var recItem dm.PortfolioModel
	var recList []dm.PortfolioModel

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.PortfolioModel_SYSId, "0")
   recItem.Code  = get_String(rec, dm.PortfolioModel_Code, "")
   recItem.Name  = get_String(rec, dm.PortfolioModel_Name, "")
   recItem.SYSCreated  = get_String(rec, dm.PortfolioModel_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.PortfolioModel_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.PortfolioModel_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.PortfolioModel_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.PortfolioModel_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.PortfolioModel_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\PortfolioModel_Impl.go



















	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func PortfolioModel_NewID(r dm.PortfolioModel) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

