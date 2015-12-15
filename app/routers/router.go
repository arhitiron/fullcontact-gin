package routers

import (
	"github.com/gin-gonic/gin"
	ctrl "fullcontact-gin/app/controllers"
	"strconv"
	"fullcontact-gin/app/models"
)

type Router struct {

}

func (r *Router) createApiWith(e *gin.Engine) *gin.RouterGroup {
	return e.Group("/api")
}

func (r *Router) setVersion1(e *gin.RouterGroup) *gin.RouterGroup {
	return e.Group("/v1")
}

func (r *Router) GetRouter() *gin.Engine {
	router := gin.Default()
	api := r.createApiWith(router)
	apiV1 := r.setVersion1(api)

	apiV1.GET("/contacts", func(c *gin.Context) {
		c.JSON(200, ctrl.ContactController{}.GetAllItems())
	})

	apiV1.GET("/contacts/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		c.JSON(200, ctrl.ContactController{}.GetItem(id))
	})

	apiV1.POST("/contacts", func(c *gin.Context) {
		var item models.Contact
		c.BindJSON(&item)

		kafka := ctrl.KafkaController{}
		kafka.PublishContact(item)

		c.JSON(200, item)
	})

	return router
}

