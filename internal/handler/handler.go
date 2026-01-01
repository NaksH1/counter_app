package handler

import (
	"counterapp/internal/dao"
	"counterapp/internal/model"
	"counterapp/internal/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	logKeyError = "error"
)

func GetProfiles(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		profiles, err := dao.GetProfilesData(db)
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}
		c.JSON(200, profiles)
	}
}

type CreateProfileRequest struct {
	Profile model.Profile `json:"profile"`
}

func CreateProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateProfileRequest
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		profile, err := dao.CreateProfile(db, &req.Profile)
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}
		c.JSON(200, profile)
	}
}

type UpdateProfileRequest struct {
	Name        *string         `json:"name,omitempty"`
	PhoneNumber *string         `json:"phone_number,omitempty"`
	Gender      *model.Gender   `json:"gender,omitempty"`
	Category    *model.Category `json:"category,omitempty"`
	IsBlocked   *bool           `json:"is_blocked,omitempty"`
	Remarks     *string         `json:"remarks,omitempty"`
}

func UpdateProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		profileID := c.Param("id")
		var req UpdateProfileRequest
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		updatedProfile, err := dao.UpdateProfile(db, profileID, &dao.ProfileUpdate{
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
			Gender:      req.Gender,
			Category:    req.Category,
			IsBlocked:   req.IsBlocked,
			Remarks:     req.Remarks,
		})
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}
		c.JSON(200, updatedProfile)
	}
}

func GetVisitsForProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		profileID := c.Param("id")
		if profileID == "" {
			c.JSON(400, gin.H{logKeyError: "Please enter profile_id"})
		}

		visits, err := dao.GetVisitsByProfileID(db, profileID)
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}
		c.JSON(200, visits)
	}
}

func GetAllLockersDetails(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		lockers, err := dao.GetAllLockers(db)
		if err != nil {
			c.JSON(500, gin.H{
				logKeyError: err.Error(),
			})
			return
		}
		c.JSON(200, lockers)
	}
}

func GetScheduleForDateRange(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sDate := c.Query("start_date")
		eDate := c.Query("end_date")
		if sDate == "" || eDate == "" {
			c.JSON(400, gin.H{logKeyError: "Cannot have empty values for start and end date"})
			return
		}

		startDate, err := util.FormatDateToISO(sDate)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: err})
			return
		}

		endDate, err := util.FormatDateToISO(eDate)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: err})
			return
		}

		if !util.CompareDates(*startDate, *endDate) {
			c.JSON(400, gin.H{logKeyError: "Start date cannot be less than end date"})
			return
		}

		schedule, err := dao.GetScheduleForDateRange(db, *startDate, *endDate)
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}
		c.JSON(200, schedule)
	}
}

type AddVisitRequest struct {
	ProfileID     string  `json:"profile_id"`
	Email         string  `json:"email"`
	ArrivalDate   string  `json:"arrival_date"`
	DepartureDate *string `json:"departure_date,omitempty"`
	StayAreaID    string  `json:"stay_area_id"`
	ProfileStatus *string `json:"profile_status,omitempty"`
}

func AddVisit(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddVisitRequest
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Invalid request body"})
			return
		}
		profile, err := dao.GetProfileByID(db, req.ProfileID)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Profile not found"})
			return
		}

		visitsWithStatusCheckedIn, err := dao.GetVisitsByProfileID(db, profile.ID.String())
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Unable to fetch visit for profile"})
			return
		}
		for _, visit := range visitsWithStatusCheckedIn {
			if visit.Status != "checked-in" {
				continue
			}
			checkedOutStatus := model.StatusCheckedOut
			_, err = dao.UpdateVisit(db, visit.ID.String(), dao.UpdateVisitRequest{
				Status: &checkedOutStatus,
			})
			if err != nil {
				fmt.Printf("unable to update status for the visit: %s, %v", visit.ID, err)
				continue
			}
		}

		arrivalDate, err := util.FormatDateToISO(req.ArrivalDate)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: err})
			return
		}

		var departureDate *time.Time
		if req.DepartureDate != nil {
			departureDate, err = util.FormatDateToISO(*req.DepartureDate)
			if err != nil {
				c.JSON(400, gin.H{logKeyError: err})
				return
			}
		}

		profileUUID, err := uuid.Parse(req.ProfileID)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Invalid profile ID format"})
			return
		}

		stayAreaUUID, err := uuid.Parse(req.StayAreaID)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Invalid stay area ID format"})
			return
		}

		visit, err := dao.AddVisit(db, dao.AddVisitRequest{
			ProfileID:     profileUUID,
			ArrivalDate:   *arrivalDate,
			DepartureDate: departureDate,
			StayAreaID:    stayAreaUUID,
			Status:        model.StatusCheckedIn,
		})
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err})
			return
		}
		c.JSON(200, visit)
	}
}

type UpdateVisitRequest struct {
	DepartureDate *string `json:"departure_date,omitempty"`
	StayAreaID    *string `json:"stay_area_id,omitempty"`
	LockerID      *string `json:"locker_id,omitempty"`
	Remarks       *string `json:"remarks,omitempty"`
	Status        *string `json:"status,omitempty"`
}

func UpdateVisit(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		visitID := c.Param("id")
		if visitID == "" {
			c.JSON(400, gin.H{logKeyError: "Visit ID is required"})
			return
		}

		var req UpdateVisitRequest
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Invalid request body"})
			return
		}

		var departureDate *time.Time
		if req.DepartureDate != nil {
			departureDate, err = util.FormatDateToISO(*req.DepartureDate)
			if err != nil {
				c.JSON(400, gin.H{logKeyError: "Invalid departure date format"})
				return
			}
		}

		var stayAreaUUID *uuid.UUID
		if req.StayAreaID != nil {
			parsedStayAreaID, err := uuid.Parse(*req.StayAreaID)
			if err != nil {
				c.JSON(400, gin.H{logKeyError: "Invalid stay area ID format"})
				return
			}
			stayAreaUUID = &parsedStayAreaID
		}

		var lockerUUID *uuid.UUID
		if req.LockerID != nil {
			parsedLockerID, err := uuid.Parse(*req.LockerID)
			if err != nil {
				c.JSON(400, gin.H{logKeyError: "Invalid locker ID format"})
				return
			}
			lockerUUID = &parsedLockerID
		}

		var status *model.ProfileStatus
		if req.Status != nil {
			profileStatus := model.ProfileStatus(*req.Status)
			status = &profileStatus
		}

		updatedVisit, err := dao.UpdateVisit(db, visitID, dao.UpdateVisitRequest{
			DepartureDate: departureDate,
			StayAreaID:    stayAreaUUID,
			LockerID:      lockerUUID,
			Remarks:       req.Remarks,
			Status:        status,
		})
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}

		c.JSON(200, updatedVisit)
	}
}

type AddScheduleRequest struct {
	ProfileID string  `json:"profile_id"`
	VisitID   string  `json:"visit_id"`
	SevaType  string  `json:"seva_type"`
	Location  *string `json:"location,omitempty"`
	Date      string  `json:"date"`
}

func AddSchedule(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddScheduleRequest
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// 1. Parse date
		scheduleDate, err := util.FormatDateToISO(req.Date)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid date format"})
			return
		}

		// 2. Validate visit exists and is active
		visit, err := dao.GetVisitByID(db, req.VisitID)
		if err != nil {
			c.JSON(404, gin.H{"error": "Visit not found"})
			return
		}
		if visit.Status != model.StatusCheckedIn {
			c.JSON(400, gin.H{"error": "Visit is not active"})
			return
		}

		// 3. Check for existing schedule (duplicate prevention)
		existingSchedule, err := dao.GetScheduleByProfileAndDate(db, req.ProfileID, *scheduleDate)
		if err == nil && existingSchedule != nil {
			c.JSON(409, gin.H{"error": "Schedule already exists for this date"})
			return
		}

		// 4. Parse UUIDs
		profileUUID, err := uuid.Parse(req.ProfileID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid profile ID format"})
			return
		}

		visitUUID, err := uuid.Parse(req.VisitID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid visit ID format"})
			return
		}

		sevaType, err := dao.GetSevaTypeByName(db, req.SevaType)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Invalid seva type or seva type not found"})
			return
		}

		// 5. Create schedule
		schedule, err := dao.AddSchedule(db, dao.AddScheduleRequest{
			ProfileID:  profileUUID,
			VisitID:    visitUUID,
			SevaTypeID: sevaType.ID,
			Location:   req.Location,
			Date:       *scheduleDate,
		})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, schedule)
	}
}

func GetAllFeedbacks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		feedbacks, err := dao.GetAllFeedbacks(db)
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err})
			return
		}
		c.JSON(200, feedbacks)
	}
}

func GetFeedbackForProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		profileID := c.Param("id")
		feedbacks, err := dao.GetFeedbacksForProfile(db, profileID)
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err})
			return
		}
		c.JSON(200, feedbacks)
	}
}

type AddFeedbackRequest struct {
	ProfileID string  `json:"profile_id"`
	VisitID   *string `json:"visit_id,omitempty"`
	Content   string  `json:"content"`
	Type      string  `json:"type"`
	CreatedBy *string `json:"created_by,omitempty"`
}

func AddFeedback(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddFeedbackRequest
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Invalid request body"})
			return
		}

		profileUUID, err := uuid.Parse(req.ProfileID)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Invalid profile ID format"})
			return
		}

		var visitUUID *uuid.UUID
		if req.VisitID != nil {
			parsedVisitID, err := uuid.Parse(*req.VisitID)
			if err != nil {
				c.JSON(400, gin.H{logKeyError: "Invalid visit ID format"})
				return
			}
			visitUUID = &parsedVisitID
		}

		feedback, err := dao.AddFeedback(db, dao.AddFeedbackRequest{
			ProfileID: profileUUID,
			VisitID:   visitUUID,
			Content:   req.Content,
			Type:      model.FeedbackType(req.Type),
			CreatedBy: req.CreatedBy,
		})

		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}

		c.JSON(201, feedback)
	}
}

func GetAllSevaTypes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sevaTypes, err := dao.GetAllSevaTypes(db)
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}
		c.JSON(200, sevaTypes)
	}
}

type AddSevaTypeRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

func AddSevaType(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddSevaTypeRequest
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Invalid request body"})
			return
		}

		sevaType, err := dao.AddSevaType(db, dao.AddSevaTypeRequest{
			Name:        req.Name,
			Description: req.Description,
		})
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}

		c.JSON(201, sevaType)
	}
}

func GetAllStayAreas(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		stayAreas, err := dao.GetAllStayAreas(db)
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}
		c.JSON(200, stayAreas)
	}
}

type AddStayAreaRequest struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

func AddStayArea(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddStayAreaRequest
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{logKeyError: "Invalid request body"})
			return
		}

		stayArea, err := dao.AddStayArea(db, dao.AddStayAreaRequest{
			Name:     req.Name,
			Capacity: req.Capacity,
		})
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}

		c.JSON(201, stayArea)
	}
}

type StayAreaOccupancyResponse struct {
	StayAreaID           string `json:"stay_area_id"`
	StayName             string `json:"stay_name"`
	Capacity             int    `json:"capacity"`
	CurrentOccupiedCount int    `json:"current_occupied_count"`
	Available            int    `json:"available"`
}

func GetStayAreaDetailsAndOccupancy(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		stayAreasWithOccupancy, err := dao.GetAllStayAreasWithOccupancy(db)
		if err != nil {
			c.JSON(500, gin.H{logKeyError: err.Error()})
			return
		}

		responses := make([]StayAreaOccupancyResponse, 0, len(stayAreasWithOccupancy))
		for _, sa := range stayAreasWithOccupancy {
			responses = append(responses, StayAreaOccupancyResponse{
				StayAreaID:           sa.StayAreaID,
				StayName:             sa.StayName,
				Capacity:             sa.StayCapacity,
				CurrentOccupiedCount: sa.CurrentOccupiedCount,
				Available:            sa.StayCapacity - sa.CurrentOccupiedCount,
			})
		}

		c.JSON(200, responses)
	}
}
