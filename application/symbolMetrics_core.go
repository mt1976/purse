package application
// ----------------------------------------------------------------
// Automatically generated  "/application/symbolmetrics.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SymbolMetrics (symbolmetrics)
// Endpoint 	        : SymbolMetrics (Symbol)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:38
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/purse/core"
	dao     "github.com/mt1976/purse/dao"
	dm      "github.com/mt1976/purse/datamodel"
	logs    "github.com/mt1976/purse/logs"
)

//symbolmetrics_PageList provides the information for the template for a list of SymbolMetricss
type SymbolMetrics_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.SymbolMetrics
}

//symbolmetrics_Page provides the information for the template for an individual SymbolMetrics
type SymbolMetrics_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		SYSId string
		Symbol string
		Price string
		URI string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
		Bid string
		Offer string
		RawPrice string
		RawBid string
		RawOffer string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	SymbolMetrics_Redirect = dm.SymbolMetrics_PathList
)

//SymbolMetrics_Publish annouces the endpoints available for this object
func SymbolMetrics_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.SymbolMetrics_Path, SymbolMetrics_Handler)
	mux.HandleFunc(dm.SymbolMetrics_PathList, SymbolMetrics_HandlerList)
	mux.HandleFunc(dm.SymbolMetrics_PathView, SymbolMetrics_HandlerView)
	mux.HandleFunc(dm.SymbolMetrics_PathEdit, SymbolMetrics_HandlerEdit)
	mux.HandleFunc(dm.SymbolMetrics_PathNew, SymbolMetrics_HandlerNew)
	mux.HandleFunc(dm.SymbolMetrics_PathSave, SymbolMetrics_HandlerSave)
	mux.HandleFunc(dm.SymbolMetrics_PathDelete, SymbolMetrics_HandlerDelete)
	logs.Publish("Application", dm.SymbolMetrics_Title)
    core.Catalog_Add(dm.SymbolMetrics_Title, dm.SymbolMetrics_Path, "", dm.SymbolMetrics_QueryString, "APP")
}

//SymbolMetrics_HandlerList is the handler for the list page
func SymbolMetrics_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.SymbolMetrics
	noItems, returnList, _ := dao.SymbolMetrics_GetList()

	pageDetail := SymbolMetrics_PageList{
		Title:            CardTitle(dm.SymbolMetrics_Title, core.Action_List),
		PageTitle:        PageTitle(dm.SymbolMetrics_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.SymbolMetrics_TemplateList, w, r, pageDetail)

}

//SymbolMetrics_HandlerView is the handler used to View a page
func SymbolMetrics_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SymbolMetrics_QueryString)
	_, rD, _ := dao.SymbolMetrics_GetByID(searchID)

	pageDetail := SymbolMetrics_Page{
		Title:       CardTitle(dm.SymbolMetrics_Title, core.Action_View),
		PageTitle:   PageTitle(dm.SymbolMetrics_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Symbol = rD.Symbol
pageDetail.Price = rD.Price
pageDetail.URI = rD.URI
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.Bid = rD.Bid
pageDetail.Offer = rD.Offer
pageDetail.RawPrice = rD.RawPrice
pageDetail.RawBid = rD.RawBid
pageDetail.RawOffer = rD.RawOffer


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//


	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.SymbolMetrics_TemplateView, w, r, pageDetail)

}

//SymbolMetrics_HandlerEdit is the handler used generate the Edit page
func SymbolMetrics_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SymbolMetrics_QueryString)
	_, rD, _ := dao.SymbolMetrics_GetByID(searchID)
	
	pageDetail := SymbolMetrics_Page{
		Title:       CardTitle(dm.SymbolMetrics_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.SymbolMetrics_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Symbol = rD.Symbol
pageDetail.Price = rD.Price
pageDetail.URI = rD.URI
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.Bid = rD.Bid
pageDetail.Offer = rD.Offer
pageDetail.RawPrice = rD.RawPrice
pageDetail.RawBid = rD.RawBid
pageDetail.RawOffer = rD.RawOffer


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.SymbolMetrics_TemplateEdit, w, r, pageDetail)


}

//SymbolMetrics_HandlerSave is the handler used process the saving of an SymbolMetrics
func SymbolMetrics_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Symbol"))

	var item dm.SymbolMetrics
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.SymbolMetrics_SYSId)
		item.Symbol = r.FormValue(dm.SymbolMetrics_Symbol)
		item.Price = r.FormValue(dm.SymbolMetrics_Price)
		item.URI = r.FormValue(dm.SymbolMetrics_URI)
		item.SYSCreated = r.FormValue(dm.SymbolMetrics_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.SymbolMetrics_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.SymbolMetrics_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.SymbolMetrics_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.SymbolMetrics_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.SymbolMetrics_SYSUpdatedHost)
		item.Bid = r.FormValue(dm.SymbolMetrics_Bid)
		item.Offer = r.FormValue(dm.SymbolMetrics_Offer)
		item.RawPrice = r.FormValue(dm.SymbolMetrics_RawPrice)
		item.RawBid = r.FormValue(dm.SymbolMetrics_RawBid)
		item.RawOffer = r.FormValue(dm.SymbolMetrics_RawOffer)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.SymbolMetrics_Store(item,r)	

	http.Redirect(w, r, SymbolMetrics_Redirect, http.StatusFound)
}

//SymbolMetrics_HandlerNew is the handler used process the creation of an SymbolMetrics
func SymbolMetrics_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := SymbolMetrics_Page{
		Title:       CardTitle(dm.SymbolMetrics_Title, core.Action_New),
		PageTitle:   PageTitle(dm.SymbolMetrics_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Symbol = ""
pageDetail.Price = ""
pageDetail.URI = ""
pageDetail.SYSCreated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.Bid = ""
pageDetail.Offer = ""
pageDetail.RawPrice = ""
pageDetail.RawBid = ""
pageDetail.RawOffer = ""


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.SymbolMetrics_TemplateNew, w, r, pageDetail)

}

//SymbolMetrics_HandlerDelete is the handler used process the deletion of an SymbolMetrics
func SymbolMetrics_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.SymbolMetrics_QueryString)

	dao.SymbolMetrics_Delete(searchID)	

	http.Redirect(w, r, SymbolMetrics_Redirect, http.StatusFound)
}