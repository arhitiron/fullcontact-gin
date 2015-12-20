package dao

import (
	"fullcontact-gin/app/core"
	"fullcontact-gin/app/models"
	"log"
)

type TargetPersonDao struct {}

func (d *TargetPersonDao) Save(person models.TargetPerson) {
	var err error
	if (person.Id != nil) {
		err = core.DbMap.Update(&person)
	} else {
		err = core.DbMap.Insert(&person)
	}

	if err != nil {
		log.Fatalf("Failed to store data: %s", err)
	}
}

func (d *TargetPersonDao) FindAll() []models.TargetPerson {
	var persons []models.TargetPerson
	_, err := core.DbMap.Select(&persons, "select * from target_persons")
	if err != nil {
		log.Fatalf("Error due select from database: %s", err)
	}
	return persons
}

func (d *TargetPersonDao) FindAllNotProcessed() []models.TargetPerson {
	var persons []models.TargetPerson
	_, err := core.DbMap.Select(&persons, "select * from target_persons where processed = 0")
	if err != nil {
		log.Fatalf("Error due select from database: %s", err)
	}
	return persons
}
