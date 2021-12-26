package controller

import (
	"MyTasks/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (s *Server) AddTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		task := models.Task{
			Reference:    uuid.New().String(),
			CreatedAt:    time.Now().String(),
			LastModified: time.Now().String(),
		}
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := s.Task.AddTask(task); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}

func (s *Server) GetAllTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := s.Task.ReadAll()
		if err != nil {
			return
		}
		c.JSON(200, gin.H{"tasks": tasks})
	}
}

func (s Server) UpdateTaskStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param("reference")
		if err := s.Task.UpdateStatus(reference); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "successful"})
	}
}

func (s Server) DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param("reference")
		if err := s.Task.DeleteTask(reference); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "successful"})
	}
}
