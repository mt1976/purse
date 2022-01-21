package application
// ----------------------------------------------------------------
// Automatically generated  "/application/holdingtype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : HoldingType (holdingtype)
// Endpoint 	        : HoldingType (Code)
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

//holdingtype_PageList provides the information for the template for a list of HoldingTypes
type HoldingType_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.HoldingType
}

//holdingtype_Page provides the information for the template for an individual HoldingType
type HoldingType_Page struct {
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
	HoldingType_Redirect = dm.HoldingType_PathList
)

//HoldingType_Publish annouces the endpoints available for this object
func HoldingType_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.HoldingType_Path, HoldingType_Handler)
	mux.HandleFunc(dm.HoldingType_PathList, HoldingType_HandlerList)
	mux.HandleFunc(dm.HoldingType_PathView, HoldingType_HandlerView)
	mux.HandleFunc(dm.HoldingType_PathEdit, HoldingType_HandlerEdit)
	mux.HandleFunc(dm.HoldingType_PathNew, HoldingType_HandlerNew)
	mux.HandleFunc(dm.HoldingType_PathSave, HoldingType_HandlerSave)
	mux.HandleFunc(dm.HoldingType_PathDelete, HoldingType_HandlerDelete)
	logs.Publish("Application", dm.HoldingType_Title)
    core.Catalog_Add(dm.HoldingType_Title, dm.HoldingType_Path, "", dm.HoldingType_QueryString, "APP")
}

//HoldingType_HandlerList is the handler for the list page
func HoldingType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.HoldingType
	noItems, returnList, _ := dao.HoldingType_GetList()

	pageDetail := HoldingType_PageList{
		Title:            CardTitle(dm.HoldingType_Title, core.Action_List),
		PageTitle:        PageTitle(dm.HoldingType_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.HoldingType_TemplateList, w, r, pageDetail)

}

//HoldingType_HandlerView is the handler used to View a page
func HoldingType_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.HoldingType_QueryString)
	_, rD, _ := dao.HoldingType_GetByID(searchID)

	pageDetail := HoldingType_Page{
		Title:       CardTitle(dm.HoldingType_Title, core.Action_View),
		PageTitle:   PageTitle(dm.HoldingType_Title, core.Action_View),
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

	ExecuteTemplate(dm.HoldingType_TemplateView, w, r, pageDetail)

}

//HoldingType_HandlerEdit is the handler used generate the Edit page
func HoldingType_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.HoldingType_QueryString)
	_, rD, _ := dao.HoldingType_GetByID(searchID)
	
	pageDetail := HoldingType_Page{
		Title:       CardTitle(dm.HoldingType_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.HoldingType_Title, core.Action_Edit),
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

	ExecuteTemplate(dm.HoldingType_TemplateEdit, w, r, pageDetail)


}

//HoldingType_HandlerSave is the handler used process the saving of an HoldingType
func HoldingType_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.HoldingType
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.HoldingType_SYSId)
		item.Code = r.FormValue(dm.HoldingType_Code)
		item.Name = r.FormValue(dm.HoldingType_Name)
		item.Factor = r.FormValue(dm.HoldingType_Factor)
		item.SYSCreated = r.FormValue(dm.HoldingType_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.HoldingType_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.HoldingType_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.HoldingType_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.HoldingType_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.HoldingType_SYSUpdatedHost)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.HoldingType_Store(item,r)	

	http.Redirect(w, r, HoldingType_Redirect, http.StatusFound)
}

//HoldingType_HandlerNew is the handler used process the creation of an HoldingType
func HoldingType_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := HoldingType_Page{
		Title:       CardTitle(dm.HoldingType_Title, core.Action_New),
		PageTitle:   PageTitle(dm.HoldingType_Title, core.Action_New),
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

	ExecuteTemplate(dm.HoldingType_TemplateNew, w, r, pageDetail)

}

//HoldingType_HandlerDelete is the handler used process the deletion of an HoldingType
func HoldingType_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.HoldingType_QueryString)

	dao.HoldingType_Delete(searchID)	

	http.Redirect(w, r, HoldingType_Redirect, http.StatusFound)
}