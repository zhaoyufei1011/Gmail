package main

import (
	"Autosino/router"
)
func main()  {
	router:=router.InitRouter()
	router.Run(":8080")
}