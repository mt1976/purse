package application
// ----------------------------------------------------------------
// Automatically generated  "/application/watchlist.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Watchlist (watchlist)
// Endpoint 	        : Watchlist (Id)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:39
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/purse/core"
	dao     "github.com/mt1976/purse/dao"
	dm      "github.com/mt1976/purse/datamodel"
	logs    "github.com/mt1976/purse/logs"
)

//watchlist_PageList provides the information for the template for a list of Watchlists
type Watchlist_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Watchlist
}

//watchlist_Page provides the information for the template for an individual Watchlist
type Watchlist_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Code string
		Name string
		Symbol string
		Notes string
		SYSCreated string
		SYSUpdated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	Watchlist_Redirect = dm.Watchlist_PathList
)

//Watchlist_Publish annouces the endpoints available for this object
func Watchlist_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Watchlist_Path, Watchlist_Handler)
	mux.HandleFunc(dm.Watchlist_PathList, Watchlist_HandlerList)
	mux.HandleFunc(dm.Watchlist_PathView, Watchlist_HandlerView)
	mux.HandleFunc(dm.Watchlist_PathEdit, Watchlist_HandlerEdit)
	mux.HandleFunc(dm.Watchlist_PathNew, Watchlist_HandlerNew)
	mux.HandleFunc(dm.Watchlist_PathSave, Watchlist_HandlerSave)
	mux.HandleFunc(dm.Watchlist_PathDelete, Watchlist_HandlerDelete)
	logs.Publish("Application", dm.Watchlist_Title)
    core.Catalog_Add(dm.Watchlist_Title, dm.Watchlist_Path, "", dm.Watchlist_QueryString, "APP")
}

//Watchlist_HandlerList is the handler for the list page
func Watchlist_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Watchlist
	noItems, returnList, _ := dao.Watchlist_GetList()

	pageDetail := Watchlist_PageList{
		Title:            CardTitle(dm.Watchlist_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Watchlist_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Watchlist_TemplateList, w, r, pageDetail)

}

//Watchlist_HandlerView is the handler used to View a page
func Watchlist_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Watchlist_QueryString)
	_, rD, _ := dao.Watchlist_GetByID(searchID)

	pageDetail := Watchlist_Page{
		Title:       CardTitle(dm.Watchlist_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Watchlist_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Symbol = rD.Symbol
pageDetail.Notes = rD.Notes
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//


	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Watchlist_TemplateView, w, r, pageDetail)

}

//Watchlist_HandlerEdit is the handler used generate the Edit page
func Watchlist_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Watchlist_QueryString)
	_, rD, _ := dao.Watchlist_GetByID(searchID)
	
	pageDetail := Watchlist_Page{
		Title:       CardTitle(dm.Watchlist_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Watchlist_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Symbol = rD.Symbol
pageDetail.Notes = rD.Notes
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Watchlist_TemplateEdit, w, r, pageDetail)


}

//Watchlist_HandlerSave is the handler used process the saving of an Watchlist
func Watchlist_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Watchlist
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Watchlist_SYSId)
		item.Id = r.FormValue(dm.Watchlist_Id)
		item.Code = r.FormValue(dm.Watchlist_Code)
		item.Name = r.FormValue(dm.Watchlist_Name)
		item.Symbol = r.FormValue(dm.Watchlist_Symbol)
		item.Notes = r.FormValue(dm.Watchlist_Notes)
		item.SYSCreated = r.FormValue(dm.Watchlist_SYSCreated)
		item.SYSUpdated = r.FormValue(dm.Watchlist_SYSUpdated)
		item.SYSCreatedBy = r.FormValue(dm.Watchlist_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Watchlist_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Watchlist_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Watchlist_SYSUpdatedHost)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.Watchlist_Store(item,r)	

	http.Redirect(w, r, Watchlist_Redirect, http.StatusFound)
}

//Watchlist_HandlerNew is the handler used process the creation of an Watchlist
func Watchlist_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Watchlist_Page{
		Title:       CardTitle(dm.Watchlist_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Watchlist_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.Symbol = ""
pageDetail.Notes = ""
pageDetail.SYSCreated = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Watchlist_TemplateNew, w, r, pageDetail)

}

//Watchlist_HandlerDelete is the handler used process the deletion of an Watchlist
func Watchlist_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Watchlist_QueryString)

	dao.Watchlist_Delete(searchID)	

	http.Redirect(w, r, Watchlist_Redirect, http.StatusFound)
}