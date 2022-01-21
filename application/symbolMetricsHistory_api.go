package application

// ----------------------------------------------------------------
// Automatically generated  "/application/symbolmetricshistory_api.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SymbolMetricsHistory (symbolmetricshistory)
// Endpoint 	        : SymbolMetricsHistory (TSKey)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:38
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"encoding/json"

	"net/http"


	core    "github.com/mt1976/purse/core"
	dao     "github.com/mt1976/purse/dao"
	dm      "github.com/mt1976/purse/datamodel"
)

//SymbolMetricsHistory_Handler is the handler for the api calls
func SymbolMetricsHistory_Handler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	//TODO: Add your security validation here
	//        new => POST
	//  	 read => GET
	// 	   update => PUT
	//     delete => DELETE

	httpMethod := r.Method
	
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	//responseStatus := http.StatusOK

	switch httpMethod {
	case http.MethodGet:
		symbolmetricshistory_MethodGet(w, r)

	case http.MethodPost:
		symbolmetricshistory_MethodPost(w, r)

	case http.MethodPut:
		symbolmetricshistory_MethodPost(w, r)
	case http.MethodDelete:

		symbolmetricshistory_MethodDelete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// Handles GET requests for SymbolMetricsHistory
func symbolmetricshistory_MethodGet(w http.ResponseWriter, r *http.Request) {
	//logs.Information("PATH", r.URL.Path)
	searchID := core.GetURLparam(r, dm.SymbolMetricsHistory_QueryString)
	//logs.Information("GET", searchID)
	w.Header().Set("Content-Type", "application/json")

	if searchID == "" {
		//Get All Entities
		//logs.Information("GET", "All")
		noRecs, records, _ := dao.SymbolMetricsHistory_GetList()
		var ci core.ContentList
		ci.Count = noRecs
		ci.Key = dm.SymbolMetricsHistory_QueryString
		for _, v := range records {
			ciContent := core.ContentListItem{ID:v.TSKey,Query:"?" + ci.Key +"="+ v.TSKey}
			ci.Items= append(ci.Items, ciContent)
		}
		json_data, _ := json.Marshal(ci)
		w.Write(json_data)

		if noRecs == 0 {
			w.WriteHeader(int(http.StatusNotFound))
		} else {
			w.WriteHeader(int(http.StatusOK))
		}


	} else {
		//Get a specific entity
		_, record, _ := dao.SymbolMetricsHistory_GetByID(searchID)
		//spew.Dump(record)
		json_data, _ := json.Marshal(record)
		w.Write(json_data)

		if record.TSKey == "" {
		    w.WriteHeader(int(http.StatusNotFound))
		} else {
			w.WriteHeader(int(http.StatusOK))
		}
	}


}

//Handles POST & PUT requests for SymbolMetricsHistory
func symbolmetricshistory_MethodPost(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("POST")
	//fmt.Printf("r.Body: %v\n", r.Body)

	decoder := json.NewDecoder(r.Body)
	var t dm.SymbolMetricsHistory
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(int(http.StatusNotFound))
		panic(err)
	} else {
		w.WriteHeader(int(http.StatusOK))
	}
	//spew.Dump(t)
	err = dao.SymbolMetricsHistory_StoreSystem(t)
	//logs.Processing("POST BACK")
	//logs.Information("POST", err.Error())
	if err != nil {
		//	panic(err)
		w.WriteHeader(int(http.StatusNotFound))
	} else {
		w.WriteHeader(int(http.StatusOK))
	}
	//logs.Success("POST")
}
//Handles DELETE requests for SymbolMetricsHistory
func symbolmetricshistory_MethodDelete(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("DELETE")
	deleteID := core.GetURLparam(r, dm.SymbolMetricsHistory_QueryString)
	//logs.Information("DELETE", deleteID)

	dao.SymbolMetricsHistory_Delete(deleteID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	//fmt.Printf("json_data: %v\n", json_data)

	//logs.Success("DELETE")
}
