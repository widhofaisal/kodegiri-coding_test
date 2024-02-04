package main

import (
	"kodegiri/kodingtest/config"
	"kodegiri/kodingtest/route"

	"github.com/labstack/echo/v4/middleware"

)


func main() {
	config.InitDB()
	e := route.Init()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// utils.AutoCreateAdmin()
	// utils.AutoInsertMaterials()
	// utils.AutoCreateFolderReceipts()
	
	e.Start(":8000")
	
}
