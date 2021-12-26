package main

import (
	"MyTasks/controller"
	"MyTasks/models/mongoDb"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.Route(r)
	db := mongoDb.MongoRepository{}
	db.Init()
	_ = controller.Server{Task: db}
	err := r.Run(":8081")
	if err != nil {
		panic("Unable to start application")
	}
}
