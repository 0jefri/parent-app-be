package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/parent-app-be/middleware"
	"github.com/parent-app-be/models"
)

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	parentID := r.Context().Value(middleware.ContextKeyParentID).(uint)

	var input struct {
		ChildID    uint   `json:"child_id"`
		DeviceName string `json:"device_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	// Validasi: pastikan child ini milik parent yang sedang login
	var child models.Child
	if err := models.DB.First(&child, input.ChildID).Error; err != nil || child.ParentID != parentID {
		http.Error(w, "Invalid child_id", http.StatusForbidden)
		return
	}

	device := models.Device{
		ChildID:    input.ChildID,
		DeviceName: input.DeviceName,
		Status:     "unlocked",
	}

	if err := models.DB.Create(&device).Error; err != nil {
		http.Error(w, "Failed to create device", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Device created",
		"device_id": device.ID,
	})
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	parentID := r.Context().Value(middleware.ContextKeyParentID).(uint)

	var devices []models.Device
	err := models.DB.Joins("JOIN children ON children.id = devices.child_id").
		Where("children.parent_id = ?", parentID).
		Find(&devices).Error
	if err != nil {
		http.Error(w, "Failed to fetch devices", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(devices)
}

func LockDevice(w http.ResponseWriter, r *http.Request) {
	deviceID, _ := strconv.Atoi(mux.Vars(r)["id"])

	var device models.Device
	if err := models.DB.First(&device, deviceID).Error; err != nil {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}

	device.Status = "locked"
	if err := models.DB.Save(&device).Error; err != nil {
		http.Error(w, "Failed to lock device", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Device locked",
	})
}

func UnlockDevice(w http.ResponseWriter, r *http.Request) {
	deviceID, _ := strconv.Atoi(mux.Vars(r)["id"])

	var device models.Device
	if err := models.DB.First(&device, deviceID).Error; err != nil {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}

	device.Status = "unlocked"
	if err := models.DB.Save(&device).Error; err != nil {
		http.Error(w, "Failed to unlock device", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Device unlocked",
	})
}
