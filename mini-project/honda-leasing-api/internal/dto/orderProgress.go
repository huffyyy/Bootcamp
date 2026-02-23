package dto

import "time"

type OrderProgressResponse struct {
	TaskID   int64      `json:"task_id"`
	TaskName string     `json:"task_name"`
	Status   string     `json:"status"`
	Date     *time.Time `json:"date,omitempty"`
	IsDone   bool       `json:"is_done"`
}

type UpdateTaskStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
