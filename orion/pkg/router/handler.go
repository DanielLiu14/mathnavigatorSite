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
		apiPrograms.GET("/v1/unpublished", controllers.GetAllUnpublishedPrograms)
		apiPrograms.POST("/v1/publish", controllers.PublishPrograms)
	}
	apiClasses := h.Engine.Group("api/classes/")
	{
		apiClasses.GET("/v1/all", controllers.GetAllClasses)
		apiClasses.POST("/v1/create", controllers.CreateClass)
		apiClasses.GET("/v1/class/:classId", controllers.GetClassById)
		apiClasses.POST("/v1/class/:classId", controllers.UpdateClass)
		apiClasses.DELETE("/v1/class/:classId", controllers.DeleteClass)
		apiClasses.GET("/v1/classes/program/:programId", controllers.GetClassesByProgram)
		apiClasses.GET("/v1/classes/semester/:semesterId", controllers.GetClassesBySemester)
		apiClasses.GET("/v1/classes/program/:programId/semester/:semesterId", controllers.GetClassesByProgramAndSemester)
	}
	apiLocations := h.Engine.Group("api/locations/")
	{
		apiLocations.GET("/v1/all", controllers.GetAllLocations)
		apiLocations.POST("/v1/create", controllers.CreateLocation)
		apiLocations.GET("/v1/location/:locId", controllers.GetLocationById)
		apiLocations.POST("/v1/location/:locId", controllers.UpdateLocation)
		apiLocations.DELETE("/v1/location/:locId", controllers.DeleteLocation)
		apiLocations.GET("/v1/unpublished", controllers.GetAllUnpublishedLocations)
		apiLocations.POST("/v1/publish", controllers.PublishLocations)
	}
	apiAnnounces := h.Engine.Group("api/announcements/")
	{
		apiAnnounces.GET("/v1/all", controllers.GetAllAnnouncements)
		apiAnnounces.POST("/v1/create", controllers.CreateAnnouncement)
		apiAnnounces.GET("/v1/announcement/:id", controllers.GetAnnouncementById)
		apiAnnounces.POST("/v1/announcement/:id", controllers.UpdateAnnouncement)
		apiAnnounces.DELETE("/v1/announcement/:id", controllers.DeleteAnnouncement)
	}
	apiAchieves := h.Engine.Group("api/achievements/")
	{
		apiAchieves.GET("/v1/all", controllers.GetAllAchievements)
		apiAchieves.POST("/v1/create", controllers.CreateAchievement)
		apiAchieves.GET("/v1/achievement/:id", controllers.GetAchievementById)
		apiAchieves.POST("/v1/achievement/:id", controllers.UpdateAchievement)
		apiAchieves.DELETE("/v1/achievement/:id", controllers.DeleteAchievement)
		apiAchieves.GET("/v1/years", controllers.GetAllAchievementsGroupedByYear)
	}
	apiSemesters := h.Engine.Group("api/semesters/")
	{
		apiSemesters.GET("/v1/all", controllers.GetAllSemesters)
		apiSemesters.POST("/v1/create", controllers.CreateSemester)
		apiSemesters.GET("/v1/semester/:semesterId", controllers.GetSemesterById)
		apiSemesters.POST("/v1/semester/:semesterId", controllers.UpdateSemester)
		apiSemesters.DELETE("/v1/semester/:semesterId", controllers.DeleteSemester)
	}
	apiSessions := h.Engine.Group("api/sessions/")
	{
		apiSessions.POST("/v1/create", controllers.CreateSession)
		apiSessions.GET("/v1/session/:id", controllers.GetSessionById)
		apiSessions.POST("/v1/session/:id", controllers.UpdateSession)
		apiSessions.DELETE("/v1/session/:id", controllers.DeleteSession)
		apiSessions.GET("/v1/class/:classId", controllers.GetAllSessionsByClassId)
		apiSessions.GET("/v1/unpublished", controllers.GetAllUnpublishedSessions)
		apiSessions.POST("/v1/publish", controllers.PublishSessions)
	}
	apiUsers := h.Engine.Group("api/users/")
	{
		apiUsers.GET("/v1/all", controllers.GetAllUsers)
		apiUsers.POST("/v1/create", controllers.CreateUser)
		apiUsers.GET("/v1/user/:id", controllers.GetUserById)
		apiUsers.POST("/v1/user/:id", controllers.UpdateUser)
		apiUsers.DELETE("/v1/user/:id", controllers.DeleteUser)
	}
	// apiAccounts := router.Group("api/accounts/")
}
