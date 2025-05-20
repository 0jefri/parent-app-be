package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/parent-app-be/middleware"
	"github.com/parent-app-be/models"
)

func GetChildren(w http.ResponseWriter, r *http.Request) {
	parentID := r.Context().Value(middleware.ContextKeyParentID).(uint)

	var children []models.Child
	if err := models.DB.Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		http.Error(w, "Failed to get children", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(children)
}

func CreateChild(w http.ResponseWriter, r *http.Request) {
	parentID := r.Context().Value(middleware.ContextKeyParentID).(uint)

	var input struct {
		Name      string     `json:"name"`
		BirthDate *time.Time `json:"birth_date,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	child := models.Child{
		ParentID:  parentID,
		Name:      input.Name,
		BirthDate: input.BirthDate,
	}

	if err := models.DB.Create(&child).Error; err != nil {
		http.Error(w, "Failed to create child", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Child added",
		"child_id": child.ID,
	})
}

func UpdateChild(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var child models.Child
	if err := models.DB.First(&child, id).Error; err != nil {
		http.Error(w, "Child not found", http.StatusNotFound)
		return
	}

	var updated models.Child
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	child.Name = updated.Name
	child.BirthDate = updated.BirthDate

	if err := models.DB.Save(&child).Error; err != nil {
		http.Error(w, "Failed to update child", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Child updated",
	})
}

func DeleteChild(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := models.DB.Delete(&models.Child{}, id).Error; err != nil {
		http.Error(w, "Failed to delete child", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Child deleted",
	})
}
