package handler

import (
	"database/sql"
	"dts-task/helper"
	"dts-task/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskService.GetTasks()

	if err != nil {
		response := helper.APIResponse("Get task failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Get task success", http.StatusOK, "success", tasks)
	c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	var request task.GetTaskRequest

	err := c.ShouldBindUri(&request)

	if err != nil {
		response := helper.APIResponse("Failed to get task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	task, err := h.taskService.GetTask(request)

	if err != nil {
		if err == sql.ErrNoRows {
			response := helper.APIResponse("Data not found", http.StatusNotFound, "error", nil)
			c.JSON(http.StatusNotFound, response)
		} else {
			response := helper.APIResponse("Failed to get task", http.StatusInternalServerError, "error", nil)
			c.JSON(http.StatusInternalServerError, response)
		}
		return
	}

	response := helper.APIResponse("Success get task", http.StatusOK, "success", task)
	c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var request task.CreateTaskRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		response := helper.APIResponse("Failed to create task", http.StatusUnprocessableEntity, "error", gin.H{
			"errors": helper.FormatValidationError(err),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	task, err := h.taskService.CreateTask(request)
	if err != nil {
		response := helper.APIResponse("Failed to create task", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Success to create task", http.StatusOK, "success", task)
	c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var requestId task.GetTaskRequest

	err := c.ShouldBindUri(&requestId)
	if err != nil {
		response := helper.APIResponse("Failed to update task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var requestData task.CreateTaskRequest
	err = c.ShouldBindJSON(&requestData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to update task", http.StatusUnprocessableEntity, "error", map[string]interface{}{
			"errors": errors,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	task, err := h.taskService.UpdateTask(requestId, requestData)

	if err != nil {
		if err == sql.ErrNoRows {
			response := helper.APIResponse("Failed to update task", http.StatusNotFound, "error", map[string]interface{}{
				"errors": "Data not found",
			})
			c.JSON(http.StatusNotFound, response)
		} else {
			response := helper.APIResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", map[string]interface{}{
				"errors": err.Error(),
			})
			c.JSON(http.StatusUnprocessableEntity, response)
		}
		return
	}

	response := helper.APIResponse("Success update task", http.StatusOK, "success", task)
	c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	var request task.GetTaskRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		response := helper.APIResponse("Failed to delete task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.taskService.DeleteTask(request)

	if err != nil {
		if err == sql.ErrNoRows {
			response := helper.APIResponse("Failed to update task", http.StatusNotFound, "error", map[string]interface{}{
				"errors": "Data not found",
			})
			c.JSON(http.StatusNotFound, response)
		} else {
			response := helper.APIResponse("Failed to update task", http.StatusUnprocessableEntity, "error", map[string]interface{}{
				"errors": err.Error(),
			})
			c.JSON(http.StatusUnprocessableEntity, response)
		}
		return
	}

	response := helper.APIResponse("Success delete task", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
