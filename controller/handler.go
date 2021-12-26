package controller

import "github.com/gin-gonic/gin"

func (s *Server) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := s.Task.Add()
		if err != nil {
			c.JSON(500, gin.H{"error": "can't ==="})
			return
		}
	}
}
