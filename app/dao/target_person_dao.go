package dao

import (
	"fullcontact-gin/app/core"
	"fullcontact-gin/app/models"
)

type TargetPersonDao struct {}

func (d *TargetPersonDao) Save(person models.TargetPerson) {
	err := core.DbMap.Insert(&person)
	if err != nil {
		panic(err)
	}
}

func (d *TargetPersonDao) FindAll() []models.TargetPerson {
	var persons []models.TargetPerson
	_, err := core.DbMap.Select(&persons, "select * from target_persons")
	if err != nil {
		panic(err)
	}

	return persons
}

func (d *TargetPersonDao) FindAllNotProcessed() []models.TargetPerson {
	var persons []models.TargetPerson
	_, err := core.DbMap.Select(&persons, "select * from target_persons where processed = 0")
	if err != nil {
		panic(err)
	}

	return persons
}
