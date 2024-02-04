package main

import (
	"kodegiri/kodingtest/config"
	"kodegiri/kodingtest/route"
)


func main() {
	config.InitDB()
	e := route.Init()
	
	// utils.AutoCreateAdmin()
	// utils.AutoInsertMaterials()
	// utils.AutoCreateFolderReceipts()
	
	e.Start(":8000")
	
}
