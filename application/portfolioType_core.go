package application
// ----------------------------------------------------------------
// Automatically generated  "/application/portfoliotype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : PortfolioType (portfoliotype)
// Endpoint 	        : PortfolioType (Code)
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

//portfoliotype_PageList provides the information for the template for a list of PortfolioTypes
type PortfolioType_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.PortfolioType
}

//portfoliotype_Page provides the information for the template for an individual PortfolioType
type PortfolioType_Page struct {
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
	PortfolioType_Redirect = dm.PortfolioType_PathList
)

//PortfolioType_Publish annouces the endpoints available for this object
func PortfolioType_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.PortfolioType_Path, PortfolioType_Handler)
	mux.HandleFunc(dm.PortfolioType_PathList, PortfolioType_HandlerList)
	mux.HandleFunc(dm.PortfolioType_PathView, PortfolioType_HandlerView)
	mux.HandleFunc(dm.PortfolioType_PathEdit, PortfolioType_HandlerEdit)
	mux.HandleFunc(dm.PortfolioType_PathNew, PortfolioType_HandlerNew)
	mux.HandleFunc(dm.PortfolioType_PathSave, PortfolioType_HandlerSave)
	mux.HandleFunc(dm.PortfolioType_PathDelete, PortfolioType_HandlerDelete)
	logs.Publish("Application", dm.PortfolioType_Title)
    core.Catalog_Add(dm.PortfolioType_Title, dm.PortfolioType_Path, "", dm.PortfolioType_QueryString, "APP")
}

//PortfolioType_HandlerList is the handler for the list page
func PortfolioType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.PortfolioType
	noItems, returnList, _ := dao.PortfolioType_GetList()

	pageDetail := PortfolioType_PageList{
		Title:            CardTitle(dm.PortfolioType_Title, core.Action_List),
		PageTitle:        PageTitle(dm.PortfolioType_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.PortfolioType_TemplateList, w, r, pageDetail)

}

//PortfolioType_HandlerView is the handler used to View a page
func PortfolioType_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.PortfolioType_QueryString)
	_, rD, _ := dao.PortfolioType_GetByID(searchID)

	pageDetail := PortfolioType_Page{
		Title:       CardTitle(dm.PortfolioType_Title, core.Action_View),
		PageTitle:   PageTitle(dm.PortfolioType_Title, core.Action_View),
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

	ExecuteTemplate(dm.PortfolioType_TemplateView, w, r, pageDetail)

}

//PortfolioType_HandlerEdit is the handler used generate the Edit page
func PortfolioType_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.PortfolioType_QueryString)
	_, rD, _ := dao.PortfolioType_GetByID(searchID)
	
	pageDetail := PortfolioType_Page{
		Title:       CardTitle(dm.PortfolioType_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.PortfolioType_Title, core.Action_Edit),
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

	ExecuteTemplate(dm.PortfolioType_TemplateEdit, w, r, pageDetail)


}

//PortfolioType_HandlerSave is the handler used process the saving of an PortfolioType
func PortfolioType_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.PortfolioType
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.PortfolioType_SYSId)
		item.Code = r.FormValue(dm.PortfolioType_Code)
		item.Name = r.FormValue(dm.PortfolioType_Name)
		item.SYSCreated = r.FormValue(dm.PortfolioType_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.PortfolioType_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.PortfolioType_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.PortfolioType_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.PortfolioType_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.PortfolioType_SYSUpdatedHost)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.PortfolioType_Store(item,r)	

	http.Redirect(w, r, PortfolioType_Redirect, http.StatusFound)
}

//PortfolioType_HandlerNew is the handler used process the creation of an PortfolioType
func PortfolioType_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := PortfolioType_Page{
		Title:       CardTitle(dm.PortfolioType_Title, core.Action_New),
		PageTitle:   PageTitle(dm.PortfolioType_Title, core.Action_New),
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

	ExecuteTemplate(dm.PortfolioType_TemplateNew, w, r, pageDetail)

}

//PortfolioType_HandlerDelete is the handler used process the deletion of an PortfolioType
func PortfolioType_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.PortfolioType_QueryString)

	dao.PortfolioType_Delete(searchID)	

	http.Redirect(w, r, PortfolioType_Redirect, http.StatusFound)
}