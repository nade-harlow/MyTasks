package controller

import "github.com/gin-gonic/gin"

func (s *Server) Route(r *gin.Engine) {
	r.GET("/", s.Home())
	r.POST("/add", s.AddTask())
	r.PUT("/update/:reference", s.UpdateTaskStatus())
	r.GET("/tasks", s.GetAllTask())
	r.DELETE("/delete/:reference", s.DeleteTask())
}
