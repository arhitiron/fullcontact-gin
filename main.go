package main

import (
	"fullcontact-gin/app/core"
	"fullcontact-gin/app/routers"
)

func main() {
	core.InitCfg()
	core.InitDb()

	port := core.Cfg.Global.Port

	router := routers.Router{}
	router.GetRouter().Run(":" + port)
}
