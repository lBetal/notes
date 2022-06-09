package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lBetal/notes/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		devices := api.Group("/devices")
		{
			devices.POST("/", h.createDevice)
			devices.GET("/", h.getAllDevice)
			devices.GET("/:id", h.getDeviceById)
			devices.PUT("/:id", h.updateDevice)
			devices.DELETE("/:id", h.deleteDevice)
			photos := devices.Group(":id/photos")
			{
				photos.POST("/", h.createPhoto)
				photos.GET("/", h.getAllPhotos)
				photos.GET("/:photo_id", h.getPhotoById)
				photos.PUT("/:photo_id", h.updatePhoto)
				photos.DELETE("/:photo_id", h.deletePhoto)
			}
			videos := devices.Group(":id/videos")
			{
				videos.POST("/")
				videos.GET("/")
			}
			audios := devices.Group(":id/audios")
			{
				audios.POST("/")
				audios.GET("/")
			}
			messages := devices.Group(":id/messages")
			{
				messages.POST("/")
				messages.GET("/")
			}
		}
	}

	return router
}
