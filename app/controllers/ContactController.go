package controllers

import (
	"fullcontact-gin/app/models"
	"fullcontact-gin/app/dao"
	"log"
)

type ContactController struct {

}

var itemDao dao.ContactDao
func init() {
	itemDao = dao.ContactDao{}
}

func (c ContactController) GetAllItems() []models.Contact {
	r := itemDao.FindAll()
	log.Println(r)
	return r
}

func (c ContactController) GetItem(id int) models.Contact {
	//	kafka := KafkaController{}
	//	item := itemDao.FindById(id)
	//	itemStr, err := json.Marshal(item)
	//	if err != nil {
	//		panic(err)
	//	}
	//	kafka.Publish(itemStr)
	return models.Contact{}
}