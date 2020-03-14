package router

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine *gin.Engine
}

func (h *Handler) SetupApiEndpoints() {
	// h.Engine.Use(static.Serve("/", static.LocalFile("./sites/home", true)))
	h.Engine.Use(static.Serve("/", static.LocalFile("./sites/admin", true)))

	apiPrograms := h.Engine.Group("/api/programs/")
	{
		apiPrograms.GET("/v1/all", controllers.GetAllPrograms)
		apiPrograms.POST("/v1/create", controllers.CreateProgram)
		apiPrograms.GET("/v1/program/:programId", controllers.GetProgramById)
		apiPrograms.POST("/v1/program/:programId", controllers.UpdateProgram)
		apiPrograms.DELETE("/v1/program/:programId", controllers.DeleteProgram)
	}
	// apiClasses := router.Group("api/classes/")
	apiLocations := h.Engine.Group("api/locations/")
	{
		apiLocations.GET("v1/all", controllers.GetAllLocations)
		apiLocations.POST("/v1/create", controllers.CreateLocation)
		apiLocations.GET("/v1/location/:locId", controllers.GetLocationById)
		apiLocations.POST("/v1/location/:locId", controllers.UpdateLocation)
		apiLocations.DELETE("/v1/location/:locId", controllers.DeleteLocation)
	}
	apiAnnounces := h.Engine.Group("api/announcements/")
	{
		apiAnnounces.GET("/v1/all", controllers.GetAllAnnouncements)
		apiAnnounces.POST("/v1/create", controllers.CreateAnnouncement)
		apiAnnounces.GET("/v1/announce/:id", controllers.GetAnnouncementById)
		apiAnnounces.POST("/v1/announce/:id", controllers.UpdateAnnouncement)
		apiAnnounces.DELETE("/v1/announce/:id", controllers.DeleteAnnouncement)
	}
	// apiAchieve := router.Group("api/achieve/")
	// apiSemesters := router.Group("api/semesters/")
	// apiUsers := router.Group("api/users/")
	// apiAccounts := router.Group("api/accounts/")
}
