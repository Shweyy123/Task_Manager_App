package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-manager/config"
	"task-manager/middleware"
	"task-manager/models"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	_, err := config.DB.Exec(
		"INSERT INTO tasks(user_id, title, description, status) VALUES (?, ?, ?, 'pending')",
		userID, task.Title, task.Description,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task created"})
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	rows, err := config.DB.Query(
		"SELECT id, title, description, status FROM tasks WHERE user_id = ?",
		userID,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var update models.UpdateTask
	json.NewDecoder(r.Body).Decode(&update)

	var sets []string
	var vals []interface{}

	if update.Title != nil {
		sets = append(sets, "title = ?")
		vals = append(vals, *update.Title)
	}

	if update.Description != nil {
		sets = append(sets, "description = ?")
		vals = append(vals, *update.Description)
	}

	if update.Status != nil {
		sets = append(sets, "status = ?")
		vals = append(vals, *update.Status)
	}

	if len(sets) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := "UPDATE tasks SET " + strings.Join(sets, ", ") + " WHERE id = ? AND user_id = ?"
	vals = append(vals, taskID, userID)

	_, err = config.DB.Exec(query, vals...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated"})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("DELETE FROM tasks WHERE id = ? AND user_id = ?", taskID, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}
