package application
// ----------------------------------------------------------------
// Automatically generated  "/application/portfolio.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Portfolio (portfolio)
// Endpoint 	        : Portfolio (Code)
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

//portfolio_PageList provides the information for the template for a list of Portfolios
type Portfolio_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Portfolio
}

//portfolio_Page provides the information for the template for an individual Portfolio
type Portfolio_Page struct {
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
		DefaultModel string
		Description string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
		Type_Lookup string
		DefaultModel_Lookup string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Type_Lookup_List	[]dm.PortfolioType
	DefaultModel_Lookup_List	[]dm.PortfolioModel
	
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	Portfolio_Redirect = dm.Portfolio_PathList
)

//Portfolio_Publish annouces the endpoints available for this object
func Portfolio_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Portfolio_Path, Portfolio_Handler)
	mux.HandleFunc(dm.Portfolio_PathList, Portfolio_HandlerList)
	mux.HandleFunc(dm.Portfolio_PathView, Portfolio_HandlerView)
	mux.HandleFunc(dm.Portfolio_PathEdit, Portfolio_HandlerEdit)
	mux.HandleFunc(dm.Portfolio_PathNew, Portfolio_HandlerNew)
	mux.HandleFunc(dm.Portfolio_PathSave, Portfolio_HandlerSave)
	mux.HandleFunc(dm.Portfolio_PathDelete, Portfolio_HandlerDelete)
	logs.Publish("Application", dm.Portfolio_Title)
    core.Catalog_Add(dm.Portfolio_Title, dm.Portfolio_Path, "", dm.Portfolio_QueryString, "APP")
}

//Portfolio_HandlerList is the handler for the list page
func Portfolio_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Portfolio
	noItems, returnList, _ := dao.Portfolio_GetList()

	pageDetail := Portfolio_PageList{
		Title:            CardTitle(dm.Portfolio_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Portfolio_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Portfolio_TemplateList, w, r, pageDetail)

}

//Portfolio_HandlerView is the handler used to View a page
func Portfolio_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Portfolio_QueryString)
	_, rD, _ := dao.Portfolio_GetByID(searchID)

	pageDetail := Portfolio_Page{
		Title:       CardTitle(dm.Portfolio_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Portfolio_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Type = rD.Type
pageDetail.DefaultModel = rD.DefaultModel
pageDetail.Description = rD.Description
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
_,Type_Lookup_Name,_:= dao.PortfolioType_GetByID(rD.Type)
pageDetail.Type_Lookup = Type_Lookup_Name.Name
_,DefaultModel_Lookup_Name,_:= dao.PortfolioModel_GetByID(rD.DefaultModel)
pageDetail.DefaultModel_Lookup = DefaultModel_Lookup_Name.Name
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//


	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Portfolio_TemplateView, w, r, pageDetail)

}

//Portfolio_HandlerEdit is the handler used generate the Edit page
func Portfolio_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Portfolio_QueryString)
	_, rD, _ := dao.Portfolio_GetByID(searchID)
	
	pageDetail := Portfolio_Page{
		Title:       CardTitle(dm.Portfolio_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Portfolio_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Type = rD.Type
pageDetail.DefaultModel = rD.DefaultModel
pageDetail.Description = rD.Description
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
_,Type_Lookup_Name,_:= dao.PortfolioType_GetByID(rD.Type)
pageDetail.Type_Lookup = Type_Lookup_Name.Name
_,pageDetail.Type_Lookup_List,_ = dao.PortfolioType_GetList()
_,DefaultModel_Lookup_Name,_:= dao.PortfolioModel_GetByID(rD.DefaultModel)
pageDetail.DefaultModel_Lookup = DefaultModel_Lookup_Name.Name
_,pageDetail.DefaultModel_Lookup_List,_ = dao.PortfolioModel_GetList()
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Portfolio_TemplateEdit, w, r, pageDetail)


}

//Portfolio_HandlerSave is the handler used process the saving of an Portfolio
func Portfolio_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Portfolio
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Portfolio_SYSId)
		item.Code = r.FormValue(dm.Portfolio_Code)
		item.Name = r.FormValue(dm.Portfolio_Name)
		item.Type = r.FormValue(dm.Portfolio_Type)
		item.DefaultModel = r.FormValue(dm.Portfolio_DefaultModel)
		item.Description = r.FormValue(dm.Portfolio_Description)
		item.SYSCreated = r.FormValue(dm.Portfolio_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.Portfolio_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Portfolio_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.Portfolio_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.Portfolio_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Portfolio_SYSUpdatedHost)
		item.Type_Lookup = r.FormValue(dm.Portfolio_Type_Lookup)
		item.DefaultModel_Lookup = r.FormValue(dm.Portfolio_DefaultModel_Lookup)
		
		
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.Portfolio_Store(item,r)	

	http.Redirect(w, r, Portfolio_Redirect, http.StatusFound)
}

//Portfolio_HandlerNew is the handler used process the creation of an Portfolio
func Portfolio_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Portfolio_Page{
		Title:       CardTitle(dm.Portfolio_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Portfolio_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.Type = ""
pageDetail.DefaultModel = ""
pageDetail.Description = ""
pageDetail.SYSCreated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Type_Lookup = ""
_,pageDetail.Type_Lookup_List,_ = dao.PortfolioType_GetList()
pageDetail.DefaultModel_Lookup = ""
_,pageDetail.DefaultModel_Lookup_List,_ = dao.PortfolioModel_GetList()
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Portfolio_TemplateNew, w, r, pageDetail)

}

//Portfolio_HandlerDelete is the handler used process the deletion of an Portfolio
func Portfolio_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Portfolio_QueryString)

	dao.Portfolio_Delete(searchID)	

	http.Redirect(w, r, Portfolio_Redirect, http.StatusFound)
}