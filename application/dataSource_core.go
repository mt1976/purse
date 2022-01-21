package application
// ----------------------------------------------------------------
// Automatically generated  "/application/datasource.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataSource (datasource)
// Endpoint 	        : DataSource (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:36
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/purse/core"
	dao     "github.com/mt1976/purse/dao"
	dm      "github.com/mt1976/purse/datamodel"
	logs    "github.com/mt1976/purse/logs"
)

//datasource_PageList provides the information for the template for a list of DataSources
type DataSource_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DataSource
}

//datasource_Page provides the information for the template for an individual DataSource
type DataSource_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		SYSId string
		Code string
		Name string
		URI string
		Description string
		Notes string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
		AuthKey string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	DataSource_Redirect = dm.DataSource_PathList
)

//DataSource_Publish annouces the endpoints available for this object
func DataSource_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.DataSource_Path, DataSource_Handler)
	mux.HandleFunc(dm.DataSource_PathList, DataSource_HandlerList)
	mux.HandleFunc(dm.DataSource_PathView, DataSource_HandlerView)
	mux.HandleFunc(dm.DataSource_PathEdit, DataSource_HandlerEdit)
	mux.HandleFunc(dm.DataSource_PathNew, DataSource_HandlerNew)
	mux.HandleFunc(dm.DataSource_PathSave, DataSource_HandlerSave)
	mux.HandleFunc(dm.DataSource_PathDelete, DataSource_HandlerDelete)
	logs.Publish("Application", dm.DataSource_Title)
    core.Catalog_Add(dm.DataSource_Title, dm.DataSource_Path, "", dm.DataSource_QueryString, "APP")
}

//DataSource_HandlerList is the handler for the list page
func DataSource_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DataSource
	noItems, returnList, _ := dao.DataSource_GetList()

	pageDetail := DataSource_PageList{
		Title:            CardTitle(dm.DataSource_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DataSource_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DataSource_TemplateList, w, r, pageDetail)

}

//DataSource_HandlerView is the handler used to View a page
func DataSource_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataSource_QueryString)
	_, rD, _ := dao.DataSource_GetByID(searchID)

	pageDetail := DataSource_Page{
		Title:       CardTitle(dm.DataSource_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DataSource_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.URI = rD.URI
pageDetail.Description = rD.Description
pageDetail.Notes = rD.Notes
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.AuthKey = rD.AuthKey


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//


	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DataSource_TemplateView, w, r, pageDetail)

}

//DataSource_HandlerEdit is the handler used generate the Edit page
func DataSource_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataSource_QueryString)
	_, rD, _ := dao.DataSource_GetByID(searchID)
	
	pageDetail := DataSource_Page{
		Title:       CardTitle(dm.DataSource_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DataSource_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.URI = rD.URI
pageDetail.Description = rD.Description
pageDetail.Notes = rD.Notes
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.AuthKey = rD.AuthKey


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DataSource_TemplateEdit, w, r, pageDetail)


}

//DataSource_HandlerSave is the handler used process the saving of an DataSource
func DataSource_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.DataSource
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.DataSource_SYSId)
		item.Code = r.FormValue(dm.DataSource_Code)
		item.Name = r.FormValue(dm.DataSource_Name)
		item.URI = r.FormValue(dm.DataSource_URI)
		item.Description = r.FormValue(dm.DataSource_Description)
		item.Notes = r.FormValue(dm.DataSource_Notes)
		item.SYSCreated = r.FormValue(dm.DataSource_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.DataSource_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.DataSource_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.DataSource_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.DataSource_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.DataSource_SYSUpdatedHost)
		item.AuthKey = r.FormValue(dm.DataSource_AuthKey)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.DataSource_Store(item,r)	

	http.Redirect(w, r, DataSource_Redirect, http.StatusFound)
}

//DataSource_HandlerNew is the handler used process the creation of an DataSource
func DataSource_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := DataSource_Page{
		Title:       CardTitle(dm.DataSource_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DataSource_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.URI = ""
pageDetail.Description = ""
pageDetail.Notes = ""
pageDetail.SYSCreated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.AuthKey = ""


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DataSource_TemplateNew, w, r, pageDetail)

}

//DataSource_HandlerDelete is the handler used process the deletion of an DataSource
func DataSource_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DataSource_QueryString)

	dao.DataSource_Delete(searchID)	

	http.Redirect(w, r, DataSource_Redirect, http.StatusFound)
}