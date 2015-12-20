package dao

import (
	"fullcontact-gin/app/core"
	"fullcontact-gin/app/models"
)

type ContactDao struct {}

func (this *ContactDao) Save(item models.Person) {
	err := core.DbMap.Insert(&item)
	if err != nil {
		panic(err)
	}
}

func (this *ContactDao) FindAll() []models.Person {
	var items []models.Person
	_, err := core.DbMap.Select(&items, "select * from items") //simple fetch test
	if err != nil {
		panic(err)
	}

	return items
}

func (this *ContactDao) FindById(id int) models.Person {
	var item models.Person
	err := core.DbMap.SelectOne(&item, "SELECT * FROM items WHERE id=?", id)

	if err != nil {
		panic(err)
	}

	return item
}
