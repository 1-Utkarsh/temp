package tasks

import (
	"log"

	"gorm.io/gorm"
)

func GetTaskByID(db *gorm.DB, id string) (Task, error) {
	log.Default().Println("Fetching task with ID:", id)
	var task Task
	result := db.First(&task, id)
	return task, result.Error
}

func GetAllTasks(db *gorm.DB) ([]Task, error) {
	log.Default().Println("Fetching all tasks")
	var tasksList []Task
	result := db.Order("created_at desc").Find(&tasksList)
	return tasksList, result.Error
}

func CreateTask(db *gorm.DB, newTask Task) (Task, error) {
	log.Default().Println("Creating new task:", newTask)
	result := db.Create(&newTask)
	return newTask, result.Error
}

func UpdateTaskByID(db *gorm.DB, id string, updatedData UpdateTask) error {
	log.Default().Println("Fetching task for update with ID:", id)
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return err
	}

	log.Default().Println("Updating task with ID:", id)
	return db.Model(&task).Updates(updatedData).Error
}

func DeleteTaskByID(db *gorm.DB, id string) error {
	log.Default().Println("Deleting task with ID:", id)
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return err
	}

	return db.Delete(&Task{}, id).Error
}
