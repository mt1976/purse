package application

// ----------------------------------------------------------------
// Automatically generated  "/application/portfolio.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Portfolio (portfolio)
// Endpoint 	        : Portfolio (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 19/01/2022 at 16:06:49
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	core "github.com/mt1976/purse/core"
	dao "github.com/mt1976/purse/dao"
	dm "github.com/mt1976/purse/datamodel"
	logs "github.com/mt1976/purse/logs"
)

//Portfolio_Publish annouces the endpoints available for this object
func Portfolio_PublishImpl(mux http.ServeMux) {

	mux.HandleFunc("/PortfolioReval/", Portfolio_HandlerReval)
	logs.Publish("Application", dm.Portfolio_Title+"Reval")

}

//Portfolio_HandlerView is the handler used to View a page
func Portfolio_HandlerReval(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Portfolio_QueryString)
	_, rD, _ := dao.Portfolio_GetByID(searchID)

	portfolio_Revalue(rD)

	http.Redirect(w, r, Portfolio_Redirect, http.StatusFound)

}

func portfolio_Revalue(portfolio dm.Portfolio) error {

	// Loop through the data and reval each holding

	rCount, rList, _ := dao.Holding_GetListByPortfolio(portfolio.Code)

	for i := 0; i < rCount; i++ {
		localValue, err := holding_Revalue(rList[i])
		if err != nil {
			return err
		}
		fmt.Println(localValue)
	}
	return nil
}

func Portgolio_Revalue(portID string) {

	_, rD, _ := dao.Portfolio_GetByID(portID)

	err := portfolio_Revalue(rD)
	if err != nil {
		fmt.Println(err)
	}

}
