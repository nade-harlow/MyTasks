package mongoDb

import (
	"MyTasks/models"
	"log"
	"time"
)

func (r MongoRepository) Add() error {
	task := models.Task{
		Title:        "Eat Food",
		Done:         false,
		CreatedAt:    time.Now().String(),
		LastModified: time.Now().String(),
	}
	err := r.Client.C("task").Insert(task)
	if err != nil {
		log.Println("Unable to insert document: ", err.Error())
		return err
	}
	return nil
}
