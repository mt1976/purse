package application
// ----------------------------------------------------------------
// Automatically generated  "/application/symboltype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SymbolType (symboltype)
// Endpoint 	        : SymbolType (Code)
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

//symboltype_PageList provides the information for the template for a list of SymbolTypes
type SymbolType_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.SymbolType
}

//symboltype_Page provides the information for the template for an individual SymbolType
type SymbolType_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		SYSId string
		Code string
		Name string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	SymbolType_Redirect = dm.SymbolType_PathList
)

//SymbolType_Publish annouces the endpoints available for this object
func SymbolType_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.SymbolType_Path, SymbolType_Handler)
	mux.HandleFunc(dm.SymbolType_PathList, SymbolType_HandlerList)
	mux.HandleFunc(dm.SymbolType_PathView, SymbolType_HandlerView)
	mux.HandleFunc(dm.SymbolType_PathEdit, SymbolType_HandlerEdit)
	mux.HandleFunc(dm.SymbolType_PathNew, SymbolType_HandlerNew)
	mux.HandleFunc(dm.SymbolType_PathSave, SymbolType_HandlerSave)
	mux.HandleFunc(dm.SymbolType_PathDelete, SymbolType_HandlerDelete)
	logs.Publish("Application", dm.SymbolType_Title)
    core.Catalog_Add(dm.SymbolType_Title, dm.SymbolType_Path, "", dm.SymbolType_QueryString, "APP")
}

//SymbolType_HandlerList is the handler for the list page
func SymbolType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.SymbolType
	noItems, returnList, _ := dao.SymbolType_GetList()

	pageDetail := SymbolType_PageList{
		Title:            CardTitle(dm.SymbolType_Title, core.Action_List),
		PageTitle:        PageTitle(dm.SymbolType_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.SymbolType_TemplateList, w, r, pageDetail)

}

//SymbolType_HandlerView is the handler used to View a page
func SymbolType_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SymbolType_QueryString)
	_, rD, _ := dao.SymbolType_GetByID(searchID)

	pageDetail := SymbolType_Page{
		Title:       CardTitle(dm.SymbolType_Title, core.Action_View),
		PageTitle:   PageTitle(dm.SymbolType_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
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

	ExecuteTemplate(dm.SymbolType_TemplateView, w, r, pageDetail)

}

//SymbolType_HandlerEdit is the handler used generate the Edit page
func SymbolType_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SymbolType_QueryString)
	_, rD, _ := dao.SymbolType_GetByID(searchID)
	
	pageDetail := SymbolType_Page{
		Title:       CardTitle(dm.SymbolType_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.SymbolType_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
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

	ExecuteTemplate(dm.SymbolType_TemplateEdit, w, r, pageDetail)


}

//SymbolType_HandlerSave is the handler used process the saving of an SymbolType
func SymbolType_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.SymbolType
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.SymbolType_SYSId)
		item.Code = r.FormValue(dm.SymbolType_Code)
		item.Name = r.FormValue(dm.SymbolType_Name)
		item.SYSCreated = r.FormValue(dm.SymbolType_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.SymbolType_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.SymbolType_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.SymbolType_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.SymbolType_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.SymbolType_SYSUpdatedHost)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.SymbolType_Store(item,r)	

	http.Redirect(w, r, SymbolType_Redirect, http.StatusFound)
}

//SymbolType_HandlerNew is the handler used process the creation of an SymbolType
func SymbolType_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := SymbolType_Page{
		Title:       CardTitle(dm.SymbolType_Title, core.Action_New),
		PageTitle:   PageTitle(dm.SymbolType_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Code = ""
pageDetail.Name = ""
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

	ExecuteTemplate(dm.SymbolType_TemplateNew, w, r, pageDetail)

}

//SymbolType_HandlerDelete is the handler used process the deletion of an SymbolType
func SymbolType_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.SymbolType_QueryString)

	dao.SymbolType_Delete(searchID)	

	http.Redirect(w, r, SymbolType_Redirect, http.StatusFound)
}