package application
// ----------------------------------------------------------------
// Automatically generated  "/application/portfoliomodel.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : PortfolioModel (portfoliomodel)
// Endpoint 	        : PortfolioModel (Code)
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

//portfoliomodel_PageList provides the information for the template for a list of PortfolioModels
type PortfolioModel_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.PortfolioModel
}

//portfoliomodel_Page provides the information for the template for an individual PortfolioModel
type PortfolioModel_Page struct {
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
	PortfolioModel_Redirect = dm.PortfolioModel_PathList
)

//PortfolioModel_Publish annouces the endpoints available for this object
func PortfolioModel_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.PortfolioModel_Path, PortfolioModel_Handler)
	mux.HandleFunc(dm.PortfolioModel_PathList, PortfolioModel_HandlerList)
	mux.HandleFunc(dm.PortfolioModel_PathView, PortfolioModel_HandlerView)
	mux.HandleFunc(dm.PortfolioModel_PathEdit, PortfolioModel_HandlerEdit)
	mux.HandleFunc(dm.PortfolioModel_PathNew, PortfolioModel_HandlerNew)
	mux.HandleFunc(dm.PortfolioModel_PathSave, PortfolioModel_HandlerSave)
	mux.HandleFunc(dm.PortfolioModel_PathDelete, PortfolioModel_HandlerDelete)
	logs.Publish("Application", dm.PortfolioModel_Title)
    core.Catalog_Add(dm.PortfolioModel_Title, dm.PortfolioModel_Path, "", dm.PortfolioModel_QueryString, "APP")
}

//PortfolioModel_HandlerList is the handler for the list page
func PortfolioModel_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.PortfolioModel
	noItems, returnList, _ := dao.PortfolioModel_GetList()

	pageDetail := PortfolioModel_PageList{
		Title:            CardTitle(dm.PortfolioModel_Title, core.Action_List),
		PageTitle:        PageTitle(dm.PortfolioModel_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.PortfolioModel_TemplateList, w, r, pageDetail)

}

//PortfolioModel_HandlerView is the handler used to View a page
func PortfolioModel_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.PortfolioModel_QueryString)
	_, rD, _ := dao.PortfolioModel_GetByID(searchID)

	pageDetail := PortfolioModel_Page{
		Title:       CardTitle(dm.PortfolioModel_Title, core.Action_View),
		PageTitle:   PageTitle(dm.PortfolioModel_Title, core.Action_View),
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

	ExecuteTemplate(dm.PortfolioModel_TemplateView, w, r, pageDetail)

}

//PortfolioModel_HandlerEdit is the handler used generate the Edit page
func PortfolioModel_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.PortfolioModel_QueryString)
	_, rD, _ := dao.PortfolioModel_GetByID(searchID)
	
	pageDetail := PortfolioModel_Page{
		Title:       CardTitle(dm.PortfolioModel_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.PortfolioModel_Title, core.Action_Edit),
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

	ExecuteTemplate(dm.PortfolioModel_TemplateEdit, w, r, pageDetail)


}

//PortfolioModel_HandlerSave is the handler used process the saving of an PortfolioModel
func PortfolioModel_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.PortfolioModel
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.PortfolioModel_SYSId)
		item.Code = r.FormValue(dm.PortfolioModel_Code)
		item.Name = r.FormValue(dm.PortfolioModel_Name)
		item.SYSCreated = r.FormValue(dm.PortfolioModel_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.PortfolioModel_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.PortfolioModel_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.PortfolioModel_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.PortfolioModel_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.PortfolioModel_SYSUpdatedHost)
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.PortfolioModel_Store(item,r)	

	http.Redirect(w, r, PortfolioModel_Redirect, http.StatusFound)
}

//PortfolioModel_HandlerNew is the handler used process the creation of an PortfolioModel
func PortfolioModel_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := PortfolioModel_Page{
		Title:       CardTitle(dm.PortfolioModel_Title, core.Action_New),
		PageTitle:   PageTitle(dm.PortfolioModel_Title, core.Action_New),
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

	ExecuteTemplate(dm.PortfolioModel_TemplateNew, w, r, pageDetail)

}

//PortfolioModel_HandlerDelete is the handler used process the deletion of an PortfolioModel
func PortfolioModel_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.PortfolioModel_QueryString)

	dao.PortfolioModel_Delete(searchID)	

	http.Redirect(w, r, PortfolioModel_Redirect, http.StatusFound)
}