package tasks

import (
	"encoding/json"
	"log"
	"net/http"

	db "github.com/1-Utkarsh/temp/store"
	"github.com/1-Utkarsh/temp/store/tasks"
	"github.com/1-Utkarsh/temp/util"
	"github.com/go-chi/chi"
)

func Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", GetTasksByIdHandler)
	r.Get("/", GetAllTasksHandler)
	r.Post("/tasks", CreateTasksHandler)

	return r
}

type TaskResponse struct {
	Task  *tasks.Task
	Error *util.ErrorResponse
}

type AllTasksResponse struct {
	Tasks *[]tasks.Task
	Error *util.ErrorResponse
}

func GetTasksByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	w.Header().Set("Content-Type", "application/json")

	var taskRes TaskResponse
	id := chi.URLParam(r, "id")
	if id == "" {
		taskRes.Error = &util.ErrorResponse{Message: "Empty task id provided", Code: http.StatusBadRequest}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	log.Default().Println("Fetching task with ID:", id)
	t, err := tasks.GetTaskByID(db, id)
	if err.Error() == "record not found" {
		taskRes.Error = &util.ErrorResponse{Message: "Task not found", Code: http.StatusNotFound}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusNotFound)
		w.Write(data)
		return
	} else if err != nil {
		taskRes.Error = &util.ErrorResponse{Message: err.Error(), Code: http.StatusInternalServerError}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	taskRes.Task = &t
	taskRes.Error = nil

	data, _ := json.Marshal(taskRes)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	w.Header().Set("Content-Type", "application/json")

	var allTasksRes AllTasksResponse

	log.Default().Println("Fetching all tasks")
	var tasksList []tasks.Task
	result := db.Find(&tasksList)
	if result.Error != nil {
		allTasksRes.Error = &util.ErrorResponse{Message: result.Error.Error(), Code: http.StatusInternalServerError}
		data, _ := json.Marshal(allTasksRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	allTasksRes.Tasks = &tasksList
	allTasksRes.Error = nil

	data, _ := json.Marshal(allTasksRes)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func CreateTasksHandler(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	w.Header().Set("Content-Type", "application/json")

	var taskRes TaskResponse

	// Parse request body
	var newTask tasks.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		taskRes.Error = &util.ErrorResponse{Message: "Invalid request body", Code: http.StatusBadRequest}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	log.Default().Println("Creating new task:", newTask)
	t, err := tasks.CreateTask(db, newTask)
	if err != nil {
		taskRes.Error = &util.ErrorResponse{Message: err.Error(), Code: http.StatusInternalServerError}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	taskRes.Task = &t
	taskRes.Error = nil

	data, _ := json.Marshal(taskRes)
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
