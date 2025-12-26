package tasks

import "gorm.io/gorm"

func GetTaskByID(db *gorm.DB, id string) (Task, error) {
	var task Task
	result := db.First(&task, id)
	return task, result.Error
}

func CreateTask(db *gorm.DB, newTask Task) (Task, error) {
	result := db.Create(&newTask)
	return newTask, result.Error
}