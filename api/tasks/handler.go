package tasks

import (
	"encoding/json"
	"net/http"

	"github.com/1-Utkarsh/temp/store/tasks"
	"github.com/1-Utkarsh/temp/util"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// Routes sets up the routing of task-related endpoints
func Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", GetTasksByIdHandler)
	r.Get("/", GetAllTasksHandler)
	r.Post("/", CreateTasksHandler)
	r.Put("/{id}", UpdateTaskByIdHandler)
	r.Delete("/{id}", DeleteTaskHandler)
	return r
}

// Response structures
type TaskResponse struct {
	Task  *tasks.Task
	Error *util.ErrorResponse
}

type AllTasksResponse struct {
	Tasks *[]tasks.Task
	Error *util.ErrorResponse
}

// CreateTasksHandler handles creating a new task
func CreateTasksHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(util.DbKey).(*gorm.DB)
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

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

	// data quality check
	if newTask.Title == "" {
		taskRes.Error = &util.ErrorResponse{Message: "Task title cannot be empty", Code: http.StatusBadRequest}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	t, err := tasks.CreateTask(db, newTask)
	if err != nil {
		taskRes.Error = &util.ErrorResponse{Message: err.Error(), Code: http.StatusInternalServerError}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	taskRes.Task = &t
	data, _ := json.Marshal(taskRes)
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// GetAllTasksHandler handles fetching all tasks
func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(util.DbKey).(*gorm.DB)
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var allTasksRes AllTasksResponse

	var tasksList []tasks.Task
	tasksList, err := tasks.GetAllTasks(db)
	if err != nil {
		allTasksRes.Error = &util.ErrorResponse{Message: err.Error(), Code: http.StatusInternalServerError}
		data, _ := json.Marshal(allTasksRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	allTasksRes.Tasks = &tasksList
	data, _ := json.Marshal(allTasksRes)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetTasksByIdHandler handles fetching a task by ID
func GetTasksByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(util.DbKey).(*gorm.DB)
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var taskRes TaskResponse
	id := chi.URLParam(r, "id")
	if id == "" {
		taskRes.Error = &util.ErrorResponse{Message: "Empty task id provided", Code: http.StatusBadRequest}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	t, err := tasks.GetTaskByID(db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			taskRes.Error = &util.ErrorResponse{Message: "Task not found", Code: http.StatusNotFound}
			data, _ := json.Marshal(taskRes)
			w.WriteHeader(http.StatusNotFound)
			w.Write(data)
			return
		}
		taskRes.Error = &util.ErrorResponse{Message: err.Error(), Code: http.StatusInternalServerError}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	taskRes.Task = &t
	data, _ := json.Marshal(taskRes)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// UpdateTaskByIdHandler handles updating a task by ID
func UpdateTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(util.DbKey).(*gorm.DB)
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var taskRes TaskResponse
	id := chi.URLParam(r, "id")
	if id == "" {
		taskRes.Error = &util.ErrorResponse{Message: "Empty task id provided", Code: http.StatusBadRequest}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	// check if task exists
	_, err := tasks.GetTaskByID(db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			taskRes.Error = &util.ErrorResponse{Message: "Task not found", Code: http.StatusNotFound}
			data, _ := json.Marshal(taskRes)
			w.WriteHeader(http.StatusNotFound)
			w.Write(data)
			return
		}
		taskRes.Error = &util.ErrorResponse{Message: err.Error(), Code: http.StatusInternalServerError}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	// Parse request body
	var updatedTask tasks.UpdateTask
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		taskRes.Error = &util.ErrorResponse{Message: "Invalid request body", Code: http.StatusBadRequest}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	task, err := tasks.UpdateTaskByID(db, id, updatedTask)
	if err != nil {
		taskRes.Task = &task
		taskRes.Error = &util.ErrorResponse{Message: err.Error(), Code: http.StatusInternalServerError}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	taskRes.Task = &task
	data, _ := json.Marshal(taskRes)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// DeleteTaskHandler handles deleting a task by ID
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(util.DbKey).(*gorm.DB)
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var taskRes TaskResponse
	id := chi.URLParam(r, "id")
	if id == "" {
		taskRes.Error = &util.ErrorResponse{Message: "Empty task id provided", Code: http.StatusBadRequest}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	err := tasks.DeleteTaskByID(db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			taskRes.Error = &util.ErrorResponse{Message: "Task not found", Code: http.StatusNotFound}
			data, _ := json.Marshal(taskRes)
			w.WriteHeader(http.StatusNotFound)
			w.Write(data)
			return
		}
		taskRes.Error = &util.ErrorResponse{Message: err.Error(), Code: http.StatusInternalServerError}
		data, _ := json.Marshal(taskRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	data, _ := json.Marshal(taskRes)
	w.WriteHeader(http.StatusNoContent)
	w.Write(data)
}
