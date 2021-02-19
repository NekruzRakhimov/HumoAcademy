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
		main.GET("/images/:f2/:img", h.getImg)
	}

	api := router.Group("/api")
	{
		courses := api.Group("/courses")
		{
			//courses.GET("/", h.getAllCourses)
			courses.GET("/", h.getAllMiniCourses)
			courses.GET("/:id", h.getCourseById)
			courses.POST("/:id", h.registerToCourse)
		}

		news := api.Group("/news")
		{
			//news.GET("", h.getAllNews)
			news.GET("/:id", h.getNewsById)
			news.GET("/", h.getAllMiniNews)
		}

		admin := api.Group("/admin",)
		{
			admin.POST("/sign-in", h.adminSignIn)
			admin.GET("/subscribed-users", h.adminIdentity, h.getAllSubscribedUsers)
			admin.POST("/send-mail", h.adminIdentity, h.SendMail)
			courses := admin.Group("/courses", h.adminIdentity)
			{
				courses.POST("/", h.createCourse)
				courses.PATCH("/:id", h.changeCourseImg)
				courses.PUT("/:id", h.editCourse)
				courses.GET("/:id", h.changeCourseStatus)
			}
			news := admin.Group("/news", h.adminIdentity)
			{
				news.POST("/", h.createNews)
				news.PATCH("/:id", h.changeNewsImg)
				news.PUT("/:id", h.editNews)
				news.GET("/:id/", h.changeNewsStatus)
			}
			users := admin.Group("/users", h.adminIdentity)
			{
				users.GET("/:course_id", h.getAllCourseUsers)
				//users.GET("/:id", h.getUserById)
			}

		}

		user := api.Group("/user",)
		{
			user.POST("/", h.createUser)
		}
	}

	return router
}
