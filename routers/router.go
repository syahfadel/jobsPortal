package routers

import (
	"jobsPortal/controllers"
	"jobsPortal/middlewares"
	"jobsPortal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartService(db *gorm.DB) *gin.Engine {
	app := gin.Default()

	jobsPortalService := services.JobsPortalService{
		DB: db,
	}

	jobsPortalController := controllers.JobsPortalController{
		DB:                db,
		JobsPortalService: &jobsPortalService,
	}

	user := app.Group("/user")
	{
		user.POST("/register", jobsPortalController.CreateUser)
		user.POST("/login", jobsPortalController.Login)
	}

	job := app.Group("/job")
	{
		job.Use(middlewares.Authentication())
		job.GET("/getJobList", jobsPortalController.GetJobsList)
		job.GET("/getJobDetail/:id", jobsPortalController.GetJobDetail)
	}

	return app
}
