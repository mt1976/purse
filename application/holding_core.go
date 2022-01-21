package application
// ----------------------------------------------------------------
// Automatically generated  "/application/holding.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Holding (holding)
// Endpoint 	        : Holding (Code)
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

//holding_PageList provides the information for the template for a list of Holdings
type Holding_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Holding
}

//holding_Page provides the information for the template for an individual Holding
type Holding_Page struct {
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
		Portfolio string
		Description string
		Amount string
		Instrument string
		Price string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
		Symbol string
		Type_Lookup string
		Portfolio_Lookup string
	
	
		Instrument_Lookup string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Type_Lookup_List	[]dm.HoldingType
	Portfolio_Lookup_List	[]dm.Portfolio
	
	
	Instrument_Lookup_List	[]dm.Watchlist
	
	
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
}

const (
	Holding_Redirect = dm.Holding_PathList
)

//Holding_Publish annouces the endpoints available for this object
func Holding_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Holding_Path, Holding_Handler)
	mux.HandleFunc(dm.Holding_PathList, Holding_HandlerList)
	mux.HandleFunc(dm.Holding_PathView, Holding_HandlerView)
	mux.HandleFunc(dm.Holding_PathEdit, Holding_HandlerEdit)
	mux.HandleFunc(dm.Holding_PathNew, Holding_HandlerNew)
	mux.HandleFunc(dm.Holding_PathSave, Holding_HandlerSave)
	mux.HandleFunc(dm.Holding_PathDelete, Holding_HandlerDelete)
	logs.Publish("Application", dm.Holding_Title)
    core.Catalog_Add(dm.Holding_Title, dm.Holding_Path, "", dm.Holding_QueryString, "APP")
}

//Holding_HandlerList is the handler for the list page
func Holding_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Holding
	noItems, returnList, _ := dao.Holding_GetList()

	pageDetail := Holding_PageList{
		Title:            CardTitle(dm.Holding_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Holding_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Holding_TemplateList, w, r, pageDetail)

}

//Holding_HandlerView is the handler used to View a page
func Holding_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Holding_QueryString)
	_, rD, _ := dao.Holding_GetByID(searchID)

	pageDetail := Holding_Page{
		Title:       CardTitle(dm.Holding_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Holding_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Type = rD.Type
pageDetail.Portfolio = rD.Portfolio
pageDetail.Description = rD.Description
pageDetail.Amount = rD.Amount
pageDetail.Instrument = rD.Instrument
pageDetail.Price = rD.Price
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.Symbol = rD.Symbol


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
_,Type_Lookup_Name,_:= dao.HoldingType_GetByID(rD.Type)
pageDetail.Type_Lookup = Type_Lookup_Name.Name
_,Portfolio_Lookup_Name,_:= dao.Portfolio_GetByID(rD.Portfolio)
pageDetail.Portfolio_Lookup = Portfolio_Lookup_Name.Name
_,Code_Lookup_Name,_:= dao.Watchlist_GetByID(rD.Code)
pageDetail.Instrument_Lookup = Code_Lookup_Name.Name
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//


	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Holding_TemplateView, w, r, pageDetail)

}

//Holding_HandlerEdit is the handler used generate the Edit page
func Holding_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Holding_QueryString)
	_, rD, _ := dao.Holding_GetByID(searchID)
	
	pageDetail := Holding_Page{
		Title:       CardTitle(dm.Holding_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Holding_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Type = rD.Type
pageDetail.Portfolio = rD.Portfolio
pageDetail.Description = rD.Description
pageDetail.Amount = rD.Amount
pageDetail.Instrument = rD.Instrument
pageDetail.Price = rD.Price
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.Symbol = rD.Symbol


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
_,Type_Lookup_Name,_:= dao.HoldingType_GetByID(rD.Type)
pageDetail.Type_Lookup = Type_Lookup_Name.Name
_,pageDetail.Type_Lookup_List,_ = dao.HoldingType_GetList()
_,Portfolio_Lookup_Name,_:= dao.Portfolio_GetByID(rD.Portfolio)
pageDetail.Portfolio_Lookup = Portfolio_Lookup_Name.Name
_,pageDetail.Portfolio_Lookup_List,_ = dao.Portfolio_GetList()
_,Code_Lookup_Name,_:= dao.Watchlist_GetByID(rD.Code)
pageDetail.Instrument_Lookup = Code_Lookup_Name.Name
_,pageDetail.Instrument_Lookup_List,_ = dao.Watchlist_GetList()
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Holding_TemplateEdit, w, r, pageDetail)


}

//Holding_HandlerSave is the handler used process the saving of an Holding
func Holding_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Holding
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Holding_SYSId)
		item.Code = r.FormValue(dm.Holding_Code)
		item.Name = r.FormValue(dm.Holding_Name)
		item.Type = r.FormValue(dm.Holding_Type)
		item.Portfolio = r.FormValue(dm.Holding_Portfolio)
		item.Description = r.FormValue(dm.Holding_Description)
		item.Amount = r.FormValue(dm.Holding_Amount)
		item.Instrument = r.FormValue(dm.Holding_Instrument)
		item.Price = r.FormValue(dm.Holding_Price)
		item.SYSCreated = r.FormValue(dm.Holding_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.Holding_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Holding_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.Holding_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.Holding_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Holding_SYSUpdatedHost)
		item.Symbol = r.FormValue(dm.Holding_Symbol)
		item.Type_Lookup = r.FormValue(dm.Holding_Type_Lookup)
		item.Portfolio_Lookup = r.FormValue(dm.Holding_Portfolio_Lookup)
		
		
		item.Instrument_Lookup = r.FormValue(dm.Holding_Instrument_Lookup)
		
	

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END

	dao.Holding_Store(item,r)	

	http.Redirect(w, r, Holding_Redirect, http.StatusFound)
}

//Holding_HandlerNew is the handler used process the creation of an Holding
func Holding_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Holding_Page{
		Title:       CardTitle(dm.Holding_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Holding_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.Type = ""
pageDetail.Portfolio = ""
pageDetail.Description = ""
pageDetail.Amount = ""
pageDetail.Instrument = ""
pageDetail.Price = ""
pageDetail.SYSCreated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.Symbol = ""


// Automatically generated 20/01/2022 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Type_Lookup = ""
_,pageDetail.Type_Lookup_List,_ = dao.HoldingType_GetList()
pageDetail.Portfolio_Lookup = ""
_,pageDetail.Portfolio_Lookup_List,_ = dao.Portfolio_GetList()
pageDetail.Instrument_Lookup = ""
_,pageDetail.Instrument_Lookup_List,_ = dao.Watchlist_GetList()
// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Holding_TemplateNew, w, r, pageDetail)

}

//Holding_HandlerDelete is the handler used process the deletion of an Holding
func Holding_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Holding_QueryString)

	dao.Holding_Delete(searchID)	

	http.Redirect(w, r, Holding_Redirect, http.StatusFound)
}