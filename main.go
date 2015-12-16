package main

import (
	"github.com/arhitiron/fullcontact-gin/app/core"
	"github.com/arhitiron/fullcontact-gin/app/routers"
)

func main() {
	core.InitCfg()
	core.InitDb()

	port := core.Cfg.Global.Port

	router := routers.Router{}
	router.GetRouter().Run(":" + port)
}
