package api

import (
	"counterapp/internal/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://ashram-connect.vercel.app", "http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//Profiles
	router.GET("/api/profiles", handler.GetProfiles(db))
	router.POST("/api/profiles", handler.CreateProfile(db))
	router.PATCH("/api/profiles/:id", handler.UpdateProfile(db))

	//Visits
	router.GET("/api/profiles/:id/visits", handler.GetVisitsForProfile(db))
	router.POST("/api/visits", handler.AddVisit(db))
	router.PATCH("/api/visits/:id", handler.UpdateVisit(db))

	//Schedules
	router.GET("/api/schedules", handler.GetScheduleForDateRange(db))
	router.POST("/api/schedules", handler.AddSchedule(db))

	//Lockers
	router.GET("/api/lockers", handler.GetAllLockersDetails(db))

	//Feedbacks
	router.GET("/api/feedbacks", handler.GetAllFeedbacks(db))
	router.GET("/api/profiles/:id/feedbacks", handler.GetFeedbackForProfile(db))
	router.POST("/api/feedbacks", handler.AddFeedback(db))

	//SevaTypes
	router.GET("/api/seva-types", handler.GetAllSevaTypes(db))
	router.POST("/api/seva-types", handler.AddSevaType(db))

	//StayAreas
	router.GET("/api/stay-areas", handler.GetAllStayAreas(db))
	router.POST("/api/stay-areas", handler.AddStayArea(db))
	router.GET("/api/stay-areas/occupancy", handler.GetStayAreaDetailsAndOccupancy(db))

	return router
}
