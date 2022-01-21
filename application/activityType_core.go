package application
// ----------------------------------------------------------------
// Automatically generated  "/application/activitytype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ActivityType (activitytype)
// Endpoint 	        : ActivityType (Code)
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

//activitytype_PageList provides the information for the template for a list of ActivityTypes
type ActivityType_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.ActivityType
}

//activitytype_Page provides the information for the template for an individual ActivityType
type ActivityType_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		SYSId string
		Code string
		Name string
		Factor string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	ActivityType_Redirect = dm.ActivityType_PathList
)

//ActivityType_Publish annouces the endpoints available for this object
func ActivityType_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.ActivityType_Path, ActivityType_Handler)
	mux.HandleFunc(dm.ActivityType_PathList, ActivityType_HandlerList)
	mux.HandleFunc(dm.ActivityType_PathView, ActivityType_HandlerView)
	mux.HandleFunc(dm.ActivityType_PathEdit, ActivityType_HandlerEdit)
	mux.HandleFunc(dm.ActivityType_PathNew, ActivityType_HandlerNew)
	mux.HandleFunc(dm.ActivityType_PathSave, ActivityType_HandlerSave)
	mux.HandleFunc(dm.ActivityType_PathDelete, ActivityType_HandlerDelete)
	logs.Publish("Application", dm.ActivityType_Title)
    core.Catalog_Add(dm.ActivityType_Title, dm.ActivityType_Path, "", dm.ActivityType_QueryString, "APP")
}

//ActivityType_HandlerList is the handler for the list page
func ActivityType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.ActivityType
	noItems, returnList, _ := dao.ActivityType_GetList()

	pageDetail := ActivityType_PageList{
		Title:            CardTitle(dm.ActivityType_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ActivityType_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.ActivityType_TemplateList, w, r, pageDetail)

}

//ActivityType_HandlerView is the handler used to View a page
func ActivityType_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ActivityType_QueryString)
	_, rD, _ := dao.ActivityType_GetByID(searchID)

	pageDetail := ActivityType_Page{
		Title:       CardTitle(dm.ActivityType_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ActivityType_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Factor = rD.Factor
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

	ExecuteTemplate(dm.ActivityType_TemplateView, w, r, pageDetail)

}

//ActivityType_HandlerEdit is the handler used generate the Edit page
func ActivityType_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ActivityType_QueryString)
	_, rD, _ := dao.ActivityType_GetByID(searchID)
	
	pageDetail := ActivityType_Page{
		Title:       CardTitle(dm.ActivityType_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.ActivityType_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Factor = rD.Factor
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

	ExecuteTemplate(dm.ActivityType_TemplateEdit, w, r, pageDetail)


}

//ActivityType_HandlerSave is the handler used process the saving of an ActivityType
func ActivityType_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.ActivityType
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.ActivityType_SYSId)
		item.Code = r.FormValue(dm.ActivityType_Code)
		item.Name = r.FormValue(dm.ActivityType_Name)
		item.Factor = r.FormValue(dm.ActivityType_Factor)
		item.SYSCreated = r.FormValue(dm.ActivityType_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.ActivityType_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.ActivityType_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.ActivityType_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.ActivityType_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.ActivityType_SYSUpdatedHost)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.ActivityType_Store(item,r)	

	http.Redirect(w, r, ActivityType_Redirect, http.StatusFound)
}

//ActivityType_HandlerNew is the handler used process the creation of an ActivityType
func ActivityType_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := ActivityType_Page{
		Title:       CardTitle(dm.ActivityType_Title, core.Action_New),
		PageTitle:   PageTitle(dm.ActivityType_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.Factor = ""
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

	ExecuteTemplate(dm.ActivityType_TemplateNew, w, r, pageDetail)

}

//ActivityType_HandlerDelete is the handler used process the deletion of an ActivityType
func ActivityType_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.ActivityType_QueryString)

	dao.ActivityType_Delete(searchID)	

	http.Redirect(w, r, ActivityType_Redirect, http.StatusFound)
}