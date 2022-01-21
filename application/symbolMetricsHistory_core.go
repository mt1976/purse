package application
// ----------------------------------------------------------------
// Automatically generated  "/application/symbolmetricshistory.go"
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
	
	"net/http"

	core    "github.com/mt1976/purse/core"
	dao     "github.com/mt1976/purse/dao"
	dm      "github.com/mt1976/purse/datamodel"
	logs    "github.com/mt1976/purse/logs"
)

//symbolmetricshistory_PageList provides the information for the template for a list of SymbolMetricsHistorys
type SymbolMetricsHistory_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.SymbolMetricsHistory
}

//symbolmetricshistory_Page provides the information for the template for an individual SymbolMetricsHistory
type SymbolMetricsHistory_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		SYSId string
		TSKey string
		Symbol string
		Date string
		Price string
		URI string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
		Type string
		Bid string
		Offer string
		RawPrice string
		RawBid string
		RawOffer string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	SymbolMetricsHistory_Redirect = dm.SymbolMetricsHistory_PathList
)

//SymbolMetricsHistory_Publish annouces the endpoints available for this object
func SymbolMetricsHistory_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.SymbolMetricsHistory_Path, SymbolMetricsHistory_Handler)
	mux.HandleFunc(dm.SymbolMetricsHistory_PathList, SymbolMetricsHistory_HandlerList)
	mux.HandleFunc(dm.SymbolMetricsHistory_PathView, SymbolMetricsHistory_HandlerView)
	mux.HandleFunc(dm.SymbolMetricsHistory_PathEdit, SymbolMetricsHistory_HandlerEdit)
	mux.HandleFunc(dm.SymbolMetricsHistory_PathNew, SymbolMetricsHistory_HandlerNew)
	mux.HandleFunc(dm.SymbolMetricsHistory_PathSave, SymbolMetricsHistory_HandlerSave)
	mux.HandleFunc(dm.SymbolMetricsHistory_PathDelete, SymbolMetricsHistory_HandlerDelete)
	logs.Publish("Application", dm.SymbolMetricsHistory_Title)
    core.Catalog_Add(dm.SymbolMetricsHistory_Title, dm.SymbolMetricsHistory_Path, "", dm.SymbolMetricsHistory_QueryString, "APP")
}

//SymbolMetricsHistory_HandlerList is the handler for the list page
func SymbolMetricsHistory_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.SymbolMetricsHistory
	noItems, returnList, _ := dao.SymbolMetricsHistory_GetList()

	pageDetail := SymbolMetricsHistory_PageList{
		Title:            CardTitle(dm.SymbolMetricsHistory_Title, core.Action_List),
		PageTitle:        PageTitle(dm.SymbolMetricsHistory_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.SymbolMetricsHistory_TemplateList, w, r, pageDetail)

}

//SymbolMetricsHistory_HandlerView is the handler used to View a page
func SymbolMetricsHistory_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SymbolMetricsHistory_QueryString)
	_, rD, _ := dao.SymbolMetricsHistory_GetByID(searchID)

	pageDetail := SymbolMetricsHistory_Page{
		Title:       CardTitle(dm.SymbolMetricsHistory_Title, core.Action_View),
		PageTitle:   PageTitle(dm.SymbolMetricsHistory_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.TSKey = rD.TSKey
pageDetail.Symbol = rD.Symbol
pageDetail.Date = rD.Date
pageDetail.Price = rD.Price
pageDetail.URI = rD.URI
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.Type = rD.Type
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

	ExecuteTemplate(dm.SymbolMetricsHistory_TemplateView, w, r, pageDetail)

}

//SymbolMetricsHistory_HandlerEdit is the handler used generate the Edit page
func SymbolMetricsHistory_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SymbolMetricsHistory_QueryString)
	_, rD, _ := dao.SymbolMetricsHistory_GetByID(searchID)
	
	pageDetail := SymbolMetricsHistory_Page{
		Title:       CardTitle(dm.SymbolMetricsHistory_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.SymbolMetricsHistory_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.TSKey = rD.TSKey
pageDetail.Symbol = rD.Symbol
pageDetail.Date = rD.Date
pageDetail.Price = rD.Price
pageDetail.URI = rD.URI
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.Type = rD.Type
pageDetail.Bid = rD.Bid
pageDetail.Offer = rD.Offer
pageDetail.RawPrice = rD.RawPrice
pageDetail.RawBid = rD.RawBid
pageDetail.RawOffer = rD.RawOffer


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.SymbolMetricsHistory_TemplateEdit, w, r, pageDetail)


}

//SymbolMetricsHistory_HandlerSave is the handler used process the saving of an SymbolMetricsHistory
func SymbolMetricsHistory_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("TSKey"))

	var item dm.SymbolMetricsHistory
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.SymbolMetricsHistory_SYSId)
		item.TSKey = r.FormValue(dm.SymbolMetricsHistory_TSKey)
		item.Symbol = r.FormValue(dm.SymbolMetricsHistory_Symbol)
		item.Date = r.FormValue(dm.SymbolMetricsHistory_Date)
		item.Price = r.FormValue(dm.SymbolMetricsHistory_Price)
		item.URI = r.FormValue(dm.SymbolMetricsHistory_URI)
		item.SYSCreated = r.FormValue(dm.SymbolMetricsHistory_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.SymbolMetricsHistory_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.SymbolMetricsHistory_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.SymbolMetricsHistory_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.SymbolMetricsHistory_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.SymbolMetricsHistory_SYSUpdatedHost)
		item.Type = r.FormValue(dm.SymbolMetricsHistory_Type)
		item.Bid = r.FormValue(dm.SymbolMetricsHistory_Bid)
		item.Offer = r.FormValue(dm.SymbolMetricsHistory_Offer)
		item.RawPrice = r.FormValue(dm.SymbolMetricsHistory_RawPrice)
		item.RawBid = r.FormValue(dm.SymbolMetricsHistory_RawBid)
		item.RawOffer = r.FormValue(dm.SymbolMetricsHistory_RawOffer)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.SymbolMetricsHistory_Store(item,r)	

	http.Redirect(w, r, SymbolMetricsHistory_Redirect, http.StatusFound)
}

//SymbolMetricsHistory_HandlerNew is the handler used process the creation of an SymbolMetricsHistory
func SymbolMetricsHistory_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := SymbolMetricsHistory_Page{
		Title:       CardTitle(dm.SymbolMetricsHistory_Title, core.Action_New),
		PageTitle:   PageTitle(dm.SymbolMetricsHistory_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.TSKey = ""
pageDetail.Symbol = ""
pageDetail.Date = ""
pageDetail.Price = ""
pageDetail.URI = ""
pageDetail.SYSCreated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.Type = ""
pageDetail.Bid = ""
pageDetail.Offer = ""
pageDetail.RawPrice = ""
pageDetail.RawBid = ""
pageDetail.RawOffer = ""


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.SymbolMetricsHistory_TemplateNew, w, r, pageDetail)

}

//SymbolMetricsHistory_HandlerDelete is the handler used process the deletion of an SymbolMetricsHistory
func SymbolMetricsHistory_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.SymbolMetricsHistory_QueryString)

	dao.SymbolMetricsHistory_Delete(searchID)	

	http.Redirect(w, r, SymbolMetricsHistory_Redirect, http.StatusFound)
}