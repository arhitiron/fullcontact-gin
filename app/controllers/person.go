package controllers

import (
	"fullcontact-gin/app/dao"
	"fullcontact-gin/app/models"
	"fullcontact-gin/app/utils"
	"fullcontact-gin/app/core"
)

type PersonController struct{}

var personDao dao.TargetPersonDao

func init() {
	personDao = dao.TargetPersonDao{}
}

func (c PersonController) GetAll() []models.TargetPerson {
	r := personDao.FindAll()
	return r
}

func (c PersonController) AddItem(person models.TargetPerson) {
	personDao.Save(person)
}

func (c PersonController) AddItems(persons []models.TargetPerson) {
	for _, person := range persons {
		personDao.Save(person)
	}
}

func (c PersonController) ProcessPersons() {
	persons := personDao.FindAllNotProcessed()
	fc := core.Cfg.Fullcontact
	for _, person := range persons {
		utils.SendToFullcontact(person.Email, fc.WebhookUrl, fc.ApiKey)
		person.Processed = false
		personDao.Save(person)
	}
}

