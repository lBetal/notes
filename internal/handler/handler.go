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
				videos.POST("/", h.createVideo)
				videos.GET("/", h.getAllVideos)
				videos.GET("/:video_id", h.getVideoById)
				videos.PUT("/:video_id", h.updateVideo)
				videos.DELETE("/:video_id", h.deleteVideo)
			}
			audios := devices.Group(":id/audios")
			{
				audios.POST("/", h.createAudio)
				audios.GET("/", h.getAllAudios)
				audios.GET("/:audio_id", h.getAudioById)
				audios.PUT("/:audio_id", h.updateAudio)
				audios.DELETE("/:audio_id", h.deleteAudio)
			}
			messages := devices.Group(":id/messages")
			{
				messages.POST("/", h.createMessage)
				messages.GET("/", h.getAllMessages)
				messages.GET("/:message_id", h.getMessageById)
				messages.PUT("/:message_id", h.updateMessage)
				messages.DELETE("/:message_id", h.deleteMessage)
			}
		}
	}

	return router
}
