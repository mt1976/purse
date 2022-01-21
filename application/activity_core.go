package application
// ----------------------------------------------------------------
// Automatically generated  "/application/activity.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Activity (activity)
// Endpoint 	        : Activity (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:35
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/purse/core"
	dao     "github.com/mt1976/purse/dao"
	dm      "github.com/mt1976/purse/datamodel"
	logs    "github.com/mt1976/purse/logs"
)

//activity_PageList provides the information for the template for a list of Activitys
type Activity_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Activity
}

//activity_Page provides the information for the template for an individual Activity
type Activity_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		SYSId string
		Code string
		Symbol string
		Date string
		Type string
		Portfolio string
		Amount string
		Price string
		Notes string
		Holding string
		ValueAMT string
		ValueCCY string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	Activity_Redirect = dm.Activity_PathList
)

//Activity_Publish annouces the endpoints available for this object
func Activity_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Activity_Path, Activity_Handler)
	mux.HandleFunc(dm.Activity_PathList, Activity_HandlerList)
	mux.HandleFunc(dm.Activity_PathView, Activity_HandlerView)
	mux.HandleFunc(dm.Activity_PathEdit, Activity_HandlerEdit)
	mux.HandleFunc(dm.Activity_PathNew, Activity_HandlerNew)
	mux.HandleFunc(dm.Activity_PathSave, Activity_HandlerSave)
	mux.HandleFunc(dm.Activity_PathDelete, Activity_HandlerDelete)
	logs.Publish("Application", dm.Activity_Title)
    core.Catalog_Add(dm.Activity_Title, dm.Activity_Path, "", dm.Activity_QueryString, "APP")
}

//Activity_HandlerList is the handler for the list page
func Activity_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Activity
	noItems, returnList, _ := dao.Activity_GetList()

	pageDetail := Activity_PageList{
		Title:            CardTitle(dm.Activity_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Activity_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Activity_TemplateList, w, r, pageDetail)

}

//Activity_HandlerView is the handler used to View a page
func Activity_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Activity_QueryString)
	_, rD, _ := dao.Activity_GetByID(searchID)

	pageDetail := Activity_Page{
		Title:       CardTitle(dm.Activity_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Activity_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Symbol = rD.Symbol
pageDetail.Date = rD.Date
pageDetail.Type = rD.Type
pageDetail.Portfolio = rD.Portfolio
pageDetail.Amount = rD.Amount
pageDetail.Price = rD.Price
pageDetail.Notes = rD.Notes
pageDetail.Holding = rD.Holding
pageDetail.ValueAMT = rD.ValueAMT
pageDetail.ValueCCY = rD.ValueCCY
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//


	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Activity_TemplateView, w, r, pageDetail)

}

//Activity_HandlerEdit is the handler used generate the Edit page
func Activity_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Activity_QueryString)
	_, rD, _ := dao.Activity_GetByID(searchID)
	
	pageDetail := Activity_Page{
		Title:       CardTitle(dm.Activity_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Activity_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Symbol = rD.Symbol
pageDetail.Date = rD.Date
pageDetail.Type = rD.Type
pageDetail.Portfolio = rD.Portfolio
pageDetail.Amount = rD.Amount
pageDetail.Price = rD.Price
pageDetail.Notes = rD.Notes
pageDetail.Holding = rD.Holding
pageDetail.ValueAMT = rD.ValueAMT
pageDetail.ValueCCY = rD.ValueCCY
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Activity_TemplateEdit, w, r, pageDetail)


}

//Activity_HandlerSave is the handler used process the saving of an Activity
func Activity_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Activity
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Activity_SYSId)
		item.Code = r.FormValue(dm.Activity_Code)
		item.Symbol = r.FormValue(dm.Activity_Symbol)
		item.Date = r.FormValue(dm.Activity_Date)
		item.Type = r.FormValue(dm.Activity_Type)
		item.Portfolio = r.FormValue(dm.Activity_Portfolio)
		item.Amount = r.FormValue(dm.Activity_Amount)
		item.Price = r.FormValue(dm.Activity_Price)
		item.Notes = r.FormValue(dm.Activity_Notes)
		item.Holding = r.FormValue(dm.Activity_Holding)
		item.ValueAMT = r.FormValue(dm.Activity_ValueAMT)
		item.ValueCCY = r.FormValue(dm.Activity_ValueCCY)
		item.SYSCreated = r.FormValue(dm.Activity_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.Activity_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Activity_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.Activity_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.Activity_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Activity_SYSUpdatedHost)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.Activity_Store(item,r)	

	http.Redirect(w, r, Activity_Redirect, http.StatusFound)
}

//Activity_HandlerNew is the handler used process the creation of an Activity
func Activity_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Activity_Page{
		Title:       CardTitle(dm.Activity_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Activity_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Code = ""
pageDetail.Symbol = ""
pageDetail.Date = ""
pageDetail.Type = ""
pageDetail.Portfolio = ""
pageDetail.Amount = ""
pageDetail.Price = ""
pageDetail.Notes = ""
pageDetail.Holding = ""
pageDetail.ValueAMT = ""
pageDetail.ValueCCY = ""
pageDetail.SYSCreated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Activity_TemplateNew, w, r, pageDetail)

}

//Activity_HandlerDelete is the handler used process the deletion of an Activity
func Activity_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Activity_QueryString)

	dao.Activity_Delete(searchID)	

	http.Redirect(w, r, Activity_Redirect, http.StatusFound)
}