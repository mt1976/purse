package application
// ----------------------------------------------------------------
// Automatically generated  "/application/symbol.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Symbol (symbol)
// Endpoint 	        : Symbol (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:37
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/purse/core"
	dao     "github.com/mt1976/purse/dao"
	dm      "github.com/mt1976/purse/datamodel"
	logs    "github.com/mt1976/purse/logs"
)

//symbol_PageList provides the information for the template for a list of Symbols
type Symbol_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Symbol
}

//symbol_Page provides the information for the template for an individual Symbol
type Symbol_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		SYSId string
		Code string
		Name string
		Type string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
		Identifier string
		DataSource string
		StaticValue string
		Factor string
		DPS string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	Symbol_Redirect = dm.Symbol_PathList
)

//Symbol_Publish annouces the endpoints available for this object
func Symbol_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Symbol_Path, Symbol_Handler)
	mux.HandleFunc(dm.Symbol_PathList, Symbol_HandlerList)
	mux.HandleFunc(dm.Symbol_PathView, Symbol_HandlerView)
	mux.HandleFunc(dm.Symbol_PathEdit, Symbol_HandlerEdit)
	mux.HandleFunc(dm.Symbol_PathNew, Symbol_HandlerNew)
	mux.HandleFunc(dm.Symbol_PathSave, Symbol_HandlerSave)
	mux.HandleFunc(dm.Symbol_PathDelete, Symbol_HandlerDelete)
	logs.Publish("Application", dm.Symbol_Title)
    core.Catalog_Add(dm.Symbol_Title, dm.Symbol_Path, "", dm.Symbol_QueryString, "APP")
}

//Symbol_HandlerList is the handler for the list page
func Symbol_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Symbol
	noItems, returnList, _ := dao.Symbol_GetList()

	pageDetail := Symbol_PageList{
		Title:            CardTitle(dm.Symbol_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Symbol_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Symbol_TemplateList, w, r, pageDetail)

}

//Symbol_HandlerView is the handler used to View a page
func Symbol_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Symbol_QueryString)
	_, rD, _ := dao.Symbol_GetByID(searchID)

	pageDetail := Symbol_Page{
		Title:       CardTitle(dm.Symbol_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Symbol_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Type = rD.Type
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.Identifier = rD.Identifier
pageDetail.DataSource = rD.DataSource
pageDetail.StaticValue = rD.StaticValue
pageDetail.Factor = rD.Factor
pageDetail.DPS = rD.DPS


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//


	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Symbol_TemplateView, w, r, pageDetail)

}

//Symbol_HandlerEdit is the handler used generate the Edit page
func Symbol_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Symbol_QueryString)
	_, rD, _ := dao.Symbol_GetByID(searchID)
	
	pageDetail := Symbol_Page{
		Title:       CardTitle(dm.Symbol_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Symbol_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Type = rD.Type
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.Identifier = rD.Identifier
pageDetail.DataSource = rD.DataSource
pageDetail.StaticValue = rD.StaticValue
pageDetail.Factor = rD.Factor
pageDetail.DPS = rD.DPS


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Symbol_TemplateEdit, w, r, pageDetail)


}

//Symbol_HandlerSave is the handler used process the saving of an Symbol
func Symbol_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Symbol
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Symbol_SYSId)
		item.Code = r.FormValue(dm.Symbol_Code)
		item.Name = r.FormValue(dm.Symbol_Name)
		item.Type = r.FormValue(dm.Symbol_Type)
		item.SYSCreated = r.FormValue(dm.Symbol_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.Symbol_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Symbol_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.Symbol_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.Symbol_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Symbol_SYSUpdatedHost)
		item.Identifier = r.FormValue(dm.Symbol_Identifier)
		item.DataSource = r.FormValue(dm.Symbol_DataSource)
		item.StaticValue = r.FormValue(dm.Symbol_StaticValue)
		item.Factor = r.FormValue(dm.Symbol_Factor)
		item.DPS = r.FormValue(dm.Symbol_DPS)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.Symbol_Store(item,r)	

	http.Redirect(w, r, Symbol_Redirect, http.StatusFound)
}

//Symbol_HandlerNew is the handler used process the creation of an Symbol
func Symbol_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Symbol_Page{
		Title:       CardTitle(dm.Symbol_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Symbol_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.Type = ""
pageDetail.SYSCreated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.Identifier = ""
pageDetail.DataSource = ""
pageDetail.StaticValue = ""
pageDetail.Factor = ""
pageDetail.DPS = ""


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Symbol_TemplateNew, w, r, pageDetail)

}

//Symbol_HandlerDelete is the handler used process the deletion of an Symbol
func Symbol_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Symbol_QueryString)

	dao.Symbol_Delete(searchID)	

	http.Redirect(w, r, Symbol_Redirect, http.StatusFound)
}