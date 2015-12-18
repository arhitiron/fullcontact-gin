package routers

import (
	ctrl "fullcontact-gin/app/controllers"
	"fullcontact-gin/app/models"
	"github.com/gin-gonic/gin"
	"fullcontact-gin/app/core"
)

type Router struct{}

func (r *Router) createApi(e *gin.Engine) *gin.RouterGroup {
	return e.Group("/api")
}

func (r *Router) setVersion(e *gin.RouterGroup) *gin.RouterGroup {
	version := core.Cfg.Global.ApiVersion
	return e.Group(version)
}

func (r *Router) GetRouter() *gin.Engine {
	router := gin.Default()
	api := r.createApi(router)
	api = r.setVersion(api)

	api.POST("/handle-person", func(c *gin.Context) {
		var item models.Contact
		c.BindJSON(&item)
		kafka := ctrl.KafkaController{}
		kafka.PublishContact(item)
		c.JSON(200, item)
	})

	api.GET("/persons", func(c *gin.Context) {
		c.JSON(200, ctrl.PersonController{}.GetAll())
	})

	api.POST("/persons", func(c *gin.Context) {
		var persons []models.TargetPerson
		c.BindJSON(&persons)
		ctrl.PersonController{}.AddItems(persons)
		c.JSON(200, gin.H{"success": true})
	})

	api.POST("/process", func(c *gin.Context) {

		c.JSON(200, gin.H{"status": "started"})
	})

	return router
}
