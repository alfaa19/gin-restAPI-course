package main

import (
	config "github.com/alfaa19/gin-restAPI-course/config/mysql"
	"github.com/alfaa19/gin-restAPI-course/controller"
	"github.com/alfaa19/gin-restAPI-course/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	//router course
	r.GET("api/v1/courses", middleware.AuthMiddleware(), middleware.UserMiddleware(), controller.GetCourse)
	r.GET("api/v1/course/:id", middleware.AuthMiddleware(), middleware.UserMiddleware(), controller.GetDetailCourse)
	r.GET("api/v1/courses/total-courses", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.GetTotalCourse)
	r.GET("api/v1/courses/total-free-courses", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.GetTotalFreeCourse)
	r.POST("api/v1/course", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.CreateCourse)
	r.PUT("api/v1/course/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.UpdateCourse)
	r.DELETE("api/v1/course/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.DeleteCourse)

	//router category
	r.GET("api/v1/course/categories", middleware.AuthMiddleware(), middleware.UserMiddleware(), controller.GetCategory)
	r.GET("api/v1/course/popular-categories", middleware.AuthMiddleware(), middleware.UserMiddleware(), controller.GetPopularCategory)
	r.POST("api/v1/course/category", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.CreateCategory)

	//router auth
	r.POST("api/v1/register", controller.Register)
	r.POST("api/v1/login", controller.Login)

	//router user
	r.GET("api/v1/users/total-users", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.GetTotalUsers)
	r.DELETE("api/v1/user/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.DeleteUser)

	r.Run(":8001")
}
