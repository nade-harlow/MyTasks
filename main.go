package main

import (
	"MyTasks/controller"
	"MyTasks/models/mongoDb"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := mongoDb.MongoRepository{}
	db.Init()
	s := controller.Server{Task: db}
	s.Route(r)

	err := r.Run(":8081")
	if err != nil {
		panic("Unable to start application")
	}
}
