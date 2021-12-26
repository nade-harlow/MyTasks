package task

import "MyTasks/models"

type Task interface {
	AddTask(task models.Task) error
	ReadAll() ([]models.Task, error)
	UpdateStatus(reference string) error
	DeleteTask(reference string) error
}
