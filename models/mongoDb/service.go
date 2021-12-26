package mongoDb

import (
	"MyTasks/models"
	"github.com/globalsign/mgo/bson"
	"log"
	"time"
)

func (r MongoRepository) AddTask(task models.Task) error {
	err := r.Client.C("task").Insert(task)
	if err != nil {
		log.Println("Unable to insert document: ", err.Error())
		return err
	}
	return nil
}

func (r MongoRepository) ReadAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.Client.C("task").Find(bson.M{"done": false}).All(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r MongoRepository) UpdateStatus(reference string) error {
	if err := r.Client.C("task").Update(bson.M{"reference": reference}, bson.M{"$set": bson.M{"done": true, "last_modified": time.Now()}}); err != nil {
		return err
	}
	return nil
}

func (r MongoRepository) DeleteTask(reference string) error {
	if err := r.Client.C("task").RemoveId(bson.M{"reference": reference}); err != nil {
		return err
	}
	return nil
}
