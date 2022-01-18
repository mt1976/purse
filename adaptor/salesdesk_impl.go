package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/firm.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : SalesDesk (firm)
// Endpoint 	        : SalesDesk (SalesDeskName)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:17
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func SalesDesk_Delete_Impl(id string) error {
	var er error

	message := "Implement SalesDesk_Delete: " + id

	// Implement SalesDesk_Delete_Impl in firm_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := SalesDesk_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func SalesDesk_Update_Impl(item dm.SalesDesk, usr string) error {
	var er error

	message := "Implement SalesDesk_Update: " + item.Name

	// Implement SalesDesk_Update_Impl in firm_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := SalesDesk_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
