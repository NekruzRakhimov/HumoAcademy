package handler

import (
	"HumoAcademy/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(h.Cors)

	main := router.Group("/")
	{
		main.GET("/", h.getAll)
		main.POST("/", h.addUserForNews)
		main.GET("/getImg/:f1/:f2/:img", h.getImg)
	}

	api := router.Group("/api")
	{
		courses := api.Group("/courses")
		{
			courses.GET("/", h.GetAllCourses)
			courses.GET("/:id", h.getCourseById)
			courses.POST("/:id", h.registerToCourse)
		}

		news := api.Group("/news")
		{
			news.GET("/", h.GetAllNews)
			news.GET("/:id", h.getNewsById)
		}

		admin := api.Group("/admin",)
		{
			adminAuth := admin.Group("/auth")
			{
				adminAuth.POST("/sign-up", h.adminSignUp)
				adminAuth.POST("/sign-in", h.adminSignIn)
			}
			courses := admin.Group("/courses", h.adminIdentity)
			{
				courses.POST("/create-course", h.createCourse)
			}
			news := admin.Group("/news", h.adminIdentity)
			{
				news.POST("/create-news", h.createNews)
			}
		}
	}


	return router
}
