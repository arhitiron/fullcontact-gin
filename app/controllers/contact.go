package controllers

import (
	"log"

	"github.com/arhitiron/fullcontact-gin/app/dao"
	"github.com/arhitiron/fullcontact-gin/app/models"
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
